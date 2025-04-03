package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreatePartner(t *testing.T) {
	t.Parallel()

	now := time.Now()
	email := "example@gmail.com"
	pass := "examplePassword"
	contactInfo := domain.ContactInfo{
		Skype:    "skype",
		Telegram: "",
		Discord:  "",
	}

	type args struct {
		req CreatePartnerRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Partner
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success creation",
			args: args{
				req: CreatePartnerRequest{
					Email:        email,
					Pass:         pass,
					ContactInfo:  contactInfo,
					WithdrawInfo: nil,
					PostbackURL:  nil,
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
				f.timer.EXPECT().Now().Return(now)
				f.repoPartner.EXPECT().Save(partner).Return(partner)
			},
		},
		{
			name: "creation failed",
			args: args{
				req: CreatePartnerRequest{
					Email:        "",
					Pass:         "",
					ContactInfo:  domain.ContactInfo{},
					WithdrawInfo: nil,
					PostbackURL:  nil,
				},
			},
			want:    nil,
			wantErr: domain.ErrContactInfoRequired,
			before: func(f useCaseMocks, args args) {
				f.timer.EXPECT().Now().Return(now)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.CreatePartner(tt.args.req)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
