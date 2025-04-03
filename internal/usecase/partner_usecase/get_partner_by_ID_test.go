package partner_usecase

import (
	"CPAPlatform/internal/adapter/repository/partner_in_memory"
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetPartner(t *testing.T) {
	t.Parallel()

	var partnerID int64 = 1
	now := time.Now()
	email := "example@gmail.com"
	pass := "examplePassword"
	contactInfo := domain.ContactInfo{
		Skype:    "skype",
		Telegram: "",
		Discord:  "",
	}

	type args struct {
		reqGetPartner GetPartnerReq
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Partner
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success get",
			args: args{
				reqGetPartner: GetPartnerReq{
					PartnerID: partnerID,
				},
			},
			want: &domain.Partner{
				Email:        email,
				Pass:         pass,
				ContactInfo:  contactInfo,
				WithdrawInfo: nil,
				PostbackURL:  nil,
				IsActive:     true,
				Balance:      0,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			wantErr: nil,
			before: func(f useCaseMocks, args args) {
				partner := &domain.Partner{
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

				f.repoPartner.EXPECT().GetPartnerByID(partnerID).Return(partner, nil)
			},
		},
		{
			name: "partner not found",
			args: args{
				reqGetPartner: GetPartnerReq{
					PartnerID: partnerID,
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

			e, err := uc.GetPartnerByID(tt.args.reqGetPartner)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
