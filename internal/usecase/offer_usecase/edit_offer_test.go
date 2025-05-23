package offer_usecase

import (
	"CPAPlatform/internal/adapter/repository/in_memory/offer_in_memory"
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEditOffer(t *testing.T) {
	t.Parallel()

	now := time.Now()
	updatedNow := now.Add(1 * time.Hour)
	var offerID int64 = 1

	type args struct {
		req UpdateOfferReq
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Offer
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success update",
			args: args{
				req: UpdateOfferReq{
					OfferID:        offerID,
					TargetURL:      "updated-example.com",
					Name:           "Updated Offer #1",
					Description:    "Updated Test Offer",
					RedirectDomain: "updated-redirect.example.com",
					ConversionType: domain.DOI,
					Payout:         map[string]int64{"RU": 3},
				},
			},
			want: &domain.Offer{
				ID:             offerID,
				TargetUrl:      "updated-example.com",
				Name:           "Updated Offer #1",
				Description:    "Updated Test Offer",
				IsActive:       false,
				RedirectDomain: "updated-redirect.example.com",
				ConversionType: domain.DOI,
				Payout:         map[string]int64{"RU": 3},
				CreatedAt:      now,
				UpdatedAt:      updatedNow,
			},
			wantErr: nil,
			before: func(f useCaseMocks, args args) {
				// Исходный оффер
				originalOffer := &domain.Offer{
					ID:             offerID,
					TargetUrl:      "example.com",
					Name:           "Offer #1",
					Description:    "Test Offer",
					IsActive:       false,
					RedirectDomain: "redirect.example.com",
					ConversionType: domain.SOI,
					Payout:         map[string]int64{"RU": 2},
					CreatedAt:      now,
					UpdatedAt:      now,
				}

				// Ожидаемый обновленный оффер
				updatedOffer := &domain.Offer{
					ID:             offerID,
					TargetUrl:      args.req.TargetURL,
					Name:           args.req.Name,
					Description:    args.req.Description,
					IsActive:       false,
					RedirectDomain: args.req.RedirectDomain,
					ConversionType: args.req.ConversionType,
					Payout:         args.req.Payout,
					CreatedAt:      now,
					UpdatedAt:      updatedNow,
				}

				// Настройка моков
				f.timer.EXPECT().Now().Return(updatedNow)
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(originalOffer, nil)
				f.repoOffer.EXPECT().Update(updatedOffer).Return(updatedOffer)
			},
		},
		{
			name: "offer not found",
			args: args{
				req: UpdateOfferReq{
					OfferID:        offerID,
					TargetURL:      "updated-example.com",
					Name:           "Updated Offer #1",
					Description:    "Updated Test Offer",
					RedirectDomain: "updated-redirect.example.com",
					ConversionType: domain.DOI,
					Payout:         map[string]int64{"RU": 3},
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

			e, err := uc.EditOffer(tt.args.req)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
