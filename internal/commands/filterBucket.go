package commands

import (
	"fmt"
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/passdata"
	"strings"
)

func ExecuteFilterBucket(cmd []string) error {

	if len(cmd) < 2 {
		fmt.Printf("showing entries without bucket\n")
		applyBucketFilter("")
	} else {
		applyBucketFilter(strings.Join(cmd[1:], " "))
	}
	printFilteredEntries()

	if len(application.FilteredEntries) == 1 {
		return selectEntry(application.FilteredEntries[0])
	}
	return nil
}

func applyBucketFilter(filter string) {

	application.FilteredEntries = make([]*passdata.PassEntry, 0)

	for index, _ := range application.AllEntries {
		if filter == "" {
			if application.AllEntries[index].Bucket == "" {
				application.FilteredEntries = append(application.FilteredEntries, &application.AllEntries[index])
			}
		} else {
			if application.AllEntries[index].MatchesBucketPrefix(filter) {
				application.FilteredEntries = append(application.FilteredEntries, &application.AllEntries[index])
			}
		}
	}
	application.NumberOfEntriesToSelect = len(application.FilteredEntries)
	fmt.Printf("filtered entries for '%s': %d \n", filter, len(application.FilteredEntries))

	preserveSelectedIndex()
}

func RegisterBucketFilter()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Bucket Filter",
			Description:       "Filter entries by bucket",
			ShortcutHint:      "fb",
			Executable:        ExecuteFilterBucket,
			IsApplicable:      application.OnlyOnStateLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("fb")})
}
