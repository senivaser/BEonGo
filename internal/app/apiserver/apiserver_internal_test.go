package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/senivaser/BEonGo/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestAPIServer_handleHello(t *testing.T) {
	s := New(NewConfig(), &model.Store{})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
