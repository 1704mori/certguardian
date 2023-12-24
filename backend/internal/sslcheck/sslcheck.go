package sslcheck

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"os"
	"strings"
	"time"

	"github.com/1704mori/certguardian/internal/domain"
)

func FromDomain(_domain string) (*domain.Info, error) {
	conn, err := tls.Dial("tcp", _domain+":443", nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return nil, err
	}

	cert := certs[0]
	info := &domain.Info{
		Issuer:    strings.Join(cert.Issuer.Organization[:], ","),
		ValidFrom: cert.NotBefore,
		ValidTo:   cert.NotAfter,
		IsExpired: time.Now().After(cert.NotAfter),
	}
	return info, nil
}

func FromPEM(filename string) (*domain.Info, error) {
	certPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, err
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	info := &domain.Info{
		Issuer:    strings.Join(cert.Issuer.Organization[:], ","),
		ValidFrom: cert.NotBefore,
		ValidTo:   cert.NotAfter,
		IsExpired: time.Now().After(cert.NotAfter),
	}
	return info, nil
}
