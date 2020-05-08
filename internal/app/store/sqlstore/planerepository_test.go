package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ythosa/go-rest-api-server/internal/app/model"
	"github.com/ythosa/go-rest-api-server/internal/app/store/sqlstore"
)

func TestPlaneRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("planes")

	s := sqlstore.New(db)
	p := model.TestPlane(t)

	assert.NoError(t, s.Plane().Create(p))
	assert.NotNil(t, p.ID)
}

func TestPlaneRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("planes")

	s := sqlstore.New(db)
	p1 := model.TestPlane(t)
	s.Plane().Create(p1)
	p2, err := s.Plane().Find(p1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, p2)
}
