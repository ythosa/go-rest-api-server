package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

// TestPlane ...
func TestPlane(t *testing.T) *Plane {
	return &Plane{
		Name:            "GolangAirlines",
		ManufactureYear: 1997,
	}
}
