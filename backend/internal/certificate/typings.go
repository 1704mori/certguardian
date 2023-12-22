package certificate

import "time"

type Metadata struct {
	CommonName string    `json:"commonName" binding:"required"`
	Expiration time.Time `json:"expiration"`
	IsValid    bool      `json:"isValid"`
	Issuer     string    `json:"issuer"`
}
