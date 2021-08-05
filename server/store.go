package main

import (
	"github.com/zippoxer/bow"
)

// Store implements pricing storage using an embedded key-value database.
type Store struct {
	db *bow.DB
}

func OpenStore(path string) (*Store, error) {
	db, err := bow.Open("pricer.db")
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Put(token string, price float64) error {
	return s.db.Bucket("prices").Put(tokenPrice{Token: token, Price: price})
}

func (s *Store) Get(token string) (float64, error) {
	var v tokenPrice
	err := s.db.Bucket("prices").Get(token, &v)
	return v.Price, err
}

func (s *Store) Close() error {
	return s.db.Close()
}

type tokenPrice struct {
	Token string `bow:"key"`
	Price float64
}
