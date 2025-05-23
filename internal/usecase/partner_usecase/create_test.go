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
		wantErr bool
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
			wantErr: false,
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
				f.repoPartner.EXPECT().Save(partner).Return(partner, nil)
			},
		},
		{
			name: "creation failed - invalid email",
			args: args{
				req: CreatePartnerRequest{
					Email:        "",
					Pass:         "validpass",
					ContactInfo:  contactInfo,
					WithdrawInfo: nil,
					PostbackURL:  nil,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				f.timer.EXPECT().Now().Return(now)
			},
		},
		{
			name: "creation failed - empty contact info",
			args: args{
				req: CreatePartnerRequest{
					Email:        "valid@email.com",
					Pass:         "validpass",
					ContactInfo:  domain.ContactInfo{}, // пустая контактная информация
					WithdrawInfo: nil,
					PostbackURL:  nil,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				f.timer.EXPECT().Now().Return(now)
			},
		},
		{
			name: "creation failed - empty password",
			args: args{
				req: CreatePartnerRequest{
					Email:        "valid@email.com",
					Pass:         "",
					ContactInfo:  contactInfo,
					WithdrawInfo: nil,
					PostbackURL:  nil,
				},
			},
			want:    nil,
			wantErr: true,
			before: func(f useCaseMocks, args args) {
				f.timer.EXPECT().Now().Return(now)
			},
		},
		{
			name: "save failed - repository error",
			args: args{
				req: CreatePartnerRequest{
					Email:        email,
					Pass:         pass,
					ContactInfo:  contactInfo,
					WithdrawInfo: nil,
					PostbackURL:  nil,
				},
			},
			want:    nil,
			wantErr: true,
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
				f.repoPartner.EXPECT().Save(partner).Return(nil, errTest)
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
