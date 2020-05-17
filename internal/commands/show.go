package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"fmt"
)

func ExecuteShow(cmd []string) error {
	fmt.Printf("------------------ show all -------------------\n")
	fmt.Printf("file: %v\n", application.SelectedFile)
	resetFilter()
	printFilteredEntries()

	fmt.Printf("------------------------------------------\n")
	return nil
}

func printFilteredEntries() {
	for index, passEntry := range application.FilteredEntries {
		printEntrySecure(index, *passEntry)
	}
}

func RegisterShow()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Show",
			Description:  "Show all entries",
			ShortcutHint:     "s",
			Executable:   ExecuteShow,
			IsApplicable: application.OnlyOnStateLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("s")})
}