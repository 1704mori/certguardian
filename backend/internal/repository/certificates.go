package repository

import (
	"encoding/json"

	"github.com/1704mori/certguardian/internal/certificates"
	"github.com/1704mori/certguardian/internal/db"
	"go.etcd.io/bbolt"
)

type CertificateRepository struct {
	DB *db.Database
}

func NewCertificateRepository(database *db.Database) *CertificateRepository {
	return &CertificateRepository{DB: database}
}

func (repo *CertificateRepository) Add(certsDir []string) error {
	jsonBytes, err := json.Marshal(certsDir)
	if err != nil {
		return err
	}

	err = repo.DB.DB.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("Certificates"))
		if err != nil {
			return err
		}

		return bucket.Put([]byte("directories"), jsonBytes)
	})

	return err
}

func (repo *CertificateRepository) Find() (certificates.Info, error) {
	var certs certificates.Info

	err := repo.DB.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Certificates"))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}

		certsDirBytes := bucket.Get([]byte("directories"))
		if certsDirBytes == nil {
			return bbolt.ErrInvalid
		}

		return json.Unmarshal(certsDirBytes, &certs)
	})

	return certs, err
}

func (repo *CertificateRepository) FindByDirectories(directories []string) (certificates.Info, error) {
	var certs certificates.Info

	err := repo.DB.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Certificates"))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}

		certsDirBytes := bucket.Get([]byte("directories"))
		if certsDirBytes == nil {
			return bbolt.ErrInvalid
		}

		return json.Unmarshal(certsDirBytes, &certs)
	})

	return certs, err
}
