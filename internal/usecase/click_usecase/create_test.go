package click_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateClick(t *testing.T) {
	var offerID int64 = 1
	var partnerID int64 = 1

	now := time.Now()
	utmParams := map[string]string{"sub1": "example", "sub2": "example", "sub3": "example"}
	ipAddress := "123.1.1.12"
	userAgent := "ExampleUserAgent"
	country := "RU"

	type args struct {
		req CreateClickRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Click
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success creation",
			args: args{
				req: CreateClickRequest{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent),
					UTMParams: utmParams,
					IPAddress: ipAddress,
					UserAgent: userAgent,
					Country:   country,
				},
			},
			want: &domain.Click{
				OfferID:   offerID,
				PartnerID: partnerID,
				ClickID:   domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent),
				UTMParams: utmParams,
				IPAddress: ipAddress,
				Useragent: userAgent,
				Country:   country,
				IsUnique:  true,
				CreatedAt: now,
			},
			before: func(f useCaseMocks, args args) {
				click := &domain.Click{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent),
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					IsUnique:  true,
					Country:   country,
					CreatedAt: now,
				}

				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(&domain.Partner{ID: partnerID}, nil)
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(&domain.Offer{ID: offerID}, nil)
				f.repoClick.EXPECT().IsUnique(args.req.ClickID).Return(true)
				f.repoClick.EXPECT().Save(click).Return(click)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.CreateClick(tt.args.req)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
