package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllPayoutsByPartner(t *testing.T) {
	t.Parallel()

	now := time.Now()
	var partnerID int64 = 1
	var amount int64 = 1000
	withdrawInfo := domain.WithdrawInfo{
		Method:     domain.BankWithdrawMethod,
		Requisites: "1234 5678 9012 3456",
	}

	type args struct {
		req GetPayoutsByPartnerReq
	}

	tests := []struct {
		name   string
		want   []*domain.Payout
		args   args
		before func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "successful get",
			args: args{
				req: GetPayoutsByPartnerReq{
					PartnerID: partnerID,
				},
			},
			want: []*domain.Payout{
				{
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.PendingPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     now,
				},
			},
			before: func(f useCaseMocks, args args) {
				payout := &domain.Payout{
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.PendingPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     now,
				}

				f.repoPayout.EXPECT().GetAllPayoutsByPartnerID(partnerID).Return([]*domain.Payout{payout})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e := uc.GetAllPayoutsByPartnerID(tt.args.req)

			a.Equal(tt.want, e)
		})
	}
}
