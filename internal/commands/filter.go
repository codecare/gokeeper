package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"fmt"
	"github.com/codecare/gokeeper/internal/passdata"
	"reflect"
	"strings"
)

func ExecuteFilter(cmd []string) error {

	if len(cmd) < 2 {
		fmt.Printf("resetting filter\n")
		resetFilter()
	} else {
		applyFilter(strings.Join(cmd[1:], " "))
	}
	printFilteredEntries()

	if len(application.FilteredEntries) == 1 {
		return selectEntry(application.FilteredEntries[0])
	}
	return nil
}

func applyFilter(filter string) {
	application.ActiveIndex = -1
	application.FilteredEntries = make([]*passdata.PassEntry, 0)

	for index, _ := range application.AllEntries {
		if application.AllEntries[index].MatchesFilter(filter) {
			application.FilteredEntries = append(application.FilteredEntries, &application.AllEntries[index])
		}
	}
	application.NumberOfEntriesToSelect = len(application.FilteredEntries)
	fmt.Printf("filtered entries for '%s': %d \n", filter, len(application.FilteredEntries))
	preserveSelectedIndex()
}

func resetFilter() {
	application.FilteredEntries = make([]*passdata.PassEntry, 0)
	for index, _ := range application.AllEntries {
		application.FilteredEntries = append(application.FilteredEntries, &application.AllEntries[index])
	}
	preserveSelectedIndex()
	application.NumberOfEntriesToSelect = len(application.FilteredEntries)
}

func preserveSelectedIndex() {
	// find index of active entry
	application.ActiveIndex = -1
	for pos, passEntry := range application.FilteredEntries {

		if application.ActiveEntry != nil && reflect.DeepEqual(application.ActiveEntry, passEntry) {
			application.ActiveIndex = pos
		}
	}
	if application.ActiveIndex == -1 {
		application.ActiveEntry = nil
	}
}

func RegisterFilter()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Filter",
			Description:       "Filter entries",
			ShortcutHint:      "f",
			Executable:        ExecuteFilter,
			IsApplicable:      application.OnlyOnStateLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("f")})
}
