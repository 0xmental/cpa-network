package click_usecase

import (
	"CPAPlatform/internal/adapter/repository/in_memory/offer_in_memory"
	"CPAPlatform/internal/adapter/repository/in_memory/partner_in_memory"
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
	clickID := domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent)

	type args struct {
		req CreateClickRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Click
		wantErr bool
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success creation with provided clickID",
			args: args{
				req: CreateClickRequest{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					UserAgent: userAgent,
					Country:   country,
				},
			},
			want: &domain.Click{
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
			wantErr: false,
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

				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(&domain.Partner{ID: partnerID}, nil)
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(&domain.Offer{ID: offerID}, nil)
				f.repoClick.EXPECT().IsUnique(args.req.ClickID).Return(true)
				f.repoClick.EXPECT().Save(click).Return(click)
			},
		},
		{
			name: "success creation with generated clickID",
			args: args{
				req: CreateClickRequest{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   "", // Пустой clickID, должен генерироваться автоматически
					UTMParams: utmParams,
					IPAddress: ipAddress,
					UserAgent: userAgent,
					Country:   country,
				},
			},
			want: &domain.Click{
				OfferID:   offerID,
				PartnerID: partnerID,
				ClickID:   clickID, // Сгенерированный clickID
				UTMParams: utmParams,
				IPAddress: ipAddress,
				Useragent: userAgent,
				Country:   country,
				IsUnique:  true,
				CreatedAt: now,
			},
			wantErr: false,
			before: func(f useCaseMocks, args args) {
				partner := &domain.Partner{ID: partnerID}
				offer := &domain.Offer{ID: offerID}

				click := &domain.Click{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID, // Сгенерированный
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					IsUnique:  true,
					Country:   country,
					CreatedAt: now,
				}

				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(partner, nil)
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(offer, nil)
				f.repoClick.EXPECT().IsUnique("").Return(true) // Проверяем пустой ClickID из запроса
				f.repoClick.EXPECT().Save(click).Return(click)
			},
		},
		{
			name: "partner not found error",
			args: args{
				req: CreateClickRequest{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					UserAgent: userAgent,
					Country:   country,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(nil, partner_in_memory.ErrPartnerNotFound)
			},
		},
		{
			name: "offer not found error",
			args: args{
				req: CreateClickRequest{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					UserAgent: userAgent,
					Country:   country,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(&domain.Partner{ID: partnerID}, nil)
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(nil, offer_in_memory.ErrOfferNotFound)
			},
		},
		{
			name: "non-unique click",
			args: args{
				req: CreateClickRequest{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					UserAgent: userAgent,
					Country:   country,
				},
			},
			want: &domain.Click{
				OfferID:   offerID,
				PartnerID: partnerID,
				ClickID:   clickID,
				UTMParams: utmParams,
				IPAddress: ipAddress,
				Useragent: userAgent,
				Country:   country,
				IsUnique:  false, // Не уникальный клик
				CreatedAt: now,
			},
			wantErr: false,
			before: func(f useCaseMocks, args args) {
				click := &domain.Click{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   clickID,
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					IsUnique:  false,
					Country:   country,
					CreatedAt: now,
				}

				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(&domain.Partner{ID: partnerID}, nil)
				f.repoOffer.EXPECT().GetOfferByID(offerID).Return(&domain.Offer{ID: offerID}, nil)
				f.repoClick.EXPECT().IsUnique(args.req.ClickID).Return(false) // Не уникальный
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

			if tt.wantErr {
				a.Error(err)
				a.Nil(e)
			} else {
				a.NoError(err)
				a.Equal(tt.want, e)
			}
		})
	}
}
