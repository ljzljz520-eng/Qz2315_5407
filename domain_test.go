package main

import (
	"independentjournal/domain"
	"testing"
)

func TestDomainValidation(t *testing.T) {
	r := domain.Record{}
	if r.Validate() == nil {
		t.Fatal("expected invalid")
	}
	if !domain.ValidRole("developer") {
		t.Fatal("role")
	}
}
