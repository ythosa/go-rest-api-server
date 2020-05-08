package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ythosa/go-rest-api-server/internal/app/model"
)

func TestPlane_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		p       func() *model.Plane
		isValid bool
	}{
		{
			name: "valid",
			p: func() *model.Plane {
				return model.TestPlane(t)
			},
			isValid: true,
		},
		{
			name: "not valid all props",
			p: func() *model.Plane {
				p := model.TestPlane(t)
				p.Name = ""
				p.ManufactureYear = 0

				return p
			},
			isValid: false,
		},
		{
			name: "not valid year",
			p: func() *model.Plane {
				p := model.TestPlane(t)
				p.ManufactureYear = 990

				return p
			},
			isValid: false,
		},
		{
			name: "not valid name",
			p: func() *model.Plane {
				p := model.TestPlane(t)
				p.Name = ""

				return p
			},
			isValid: false,
		},
		{
			name: "not valid year",
			p: func() *model.Plane {
				p := model.TestPlane(t)
				p.ManufactureYear = 11111

				return p
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.p().Validate())
			} else {
				assert.Error(t, tc.p().Validate())
			}
		})
	}
}
