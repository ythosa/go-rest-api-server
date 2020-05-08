package store

import "github.com/ythosa/go-rest-api-server/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}

// PlaneRepository ...
type PlaneRepository interface {
	Create(*model.Plane) error
	Find(int) (*model.Plane, error)
}
