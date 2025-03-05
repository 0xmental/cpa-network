package domain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Click struct {
	ID        int64
	OfferID   int64
	PartnerID int64
	ClickID   string // Если не передан, генерируется автоматически
	UTMParams map[string]string
	IPAddress string
	Useragent string
	Country   string
	IsUnique  bool
	CreatedAt time.Time
}

func GenerateClickID(offerID, partnerID int64, ip, userAgent string) string {
	data := fmt.Sprintf("%d-%d-%s-%s", offerID, partnerID, ip, userAgent)
	hash := sha256.Sum256([]byte(data))

	return fmt.Sprintf("%x", hash[:16])
}

func NewClick(offerID, partnerID int64, clickID, country, ipAddress, useragent string,
	utmParams map[string]string, isUnique bool, time time.Time) *Click {
	return &Click{
		OfferID:   offerID,
		PartnerID: partnerID,
		ClickID:   clickID,
		UTMParams: utmParams,
		IPAddress: ipAddress,
		Useragent: useragent,
		Country:   country,
		IsUnique:  isUnique,
		CreatedAt: time,
	}
}
