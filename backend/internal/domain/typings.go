package domain

import "time"

type Info struct {
	CommonName string    `json:"commonName" binding:"required"`
	Issuer     string    `json:"issuer"`
	ValidFrom  time.Time `json:"validFrom"`
	ValidTo    time.Time `json:"validTo"`
	IsExpired  bool      `json:"isExpired"`
}
