package conversion_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateConversion(t *testing.T) {
	t.Parallel()

	var offerID int64 = 1
	var partnerID int64 = 1
	now := time.Now()
	ipAddress := "123.1.1.12"
	userAgent := "ExampleUserAgent"
	country := "RU"
	payout := map[string]int64{"RU": 2}
	clickID := domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent)

	type args struct {
		req CreateConversionRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Conversion
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{name: "success creation",
			args: args{
				req: CreateConversionRequest{
					ClickID: clickID,
				},
			},
			want: &domain.Conversion{
				ClickID:   clickID,
				Payout:    2,
				OfferID:   offerID,
				PartnerID: partnerID,
				CreatedAt: now,
			},
			before: func(f useCaseMocks, args args) {
				conversion := &domain.Conversion{
					ClickID:   clickID,
					OfferID:   offerID,
					Payout:    2,
					PartnerID: partnerID,
					CreatedAt: now,
				}
				click := &domain.Click{
					ClickID:   clickID,
					OfferID:   offerID,
					Country:   country,
					PartnerID: partnerID,
				}
				offer := &domain.Offer{
					ID:       offerID,
					Payout:   payout,
					IsActive: true,
				}
				partner := &domain.Partner{
					ID:      partnerID,
					Balance: 0,
				}

				f.timer.EXPECT().Now().Return(now)

				f.repoClick.EXPECT().GetByClickID(clickID).Return(click, nil)
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(offer, nil)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(partner, nil)
				f.repoConversion.EXPECT().Save(conversion).Return(conversion)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.CreateConversion(tt.args.req)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
