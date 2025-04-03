package click_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllClicksByOffer(t *testing.T) {
	var offerID int64 = 1
	var partnerID int64 = 1

	now := time.Now()
	utmParams := map[string]string{"sub1": "example", "sub2": "example", "sub3": "example"}
	ipAddress := "123.1.1.12"
	userAgent := "ExampleUserAgent"
	country := "RU"

	type args struct {
		req GetClicksByOfferReq
	}

	tests := []struct {
		name   string
		want   []*domain.Click
		args   args
		before func(ucMocks useCaseMocks)
	}{
		{
			name: "success get",
			args: args{
				req: GetClicksByOfferReq{
					OfferID: offerID,
				},
			},
			want: []*domain.Click{
				{
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
			},
			before: func(f useCaseMocks) {
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

				f.repoClick.EXPECT().GetAllClicksByOffer(offerID).Return([]*domain.Click{click})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks)

			e := uc.GetAllClicksByOffer(tt.args.req)

			a.Equal(tt.want, e)
		})
	}
}
