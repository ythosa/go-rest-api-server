package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ythosa/go-rest-api-server/internal/app/store/teststore"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	s := newServer(teststore.New())

	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
