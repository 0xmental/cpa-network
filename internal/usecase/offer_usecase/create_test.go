package offer_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateOffer(t *testing.T) {
	t.Parallel()

	now := time.Now()
	targetURL := "example.com"
	name := "Offer #1"
	description := "Test Offer"
	redirectDomain := "redirect.example.com"
	conversionType := domain.SOI
	payout := map[string]int64{"RU": 2}

	type args struct {
		req CreateOfferRequest
	}

	tests := []struct {
		name   string
		args   args
		want   *domain.Offer
		before func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success creation",
			args: args{
				req: CreateOfferRequest{
					TargetURL:      targetURL,
					Name:           name,
					Description:    description,
					RedirectDomain: redirectDomain,
					ConversionType: conversionType,
					Payout:         payout,
				},
			},
			want: &domain.Offer{
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
			before: func(f useCaseMocks, args args) {
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

				f.timer.EXPECT().Now().Return(now)
				f.repoOffer.EXPECT().Save(offer).Return(offer)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e := uc.CreateOffer(tt.args.req)

			a.Equal(tt.want, e)
		})
	}
}
