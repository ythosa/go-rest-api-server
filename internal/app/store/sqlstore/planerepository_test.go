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
