package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllPartners(t *testing.T) {
	t.Parallel()

	now := time.Now()
	email := "example@gmail.com"
	pass := "examplePassword"
	contactInfo := domain.ContactInfo{
		Skype:    "skype",
		Telegram: "",
		Discord:  "",
	}

	tests := []struct {
		name   string
		want   []*domain.Partner
		before func(ucMocks useCaseMocks)
	}{
		{
			name: "success get",
			want: []*domain.Partner{
				{
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
			},
			before: func(f useCaseMocks) {
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

				f.repoPartner.EXPECT().GetAllPartners().Return([]*domain.Partner{partner})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks)

			e := uc.GetAllPartners()

			a.Equal(tt.want, e)
		})
	}
}
