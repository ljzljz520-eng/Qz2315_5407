package main

import (
	"independentjournal/journal"
	"testing"
)

func TestJournalRender(t *testing.T) {
	if journal.IsBlank(journal.StatusLabel(fixtureRecord())) {
		t.Fatal("label")
	}
}
