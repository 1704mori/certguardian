package repository

import (
	"encoding/json"

	"github.com/1704mori/certguardian/internal/certificate"
	"github.com/1704mori/certguardian/internal/db"
	"go.etcd.io/bbolt"
)

type CertificateRepository struct {
	DB *db.Database
}

func NewCertificateRepository(database *db.Database) *CertificateRepository {
	return &CertificateRepository{DB: database}
}

func (repo *CertificateRepository) Add(domain string, certMeta certificate.Metadata) error {
	jsonBytes, err := json.Marshal(certMeta)
	if err != nil {
		return err
	}

	err = repo.DB.DB.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("Certificates"))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(domain), jsonBytes)
	})

	return err
}

func (repo *CertificateRepository) Find(domain string) (*certificate.Metadata, error) {
	var certMeta certificate.Metadata

	// Retrieve the certificate from the database.
	err := repo.DB.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Certificates"))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}

		certMetaBytes := bucket.Get([]byte(domain))
		if certMetaBytes == nil {
			return bbolt.ErrInvalid
		}

		return json.Unmarshal(certMetaBytes, &certMeta)
	})

	return &certMeta, err
}
