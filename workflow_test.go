package main

import (
	"independentjournal/domain"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s := fixtureService(t)
	r := fixtureRecord()
	u := domain.NewUser("u1", "Indie", "developer", r.CreatedAt)
	got, e := s.RegisterEntry(r, u)
	if e != nil || got.Status != "submitted" {
		t.Fatalf("register: %v %v", got, e)
	}
	v, e := s.QueryEntry(r.ID)
	if e != nil || v == "" {
		t.Fatal(e)
	}
}
func TestRecordFlow49(t *testing.T) {
	s := fixtureService(t)
	r := fixtureRecord()
	u := domain.NewUser("u1", "Indie", "developer", r.CreatedAt)
	_, _ = s.RegisterEntry(r, u)
	v, e := s.QueryEntry(r.ID)
	if e != nil || len(v) < 20 {
		t.Fatalf("flow49: %v", e)
	}
}
