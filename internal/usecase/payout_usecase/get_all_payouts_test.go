package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllPayouts(t *testing.T) {
	t.Parallel()

	now := time.Now()
	var partnerID int64 = 1
	var payoutID int64 = 1
	var amount int64 = 1000
	withdrawInfo := domain.WithdrawInfo{
		Method:     domain.BankWithdrawMethod,
		Requisites: "1234 5678 9012 3456",
	}

	type args struct {
		req GetAllPayoutsReq
	}

	tests := []struct {
		name   string
		args   args
		want   []*domain.Payout
		before func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "successful get all payouts",
			args: args{
				req: GetAllPayoutsReq{
					PartnerID: partnerID,
					PayoutID:  0,
					Status:    0,
				},
			},
			want: []*domain.Payout{{
				ID:           payoutID,
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       domain.PendingPayoutStatus,
				CreatedAt:    now,
				UpdateAt:     now,
			}},
			before: func(f useCaseMocks, args args) {
				payout := &domain.Payout{
					ID:           payoutID,
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.PendingPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     now,
				}

				f.repoPayout.EXPECT().GetAllPayouts(dto.PayoutFilter{
					PartnerID: partnerID,
					PayoutID:  0,
					Status:    0,
				}).Return([]*domain.Payout{payout})
			},
		},
		{
			name: "get payouts by specific ID",
			args: args{
				req: GetAllPayoutsReq{
					PartnerID: 0,
					PayoutID:  payoutID,
					Status:    0,
				},
			},
			want: []*domain.Payout{{
				ID:           payoutID,
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       domain.PendingPayoutStatus,
				CreatedAt:    now,
				UpdateAt:     now,
			}},
			before: func(f useCaseMocks, args args) {
				payout := &domain.Payout{
					ID:           payoutID,
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.PendingPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     now,
				}

				f.repoPayout.EXPECT().GetAllPayouts(dto.PayoutFilter{
					PartnerID: 0,
					PayoutID:  payoutID,
					Status:    0,
				}).Return([]*domain.Payout{payout})
			},
		},
		{
			name: "get payouts by status",
			args: args{
				req: GetAllPayoutsReq{
					PartnerID: 0,
					PayoutID:  0,
					Status:    domain.PaidPayoutStatus,
				},
			},
			want: []*domain.Payout{{
				ID:           payoutID,
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       domain.PaidPayoutStatus,
				CreatedAt:    now,
				UpdateAt:     now,
			}},
			before: func(f useCaseMocks, args args) {
				payout := &domain.Payout{
					ID:           payoutID,
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.PaidPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     now,
				}

				f.repoPayout.EXPECT().GetAllPayouts(dto.PayoutFilter{
					PartnerID: 0,
					PayoutID:  0,
					Status:    domain.PaidPayoutStatus,
				}).Return([]*domain.Payout{payout})
			},
		},
		{
			name: "empty result",
			args: args{
				req: GetAllPayoutsReq{
					PartnerID: 999,
					PayoutID:  0,
					Status:    0,
				},
			},
			want: []*domain.Payout{},
			before: func(f useCaseMocks, args args) {
				f.repoPayout.EXPECT().GetAllPayouts(dto.PayoutFilter{
					PartnerID: 999,
					PayoutID:  0,
					Status:    0,
				}).Return([]*domain.Payout{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e := uc.GetAllPayouts(tt.args.req)

			a.Equal(tt.want, e)
		})
	}
}
