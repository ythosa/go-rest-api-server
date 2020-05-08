package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Plane ...
type Plane struct {
	ID              int
	Name            string
	ManufactureYear int
}

// Validate ...
func (p *Plane) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.ManufactureYear, validation.Required),
	)
}
