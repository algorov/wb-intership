package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_HandleIndex(t *testing.T) {
	s := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	s.handleIndex().ServeHTTP(recorder, request)
	assert.Equal(t, "Unable to read HTML file\n", recorder.Body.String())
}
