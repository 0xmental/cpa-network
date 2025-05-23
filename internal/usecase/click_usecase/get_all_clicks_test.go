package click_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllClicks(t *testing.T) {
	var offerID int64 = 1
	var partnerID int64 = 1

	now := time.Now()
	utmParams := map[string]string{"sub1": "example", "sub2": "example", "sub3": "example"}
	ipAddress := "123.1.1.12"
	userAgent := "ExampleUserAgent"
	country := "RU"
	clickID := domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent)

	type args struct {
		req GetAllClicksReq
	}

	tests := []struct {
		name   string
		args   args
		want   []*domain.Click
		before func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success get all clicks",
			args: args{
				req: GetAllClicksReq{
					PartnerID: partnerID,
					OfferID:   0,
					ClickID:   "",
				},
			},
			want: []*domain.Click{
				{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					Country:   country,
					IsUnique:  true,
					CreatedAt: now,
				},
			},
			before: func(f useCaseMocks, args args) {
				click := &domain.Click{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					IsUnique:  true,
					Country:   country,
					CreatedAt: now,
				}

				f.repoClick.EXPECT().GetAllClicks(dto.ClickFilter{
					PartnerID: partnerID,
					OfferID:   0,
					ClickID:   "",
				}).Return([]*domain.Click{click})
			},
		},
		{
			name: "get clicks by offer ID",
			args: args{
				req: GetAllClicksReq{
					PartnerID: 0,
					OfferID:   offerID,
					ClickID:   "",
				},
			},
			want: []*domain.Click{
				{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					Country:   country,
					IsUnique:  true,
					CreatedAt: now,
				},
			},
			before: func(f useCaseMocks, args args) {
				click := &domain.Click{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					IsUnique:  true,
					Country:   country,
					CreatedAt: now,
				}

				f.repoClick.EXPECT().GetAllClicks(dto.ClickFilter{
					PartnerID: 0,
					OfferID:   offerID,
					ClickID:   "",
				}).Return([]*domain.Click{click})
			},
		},
		{
			name: "get clicks by click ID",
			args: args{
				req: GetAllClicksReq{
					PartnerID: 0,
					OfferID:   0,
					ClickID:   clickID,
				},
			},
			want: []*domain.Click{
				{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					Country:   country,
					IsUnique:  true,
					CreatedAt: now,
				},
			},
			before: func(f useCaseMocks, args args) {
				click := &domain.Click{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					IsUnique:  true,
					Country:   country,
					CreatedAt: now,
				}

				f.repoClick.EXPECT().GetAllClicks(dto.ClickFilter{
					PartnerID: 0,
					OfferID:   0,
					ClickID:   clickID,
				}).Return([]*domain.Click{click})
			},
		},
		{
			name: "empty result",
			args: args{
				req: GetAllClicksReq{
					PartnerID: 999,
					OfferID:   0,
					ClickID:   "",
				},
			},
			want: []*domain.Click{},
			before: func(f useCaseMocks, args args) {
				f.repoClick.EXPECT().GetAllClicks(dto.ClickFilter{
					PartnerID: 999,
					OfferID:   0,
					ClickID:   "",
				}).Return([]*domain.Click{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e := uc.GetAllClicks(tt.args.req)

			a.Equal(tt.want, e)
		})
	}
}
