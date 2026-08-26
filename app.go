package main

import (
	"independentjournal/domain"
	"independentjournal/service"
	"independentjournal/store"
	"os"
	"testing"
	"time"
)

func fixtureService(t *testing.T) *service.Service {
	t.Helper()
	f := t.TempDir() + "/db"
	s, e := store.Open(f)
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return service.New(s)
}
func fixtureRecord() domain.Record {
	return domain.NewRecord("r49", "Pixel Garden", "weekly build notes", "u1", time.Unix(1700000000, 0))
}
func ensureTemp() string {
	f, e := os.CreateTemp("", "journal")
	if e != nil {
		return ""
	}
	n := f.Name()
	f.Close()
	os.Remove(n)
	return n
}
