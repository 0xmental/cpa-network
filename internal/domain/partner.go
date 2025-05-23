package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/mail"
	"strings"
	"time"
)

var ErrContactInfoRequired = errors.New("at least one contact method (Skype, Telegram, or Discord) must be provided")
var ErrIncorrectAmount = errors.New("amount must be greater than zero")
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrInvalidEmail = errors.New("invalid email address")
var ErrEmptyPassword = errors.New("password cannot be empty")

type WithdrawMethod int8

var USDTWithdrawMethod WithdrawMethod = 1
var BankWithdrawMethod WithdrawMethod = 2

type (
	Partner struct {
		ID           int64
		Email        string
		Pass         string
		ContactInfo  ContactInfo
		WithdrawInfo *WithdrawInfo
		PostbackURL  *string
		IsActive     bool
		Balance      int64
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	ContactInfo struct {
		Skype    string
		Telegram string
		Discord  string
	}
	WithdrawInfo struct {
		Method     WithdrawMethod
		Requisites string
	}
)

func (a *ContactInfo) Value() (driver.Value, error) { return json.Marshal(a) }

func (a *ContactInfo) Scan(value interface{}) error {
	b, ok := value.(string)
	if !ok {
		return errors.New("type assertion failed on ContactInfo")
	}

	return json.Unmarshal([]byte(b), &a)
}

func (a *WithdrawInfo) Value() (driver.Value, error) { return json.Marshal(a) }

func (a *WithdrawInfo) Scan(value interface{}) error {
	b, ok := value.(string)
	if !ok {
		return errors.New("type assertion failed on WithdrawInfo")
	}

	return json.Unmarshal([]byte(b), &a)
}

func NewPartner(email, pass string, contactInfo ContactInfo, withdrawInfo *WithdrawInfo, postbackURL *string, time time.Time) (*Partner, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, ErrInvalidEmail
	}

	if strings.TrimSpace(pass) == "" {
		return nil, ErrEmptyPassword
	}

	if contactInfo.Skype == "" && contactInfo.Telegram == "" && contactInfo.Discord == "" {
		return nil, ErrContactInfoRequired
	}

	return &Partner{
		Email:        email,
		Pass:         pass,
		ContactInfo:  contactInfo,
		WithdrawInfo: withdrawInfo,
		PostbackURL:  postbackURL,
		IsActive:     true,
		Balance:      0,
		CreatedAt:    time,
		UpdatedAt:    time,
	}, nil
}

func (p *Partner) Activate(activatedAt time.Time) {
	p.IsActive = true
	p.UpdatedAt = activatedAt
}

func (p *Partner) DeductBalance(amount int64) error {
	if amount <= 0 {
		return ErrIncorrectAmount
	}

	if p.Balance < amount {
		return ErrInsufficientBalance
	}
	p.Balance -= amount

	return nil
}

func (p *Partner) AddBalance(amount int64) {
	p.Balance += amount
}
