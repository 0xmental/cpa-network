package domain

import "time"

type ConversionType int8

var SOI ConversionType = 1
var DOI ConversionType = 2
var Sale ConversionType = 3

type (
	Offer struct {
		ID             int64
		TargetUrl      string
		Name           string
		Description    string
		IsActive       bool
		RedirectDomain string
		ConversionType ConversionType
		Payout         map[string]int64
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}

	RedirectLink struct {
		Domain    string
		PartnerID int64
		OfferID   int64
	}
)

func NewOffer(
	targetURL, name, description string, redirectDomain string,
	conversionType ConversionType, payout map[string]int64, time time.Time) *Offer {
	return &Offer{
		TargetUrl:      targetURL,
		Name:           name,
		Description:    description,
		IsActive:       false,
		RedirectDomain: redirectDomain,
		ConversionType: conversionType,
		Payout:         payout,
		CreatedAt:      time,
		UpdatedAt:      time,
	}
}
