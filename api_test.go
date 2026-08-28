package main

import (
	"independentjournal/api"
	"independentjournal/domain"
	"net/http/httptest"
	"testing"
)

func TestAPIQuery(t *testing.T) {
	s := fixtureService(t)
	r := fixtureRecord()
	u := domain.NewUser("u1", "Indie", "developer", r.CreatedAt)
	_, _ = s.RegisterEntry(r, u)
	req := httptest.NewRequest("GET", "/records/"+r.ID, nil)
	w := httptest.NewRecorder()
	api.NewHandler(s).ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
