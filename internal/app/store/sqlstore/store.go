package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/ythosa/go-rest-api-server/internal/app/store"
)

// Store ...
type Store struct {
	db              *sql.DB
	userRepository  *UserRepository
	planeRepository *PlaneRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Plane ...
func (s *Store) Plane() store.PlaneRepository {
	if s.planeRepository != nil {
		return s.planeRepository
	}

	s.planeRepository = &PlaneRepository{
		store: s,
	}

	return s.planeRepository
}
