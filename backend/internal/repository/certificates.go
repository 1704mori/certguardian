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

func (repo *CertificateRepository) Add(info certificates.Info) error {
	jsonBytes, err := json.Marshal(info)
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

func (repo *CertificateRepository) List() ([]certificates.Info, error) {
	var domains []certificates.Info

	err := repo.DB.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Certificates"))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}

		return bucket.ForEach(func(k, v []byte) error {
			var domainMeta certificates.Info

			err := json.Unmarshal(v, &domainMeta)
			if err != nil {
				return err
			}

			domains = append(domains, domainMeta)

			return nil
		})
	})

	return domains, err
}
