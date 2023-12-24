package repository

import (
	"encoding/json"

	"github.com/1704mori/certguardian/internal/db"
	"github.com/1704mori/certguardian/internal/domain"
	"go.etcd.io/bbolt"
)

type DomainRepository struct {
	DB *db.Database
}

func NewDomainRepository(database *db.Database) *DomainRepository {
	return &DomainRepository{DB: database}
}

func (repo *DomainRepository) Add(domain string, domainMeta domain.Info) error {
	jsonBytes, err := json.Marshal(domainMeta)
	if err != nil {
		return err
	}

	err = repo.DB.DB.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("Domains"))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(domain), jsonBytes)
	})

	return err
}

func (repo *DomainRepository) Find(_domain string) (*domain.Info, error) {
	var domainMeta domain.Info

	err := repo.DB.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Domains"))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}

		domainMetaBytes := bucket.Get([]byte(_domain))
		if domainMetaBytes == nil {
			return bbolt.ErrBucketNotFound
		}

		return json.Unmarshal(domainMetaBytes, &domainMeta)
	})

	return &domainMeta, err
}

func (repo *DomainRepository) List() ([]domain.Info, error) {
	var domains []domain.Info

	err := repo.DB.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Domains"))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}

		return bucket.ForEach(func(k, v []byte) error {
			var domainMeta domain.Info

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

func (repo *DomainRepository) Delete(domain string) error {
	err := repo.DB.DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Domains"))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}

		if bucket.Get([]byte(domain)) == nil {
			return bbolt.ErrBucketNotFound
		}

		return bucket.Delete([]byte(domain))
	})

	return err
}
