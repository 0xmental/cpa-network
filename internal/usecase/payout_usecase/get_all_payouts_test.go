package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllPayouts(t *testing.T) {
	t.Parallel()

	now := time.Now()
	var partnerID int64 = 1
	var amount int64 = 1000
	withdrawInfo := domain.WithdrawInfo{
		Method:     domain.BankWithdrawMethod,
		Requisites: "1234 5678 9012 3456",
	}

	tests := []struct {
		name   string
		want   []*domain.Payout
		before func(ucMocks useCaseMocks)
	}{
		{
			name: "successful get",
			want: []*domain.Payout{{
				PartnerID:    partnerID,
				WithdrawInfo: withdrawInfo,
				Amount:       amount,
				Status:       domain.PendingPayoutStatus,
				CreatedAt:    now,
				UpdateAt:     now,
			},
			},
			before: func(f useCaseMocks) {
				payout := &domain.Payout{
					PartnerID:    partnerID,
					WithdrawInfo: withdrawInfo,
					Amount:       amount,
					Status:       domain.PendingPayoutStatus,
					CreatedAt:    now,
					UpdateAt:     now,
				}

				f.repoPayout.EXPECT().GetAllPayouts().Return([]*domain.Payout{payout})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks)

			e := uc.GetAllPayouts()

			a.Equal(tt.want, e)
		})
	}
}
