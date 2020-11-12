package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestServer_New(t *testing.T) {
// 	s := New(util.NewConfig())
// 	rec := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/", nil)

// 	s.handleCounter().ServeHTTP(rec, req)

// 	assert.Equal(t, rec.Code, 200)
// }

// func TestServer_HandleCounter(t *testing.T) {
// 	s := New(util.NewConfig())
// 	rec := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/", nil)

// 	s.handleCounter().ServeHTTP(rec, req)

// 	assert.Equal(t, rec.Body)
// }

func TestServer_HandleCounter(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want []byte
	}{
		{
			name: "first",
			want: []byte("1"),
		},
		{
			name: "second",
			want: []byte("2"),
		},
		{
			name: "third",
			want: []byte("3"),
		},
	}

	handler := http.HandlerFunc(handleCounter)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)

			handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.want, rec.Body.Bytes())
		})
	}
}
