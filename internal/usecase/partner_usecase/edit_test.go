package partner_usecase

import (
	"CPAPlatform/internal/adapter/repository/partner_in_memory"
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEditPartner(t *testing.T) {
	t.Parallel()

	now := time.Now()
	updatedNow := now.Add(1 * time.Hour)
	email := "example@gmail.com"
	pass := "examplePassword"
	contactInfo := domain.ContactInfo{
		Skype:    "skype",
		Telegram: "",
		Discord:  "",
	}
	var partnerID int64 = 1
	withdrawInfo := domain.WithdrawInfo{
		Method:     domain.USDTWithdrawMethod,
		Requisites: "test-requisites",
	}
	type args struct {
		req UpdateInfoReq
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Partner
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success update",
			args: args{
				req: UpdateInfoReq{
					PartnerID: partnerID,
					Pass:      "updated-pass",
					ContactInfo: domain.ContactInfo{
						Skype:    "skype-update",
						Telegram: "",
						Discord:  "",
					},
					WithdrawInfo: &withdrawInfo,
					PostbackURL:  nil,
				},
			},
			want: &domain.Partner{
				ID:    partnerID,
				Email: email,
				Pass:  "updated-pass",
				ContactInfo: domain.ContactInfo{
					Skype:    "skype-update",
					Telegram: "",
					Discord:  "",
				},
				WithdrawInfo: &withdrawInfo,
				PostbackURL:  nil,
				IsActive:     true,
				Balance:      0,
				CreatedAt:    now,
				UpdatedAt:    updatedNow,
			},
			wantErr: nil,
			before: func(f useCaseMocks, args args) {
				originalPartner := &domain.Partner{
					ID:           partnerID,
					Email:        email,
					Pass:         pass,
					ContactInfo:  contactInfo,
					WithdrawInfo: nil,
					PostbackURL:  nil,
					IsActive:     true,
					Balance:      0,
					CreatedAt:    now,
					UpdatedAt:    now,
				}

				updatedPartner := &domain.Partner{
					ID:           partnerID,
					Email:        email,
					Pass:         args.req.Pass,
					ContactInfo:  args.req.ContactInfo,
					WithdrawInfo: args.req.WithdrawInfo,
					PostbackURL:  nil,
					IsActive:     true,
					Balance:      0,
					CreatedAt:    now,
					UpdatedAt:    updatedNow,
				}

				f.timer.EXPECT().Now().Return(updatedNow)
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(originalPartner, nil)
				f.repoPartner.EXPECT().Update(updatedPartner).Return(updatedPartner)
			},
		},
		{
			name: "partner not found",
			args: args{
				req: UpdateInfoReq{
					PartnerID: partnerID,
					Pass:      "updated-pass",
					ContactInfo: domain.ContactInfo{
						Skype:    "skype-update",
						Telegram: "",
						Discord:  "",
					},
					WithdrawInfo: &withdrawInfo,
					PostbackURL:  nil,
				},
			},
			want:    nil,
			wantErr: partner_in_memory.ErrPartnerNotFound,
			before: func(f useCaseMocks, args args) {
				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(nil, partner_in_memory.ErrPartnerNotFound)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.UpdatePartnerInfo(tt.args.req)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
