package main

import (
	"independentjournal/domain"
	"testing"
)

func TestWithdrawalPreservesCompleteStatus(t *testing.T) {
	s := fixtureService(t)
	r := fixtureRecord()
	u := domain.NewUser("u1", "Indie", "developer", r.CreatedAt)
	_, _ = s.RegisterEntry(r, u)
	got, e := s.WithdrawEntry(r.ID, "u1")
	if e != nil {
		t.Fatal(e)
	}
	if !got.IsComplete() {
		t.Fatalf("withdrawal lost complete state: %+v", got)
	}
}
