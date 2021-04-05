package commands

import (
	"errors"
	"fmt"
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/passdata"
	"github.com/codecare/gokeeper/internal/shell"
)

func DeleteEntry(cmd []string) error {
	fmt.Println("-------- deleting the active entry --------")

	fmt.Printf("retype name to confirm: '%s'", application.ActiveEntry.Name)
	confirm := shell.ReadInput()
	if confirm != application.ActiveEntry.Name {
		fmt.Printf("name does not match\n\n")
		return errors.New("name does not match")
	}

	if application.AllEntries[application.CurrentActiveIndex.Global].Name != application.ActiveEntry.Name {
		return errors.New(fmt.Sprintf("index mismatch! %d %s!=%s", application.CurrentActiveIndex.Global, application.AllEntries[application.CurrentActiveIndex.Global].Name, application.ActiveEntry.Name ))
	}

	application.AllEntries = remove(application.AllEntries, application.CurrentActiveIndex.Global)
	application.ActiveEntry = nil
	resetFilter()
	return nil
}

func remove(slice []passdata.PassEntry, s int) []passdata.PassEntry {
    return append(slice[:s], slice[s+1:]...)
}


func RegisterDelete()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Delete Active Entry",
			Description:  "Deletes the active entry",
			ShortcutHint:     "del",
			Executable:   DeleteEntry,
			IsApplicable: application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("del")})
}