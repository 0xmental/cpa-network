package click_usecase

import (
	"CPAPlatform/internal/adapter/repository/click_in_memory"
	"CPAPlatform/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetClickByClickID(t *testing.T) {
	var offerID int64 = 1
	var partnerID int64 = 1

	now := time.Now()
	utmParams := map[string]string{"sub1": "example", "sub2": "example", "sub3": "example"}
	ipAddress := "123.1.1.12"
	userAgent := "ExampleUserAgent"
	country := "RU"

	type args struct {
		req GetClickByClickID
	}

	tests := []struct {
		name    string
		want    *domain.Click
		args    args
		wantErr error
		before  func(ucMocks useCaseMocks, args args)
	}{
		{
			name: "success get",
			args: args{
				req: GetClickByClickID{
					ClickID: domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent),
				},
			},
			want: &domain.Click{
				OfferID:   offerID,
				PartnerID: partnerID,
				ClickID:   domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent),
				UTMParams: utmParams,
				IPAddress: ipAddress,
				Useragent: userAgent,
				Country:   country,
				IsUnique:  true,
				CreatedAt: now,
			},
			before: func(f useCaseMocks, args args) {
				click := &domain.Click{
					OfferID:   offerID,
					PartnerID: partnerID,
					ClickID:   domain.GenerateClickID(offerID, partnerID, ipAddress, userAgent),
					UTMParams: utmParams,
					IPAddress: ipAddress,
					Useragent: userAgent,
					IsUnique:  true,
					Country:   country,
					CreatedAt: now,
				}

				f.repoClick.EXPECT().GetByClickID(click.ClickID).Return(click, nil)
			},
		},
		{
			name: "click not found",
			args: args{
				req: GetClickByClickID{
					ClickID: "12312",
				},
			},
			want:    nil,
			wantErr: click_in_memory.ErrClickIDNotFound,
			before: func(f useCaseMocks, args args) {
				f.repoClick.EXPECT().GetByClickID(args.req.ClickID).Return(nil, click_in_memory.ErrClickIDNotFound)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			uc, ucMocks := makeServiceWithMocks(t)
			tt.before(ucMocks, tt.args)

			e, err := uc.GetClickByClickID(tt.args.req)

			a.ErrorIs(err, tt.wantErr)

			a.Equal(tt.want, e)
		})
	}
}
