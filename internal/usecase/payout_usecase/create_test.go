package payout_usecase

import (
	"CPAPlatform/internal/adapter/repository/in_memory/partner_in_memory"
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreatePayout(t *testing.T) {
	t.Parallel()

	now := time.Now()
	var partnerID int64 = 1
	var amount int64 = 1000
	withdrawInfo := domain.WithdrawInfo{
		Method:     domain.BankWithdrawMethod,
		Requisites: "1234 5678 9012 3456",
	}

	type args struct {
		req CreatePayoutRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Payout
		wantErr bool
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "successful payout creation",
			args: args{
				req: CreatePayoutRequest{
					PartnerID:    partnerID,
					Amount:       amount,
					WithdrawInfo: withdrawInfo,
				},
			},
			want: &domain.Payout{
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       domain.PendingPayoutStatus,
				CreatedAt:    now,
				UpdateAt:     now,
			},
			wantErr: false,
			before: func(f useCaseMocks, args args) {
				partner := &domain.Partner{
					ID:           partnerID,
					Balance:      5000,
					WithdrawInfo: &withdrawInfo,
				}
				payout := &domain.Payout{
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.PendingPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     now,
				}

				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(partner, nil)
				f.repoPayout.EXPECT().Save(payout).Return(payout)
			},
		},
		{
			name: "partner not found",
			args: args{
				req: CreatePayoutRequest{
					PartnerID:    partnerID,
					Amount:       amount,
					WithdrawInfo: withdrawInfo,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(nil, partner_in_memory.ErrPartnerNotFound)
			},
		},
		{
			name: "insufficient balance",
			args: args{
				req: CreatePayoutRequest{
					PartnerID:    partnerID,
					Amount:       amount,
					WithdrawInfo: withdrawInfo,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				partner := &domain.Partner{
					ID:           partnerID,
					Balance:      500, // меньше чем amount (1000)
					WithdrawInfo: &withdrawInfo,
				}

				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(partner, nil)
			},
		},
		{
			name: "zero amount",
			args: args{
				req: CreatePayoutRequest{
					PartnerID:    partnerID,
					Amount:       0,
					WithdrawInfo: withdrawInfo,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				partner := &domain.Partner{
					ID:           partnerID,
					Balance:      5000,
					WithdrawInfo: &withdrawInfo,
				}

				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(partner, nil)
			},
		},
		{
			name: "negative amount",
			args: args{
				req: CreatePayoutRequest{
					PartnerID:    partnerID,
					Amount:       -100,
					WithdrawInfo: withdrawInfo,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				partner := &domain.Partner{
					ID:           partnerID,
					Balance:      5000,
					WithdrawInfo: &withdrawInfo,
				}

				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(partner, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.CreatePayout(tt.args.req)

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
