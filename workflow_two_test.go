package main

import (
	"independentjournal/domain"
	"testing"
)

func TestWorkflowTwo(t *testing.T) {
	s := fixtureService(t)
	r := fixtureRecord()
	u := domain.NewUser("u1", "Indie", "developer", r.CreatedAt)
	_, _ = s.RegisterEntry(r, u)
	got, e := s.ReviewEntry(r.ID, "mod", "approved")
	if e != nil || got.Status != "approved" {
		t.Fatalf("review: %v %v", got, e)
	}
}
