package partner_postgres

import (
	"CPAPlatform/internal/domain"
	"context"
	"fmt"
)

func (r *Repo) Save(partner *domain.Partner) (*domain.Partner, error) {
	err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO partners (email, pass, contact_info, withdraw_info, postback_url, is_active, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;",
		partner.Email, partner.Pass, partner.ContactInfo, partner.WithdrawInfo, partner.PostbackURL,
		partner.IsActive, partner.Balance, partner.CreatedAt, partner.UpdatedAt).Scan(&partner.ID)

	if err != nil {
		return nil, fmt.Errorf("Conn.QueryRow: %w", err)
	}

	return partner, nil
}
