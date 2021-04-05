package commands

import (
	"errors"
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/passdata"
	"reflect"
	"strconv"
)

func ExecuteSelect(cmd []string) error {

	if s, err := strconv.Atoi(cmd[0]); err == nil && s < application.NumberOfEntriesToSelect {
		// select the one from the complete list!
		internalErr := selectEntry(application.FilteredEntries[s])

		if internalErr != nil { return internalErr }

		preserveSelectedIndex()

		return ExecutePrintSecure(cmd)
	} else {
		return err
	}
}

func selectEntry(selectedEntry *passdata.PassEntry) error {
	// find index of active entry
	var foundAtPos = -1
	for pos, passEntry := range application.AllEntries {

		if reflect.DeepEqual(*selectedEntry, passEntry) {
			foundAtPos = pos
		}
	}

	if foundAtPos == -1 {
		return errors.New("could not select")
	} else {
		application.ActiveEntry = &application.AllEntries[foundAtPos]
	}
	return nil
}

func canHandleNumericShortCut(shortcut string) bool {
	if s, err := strconv.Atoi(shortcut); err == nil && s < application.NumberOfEntriesToSelect {
		return true
	}
	return false
}

func RegisterSelect()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Help",
			Description:  "Show Help",
			ShortcutHint:     "#",
			Executable:   ExecuteSelect,
			IsApplicable: application.OnlyOnSelectableExists,
			CanHandleShortCut: canHandleNumericShortCut})
}