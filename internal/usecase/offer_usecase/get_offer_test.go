package offer_usecase

import (
	"CPAPlatform/internal/adapter/repository/offer_in_memory"
	"CPAPlatform/internal/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetOffer(t *testing.T) {
	t.Parallel()

	now := time.Now()
	targetURL := "example.com"
	name := "Offer #1"
	description := "Test Offer"
	redirectDomain := "redirect.example.com"
	conversionType := domain.SOI
	payout := map[string]int64{"RU": 2}
	var offerID int64 = 1
	var partnerID int64 = 1

	type args struct {
		reqGetOffer GetOfferReq
	}

	tests := []struct {
		name    string
		args    args
		want    *GetOfferResponse
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success get",
			args: args{
				reqGetOffer: GetOfferReq{
					OfferID:   offerID,
					PartnerID: partnerID,
				},
			},
			want: &GetOfferResponse{
				ID:             offerID,
				Name:           name,
				TargetURL:      targetURL,
				Description:    description,
				ConversionType: conversionType,
				Payout:         payout,
				TrackingURL:    fmt.Sprintf("https://%s/click?offer=%d&partner=%d", redirectDomain, offerID, partnerID),
			},
			wantErr: nil,
			before: func(f useCaseMocks, args args) {
				offer := &domain.Offer{
					ID:             offerID,
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

				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(offer, nil)
			},
		},
		{
			name: "offer not found",
			args: args{
				reqGetOffer: GetOfferReq{
					OfferID:   offerID,
					PartnerID: partnerID,
				},
			},
			want:    nil,
			wantErr: offer_in_memory.ErrOfferNotFound,
			before: func(f useCaseMocks, args args) {
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(nil, offer_in_memory.ErrOfferNotFound)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			// Вызываем метод GetOffer
			response, err := uc.GetOffer(tt.args.reqGetOffer)

			// Проверяем ошибку
			a.ErrorIs(err, tt.wantErr)

			// Проверяем результат
			a.Equal(tt.want, response)
		})
	}
}
