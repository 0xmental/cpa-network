package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUpdatePayoutStatus(t *testing.T) {
	t.Parallel()

	now := time.Now()
	updatedNow := now.Add(1 * time.Hour)
	var partnerID int64 = 1
	var payoutID int64 = 1
	var amount int64 = 1000
	withdrawInfo := domain.WithdrawInfo{
		Method:     domain.BankWithdrawMethod,
		Requisites: "1234 5678 9012 3456",
	}
	statusOrigin := domain.PendingPayoutStatus
	statusUpdate := domain.PaidPayoutStatus

	type args struct {
		req UpdatePayoutReq
	}

	tests := []struct {
		name    string
		want    *domain.Payout
		wantErr bool
		args    args
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "successful update",
			args: args{
				req: UpdatePayoutReq{
					PayoutID: payoutID,
					Status:   statusUpdate,
				},
			},
			want: &domain.Payout{
				ID:           payoutID,
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       statusUpdate,
				CreatedAt:    now,
				UpdateAt:     updatedNow,
			},
			wantErr: false,
			before: func(f useCaseMocks, args args) {
				originalPayout := &domain.Payout{
					ID:           payoutID,
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       statusOrigin,
					CreatedAt:    now,
					UpdateAt:     now,
				}
				updatedPayout := &domain.Payout{
					ID:           payoutID,
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       statusUpdate,
					CreatedAt:    now,
					UpdateAt:     updatedNow,
				}

				f.timer.EXPECT().Now().Return(updatedNow)
				f.repoPayout.EXPECT().GetAllPayouts(dto.PayoutFilter{
					PayoutID: payoutID,
				}).Return([]*domain.Payout{originalPayout})
				f.repoPayout.EXPECT().UpdatePayoutStatus(updatedPayout).Return(updatedPayout)
			},
		},
		{
			name: "payout not found",
			args: args{
				req: UpdatePayoutReq{
					PayoutID: payoutID,
					Status:   statusUpdate,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				f.repoPayout.EXPECT().GetAllPayouts(dto.PayoutFilter{
					PayoutID: payoutID,
				}).Return([]*domain.Payout{})
			},
		},
		{
			name: "update to canceled status",
			args: args{
				req: UpdatePayoutReq{
					PayoutID: payoutID,
					Status:   domain.CanceledPayoutStatus,
				},
			},
			want: &domain.Payout{
				ID:           payoutID,
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       domain.CanceledPayoutStatus,
				CreatedAt:    now,
				UpdateAt:     updatedNow,
			},
			wantErr: false,
			before: func(f useCaseMocks, args args) {
				originalPayout := &domain.Payout{
					ID:           payoutID,
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       statusOrigin,
					CreatedAt:    now,
					UpdateAt:     now,
				}
				updatedPayout := &domain.Payout{
					ID:           payoutID,
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.CanceledPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     updatedNow,
				}

				f.timer.EXPECT().Now().Return(updatedNow)
				f.repoPayout.EXPECT().GetAllPayouts(dto.PayoutFilter{
					PayoutID: payoutID,
				}).Return([]*domain.Payout{originalPayout})
				f.repoPayout.EXPECT().UpdatePayoutStatus(updatedPayout).Return(updatedPayout)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.UpdatePayoutStatus(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				a.Nil(e)
			} else {
				a.NoError(err)
				a.Equal(tt.want, e)
			}
		})
	}
}
