package sqlstore

import (
	"github.com/ythosa/go-rest-api-server/internal/app/model"
	"github.com/ythosa/go-rest-api-server/internal/app/store"
)

// PlaneRepository ...
type PlaneRepository struct {
	store *Store
}

// Create ...
func (r *PlaneRepository) Create(p *model.Plane) error {
	if err := p.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO  planes (name, manufacture_year) VALUES ($1, $2) RETURNING id",
		p.Name,
		p.ManufactureYear,
	).Scan(&p.ID)
}

// Find ...
func (r *PlaneRepository) Find(id int) (*model.Plane, error) {
	p := &model.Plane{}

	if err := r.store.db.QueryRow(
		"SELECT id, name, manufacture_year FROM planes WHERE id = $1",
		id,
	).Scan(
		&p.ID,
		&p.Name,
		&p.ManufactureYear,
	); err != nil {
		return nil, store.ErrRecordNotFound
	}

	return p, nil
}
