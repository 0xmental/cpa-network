package conversion_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllConversionsByPartner(t *testing.T) {
	var offerID int64 = 1
	var partnerID int64 = 1

	ipAddress := "123.1.1.12"
	userAgent := "ExampleUserAgent"
	clickID := domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent)
	now := time.Now()

	type args struct {
		req GetConversionsByPartnerReq
	}

	tests := []struct {
		name   string
		want   []*domain.Conversion
		args   args
		before func(ucMocks useCaseMocks)
	}{
		{
			name: "success get",
			args: args{
				req: GetConversionsByPartnerReq{
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
			before: func(f useCaseMocks) {
				conversion := &domain.Conversion{
					ClickID:   clickID,
					Payout:    2,
					OfferID:   offerID,
					PartnerID: partnerID,
					CreatedAt: now,
				}

				f.repoConversion.EXPECT().GetAllConversionsByPartner(partnerID).Return([]*domain.Conversion{conversion})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks)

			e := uc.GetAllConversionsByPartner(tt.args.req)

			a.Equal(tt.want, e)
		})
	}
}
