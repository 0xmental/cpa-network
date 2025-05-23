package main

import (
	"CPAPlatform/cmd/service_provider"
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/usecase/partner_usecase"
	"fmt"
	"log"
)

func main() {
	sp := service_provider.NewServiceProvider()
	puc := sp.GetPartnerUseCase()

	partner, err := puc.CreatePartner(partner_usecase.CreatePartnerRequest{
		Email: "test@example.com",
		Pass:  "securePassword123",
		ContactInfo: domain.ContactInfo{
			Skype:    "",
			Telegram: "@test_user",
			Discord:  "",
		},
		WithdrawInfo: nil,
		PostbackURL:  nil,
	})
	if err != nil {
		log.Fatalf("Failed to create partner: %v", err)
	}

	fmt.Printf("Partner created successfully:\n")
	fmt.Printf("ID: %d\n", partner.ID)
	fmt.Printf("Email: %s\n", partner.Email)
	fmt.Printf("Telegram: %s\n", partner.ContactInfo.Telegram)
	fmt.Printf("Balance: %d\n", partner.Balance)
	fmt.Printf("Is Active: %t\n", partner.IsActive)
	fmt.Printf("Created At: %s\n", partner.CreatedAt.Format("2006-01-02 15:04:05"))
}
