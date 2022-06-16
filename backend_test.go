package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router, _ := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/testPing", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuth(t *testing.T) {
	router, _ := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/helloAuth", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

var accessToken = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImY5MGZiMWFlMDQ4YTU0OGZiNjgxYWQ2MDkyYjBiODY5ZWE0NjdhYzYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbWRzLXByb2plY3QtMzQ5MzE2IiwiYXVkIjoibWRzLXByb2plY3QtMzQ5MzE2IiwiYXV0aF90aW1lIjoxNjU1MzM5NTA1LCJ1c2VyX2lkIjoiZkU1bjNNRVE1bldZTUV1NVpSeFdrcHhlWE9DMiIsInN1YiI6ImZFNW4zTUVRNW5XWU1FdTVaUnhXa3B4ZVhPQzIiLCJpYXQiOjE2NTUzNjU4MDQsImV4cCI6MTY1NTM2OTQwNCwicGhvbmVfbnVtYmVyIjoiKzQwNzQxMTk4NjA2IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJwaG9uZSI6WyIrNDA3NDExOTg2MDYiXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwaG9uZSJ9fQ.XRvMth7OnHphNwlSFcqCEFmeFHC4Z7Ogo4Va6JDob-IaR2F55DibdiiN275oSLhbpp1ZIX4GNX26brEkmglsv9A9XEIcHt4q60wLWGymr54EIz62VTTJM1JAGXtSs0cUu9jaBIDKz8NsG9oMf9psx10ilmZejyJWNCHJNmjuLTGVgaUfKPsu3N1TUUDOKdnX2w_QzYQzG8g8zRTPrZOWKcFObhglu9UXrWq7yhgceWenVTt2J8m7VHmk9b9i7XczYvhcCzJ0yJGOrZWvUG04vOuELfvoNzBtMvOLBJ-6xneMpdRAv_dCZ_R3ksCThJ3i02ltzTDrB-44BFcjCLdg4A"

func TestEventsEndpoint(t *testing.T) {
	router, _ := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/event", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserEndpoint(t *testing.T) {
	router, _ := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserEndpointNotGood(t *testing.T) {
	router, _ := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
