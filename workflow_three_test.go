package main

import (
	"independentjournal/domain"
	"testing"
)

func TestWorkflowThree(t *testing.T) {
	s := fixtureService(t)
	r := fixtureRecord()
	u := domain.NewUser("u1", "Indie", "developer", r.CreatedAt)
	_, _ = s.RegisterEntry(r, u)
	_, e := s.WithdrawEntry(r.ID, "u1")
	if e != nil {
		t.Fatal(e)
	}
}
