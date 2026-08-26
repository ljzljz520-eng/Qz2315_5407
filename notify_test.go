package main

import (
	"independentjournal/notify"
	"testing"
)

func TestNotifyDelivery(t *testing.T) {
	d := notify.BuildDelivery(fixtureRecord(), "developer")
	if !notify.IsDeliverable(d) {
		t.Fatal("delivery")
	}
}
