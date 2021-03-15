package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"testing"
)

func Test_copyPointer(t *testing.T) {
	application.AllEntries = factorisePassEntries()
	applyFilter("descr")
	if application.NumberOfEntriesToSelect != 2 {
		t.Errorf("filter failed")
	}
	entries := application.FilteredEntries
	entry := *entries[0]
	entry.Login = "huhu"

	if  entries[0].Login != "us@codecare.de" {
		t.Errorf("filter failed")
	}
	if  application.AllEntries[0].Login != "us@codecare.de" {
		t.Errorf("filter failed")
	}
	if  entry.Login != "huhu" {
		t.Errorf("filter failed")
	}
}
