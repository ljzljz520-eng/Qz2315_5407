package main

import (
	"independentjournal/domain"
	"independentjournal/store"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/db"
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := fixtureRecord()
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.LoadRecord(r.ID)
	if e != nil || got.Title != r.Title {
		t.Fatalf("reopen: %v", e)
	}
	_ = domain.Event{}
}
