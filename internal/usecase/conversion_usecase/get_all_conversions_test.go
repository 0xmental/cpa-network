package conversion_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllConversions(t *testing.T) {
	var offerID int64 = 1
	var partnerID int64 = 1

	ipAddress := "123.1.1.12"
	userAgent := "ExampleUserAgent"
	clickID := domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent)
	now := time.Now()

	type args struct {
		req GetAllConversionsReq
	}

	tests := []struct {
		name   string
		args   args
		want   []*domain.Conversion
		before func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success get all conversions",
			args: args{
				req: GetAllConversionsReq{
					OfferID:   offerID,
					PartnerID: 0,
				},
			},
			want: []*domain.Conversion{
				{
					ClickID:   clickID,
					Payout:    2,
					OfferID:   offerID,
					PartnerID: partnerID,
					CreatedAt: now,
				},
			},
			before: func(f useCaseMocks, args args) {
				conversion := &domain.Conversion{
					ClickID:   clickID,
					Payout:    2,
					OfferID:   offerID,
					PartnerID: partnerID,
					CreatedAt: now,
				}

				f.repoConversion.EXPECT().GetAllConversions(dto.ConversionFilter{
					OfferID:   offerID,
					PartnerID: 0,
				}).Return([]*domain.Conversion{conversion})
			},
		},
		{
			name: "get conversions by partner ID",
			args: args{
				req: GetAllConversionsReq{
					OfferID:   0,
					PartnerID: partnerID,
				},
			},
			want: []*domain.Conversion{
				{
					ClickID:   clickID,
					Payout:    2,
					OfferID:   offerID,
					PartnerID: partnerID,
					CreatedAt: now,
				},
			},
			before: func(f useCaseMocks, args args) {
				conversion := &domain.Conversion{
					ClickID:   clickID,
					Payout:    2,
					OfferID:   offerID,
					PartnerID: partnerID,
					CreatedAt: now,
				}

				f.repoConversion.EXPECT().GetAllConversions(dto.ConversionFilter{
					OfferID:   0,
					PartnerID: partnerID,
				}).Return([]*domain.Conversion{conversion})
			},
		},
		{
			name: "get conversions by both offer and partner ID",
			args: args{
				req: GetAllConversionsReq{
					OfferID:   offerID,
					PartnerID: partnerID,
				},
			},
			want: []*domain.Conversion{
				{
					ClickID:   clickID,
					Payout:    2,
					OfferID:   offerID,
					PartnerID: partnerID,
					CreatedAt: now,
				},
			},
			before: func(f useCaseMocks, args args) {
				conversion := &domain.Conversion{
					ClickID:   clickID,
					Payout:    2,
					OfferID:   offerID,
					PartnerID: partnerID,
					CreatedAt: now,
				}

				f.repoConversion.EXPECT().GetAllConversions(dto.ConversionFilter{
					OfferID:   offerID,
					PartnerID: partnerID,
				}).Return([]*domain.Conversion{conversion})
			},
		},
		{
			name: "empty result",
			args: args{
				req: GetAllConversionsReq{
					OfferID:   999,
					PartnerID: 0,
				},
			},
			want: []*domain.Conversion{},
			before: func(f useCaseMocks, args args) {
				f.repoConversion.EXPECT().GetAllConversions(dto.ConversionFilter{
					OfferID:   999,
					PartnerID: 0,
				}).Return([]*domain.Conversion{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e := uc.GetAllConversions(tt.args.req)

			a.Equal(tt.want, e)
		})
	}
}
