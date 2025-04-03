package payout_usecase

import (
	"CPAPlatform/internal/adapter/repository/payout_in_memory"
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetPayoutByID(t *testing.T) {
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
		req GetPayoutReq
	}

	tests := []struct {
		name    string
		want    *domain.Payout
		wantErr error
		args    args
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "successful get",
			args: args{
				req: GetPayoutReq{
					PayoutID: payoutID,
				},
			},
			want: &domain.Payout{
				ID:           payoutID,
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       domain.PendingPayoutStatus,
				CreatedAt:    now,
				UpdateAt:     now,
			},
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

				f.repoPayout.EXPECT().GetPayoutByID(payoutID).Return(payout, nil)
			},
		},
		{
			name: "payout not found",
			args: args{
				req: GetPayoutReq{
					PayoutID: payoutID,
				},
			},
			want:    nil,
			wantErr: payout_in_memory.ErrPayoutNotFound,
			before: func(f useCaseMocks, args args) {
				f.repoPayout.EXPECT().GetPayoutByID(payoutID).Return(nil, payout_in_memory.ErrPayoutNotFound)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.GetPayoutByID(tt.args.req)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
