package storage

import (
	"post_service/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Post() postgres.PostI
}

type Store struct {
	db   *sqlx.DB
	post postgres.PostI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &Store{
		db:   db,
		post: postgres.NewPostRepo(db),
	}
}

func (s *Store) Post() postgres.PostI {
	if s.post == nil {
		s.post = postgres.NewPostRepo(s.db)
	}
	return s.post
}
