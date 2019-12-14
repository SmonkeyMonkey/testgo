package api_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"test/internal/app/api"
	"test/internal/app/store/mongostore"
	"testing"
)

func TestServer_HandleGetAllUsers(t *testing.T) {
	db := mongostore.TestDB(t)
	s := api.NewServer(mongostore.New(db))
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Code)
}
func TestServer_HandleUserCreate(t *testing.T) {
	db := mongostore.TestDB(t)
	s := api.NewServer(mongostore.New(db))
	testCases := []struct {
		name string
		data interface{}
		code int
	}{
		{
			name: "valid",
			data: map[string]interface{}{
				"email":      "Logan_Devonport3313@tonsy.org",
				"last_name":  "Devonport",
				"country":    "Oman",
				"city":       "Madrid",
				"gender":     "Male",
				"birth_date": "Friday, April 4, 8527 8:45 AM",
			},
			code: http.StatusCreated,
		}, {
			name: "invalid date",
			data: "hello",
			code: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.data)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.code, rec.Code)
		})
	}
}
func TestServer_HandleGameCreate(t *testing.T) {
	db := mongostore.TestDB(t)
	s := api.NewServer(mongostore.New(db))
	testCases := []struct {
		name string
		data interface{}
		code int
	}{
		{
			name: "valid",
			data: map[string]interface{}{
				"points_gained": "677",
				"win_status":    "0",
				"game_type":     "11",
				"created":       "8/17/2019 8:54 PM",
			},
			code: http.StatusCreated,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.data)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/games", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.code, rec.Code)
		})
	}
}
