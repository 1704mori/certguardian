package db

import (
	"go.etcd.io/bbolt"
)

type Database struct {
	DB *bbolt.DB
}

func New(dbPath string) (*Database, error) {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (bdb *Database) Close() error {
	if bdb.DB != nil {
		return bdb.DB.Close()
	}
	return nil
}
