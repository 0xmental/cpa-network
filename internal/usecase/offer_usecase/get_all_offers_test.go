package offer_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllOffers(t *testing.T) {
	t.Parallel()

	now := time.Now()
	targetURL := "example.com"
	name := "Offer #1"
	description := "Test Offer"
	redirectDomain := "redirect.example.com"
	conversionType := domain.SOI
	payout := map[string]int64{"RU": 2}

	tests := []struct {
		name   string
		want   []*domain.Offer
		before func(ucMocks useCaseMocks)
	}{
		{
			name: "success get",
			want: []*domain.Offer{
				{
					TargetUrl:      targetURL,
					Name:           name,
					Description:    description,
					IsActive:       false,
					RedirectDomain: redirectDomain,
					ConversionType: conversionType,
					Payout:         payout,
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			before: func(f useCaseMocks) {
				offer := &domain.Offer{
					TargetUrl:      targetURL,
					Name:           name,
					Description:    description,
					IsActive:       false,
					RedirectDomain: redirectDomain,
					ConversionType: conversionType,
					Payout:         payout,
					CreatedAt:      now,
					UpdatedAt:      now,
				}

				f.repoOffer.EXPECT().GetAllOffers().Return([]*domain.Offer{offer})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks)

			e := uc.GetAllOffers()

			a.Equal(tt.want, e)
		})
	}
}
