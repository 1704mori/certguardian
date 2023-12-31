package sslcheck

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	env "github.com/1704mori/certguardian/internal"
	certTypings "github.com/1704mori/certguardian/internal/certificates"
	"github.com/1704mori/certguardian/internal/domain"
	"github.com/1704mori/certguardian/internal/util"
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
	now := time.Now()
	isExpired := now.After(cert.NotAfter)

	ned, err := util.ConvertToDuration(env.Args.NearExpiryThreshold)
	if err != nil {
		log.Printf("could not convert NED to duration %v", err)
		return nil, err
	}

	info := &domain.Info{
		Issuer:       strings.Join(cert.Issuer.Organization[:], ","),
		ValidFrom:    cert.NotBefore,
		ValidTo:      cert.NotAfter,
		IsExpired:    isExpired,
		IsNearExpiry: cert.NotAfter.Sub(now) <= ned,
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

	now := time.Now()
	isExpired := now.After(cert.NotAfter)

	ned, err := util.ConvertToDuration(env.Args.NearExpiryThreshold)
	if err != nil {
		log.Printf("could not convert NED to duration %v", err)
		return nil, err
	}

	info := &domain.Info{
		Issuer:       strings.Join(cert.Issuer.Organization[:], ","),
		ValidFrom:    cert.NotBefore,
		ValidTo:      cert.NotAfter,
		IsExpired:    isExpired,
		IsNearExpiry: cert.NotAfter.Sub(now) <= ned,
	}
	return info, nil
}

func FindCertificates(directories []string) (certTypings.Info, error) {
	directoryCertificates := make(certTypings.DirectoryInfo)

	for _, dir := range directories {
		certificates, err := SearchDirectoryForCertificates(dir)
		if err != nil {
			return certTypings.Info{}, err
		}

		for _, cert := range certificates {
			info, err := FromPEM(cert)
			if err != nil {
				log.Println(err)
				continue
			}

			if directoryCertificates[dir] == nil {
				directoryCertificates[dir] = make(map[string]domain.Info)
			}

			directoryCertificates[dir][cert] = *info
		}
	}

	return certTypings.Info{
		Directories: directoryCertificates,
	}, nil
}

func SearchDirectoryForCertificates(dir string) ([]string, error) {
	var certificates []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), "fullchain.pem") {
			certificates = append(certificates, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return certificates, nil
}
