package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"fmt"
	"github.com/codecare/gokeeper/internal/shell"
)

func ExecuteTestKeyboard(cmd []string) error {
	fmt.Println("test keyboard called")
	return shell.VerifyImportantCharmapping()
}

func RegisterTestKeyboard()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Test Keyboard",
			Description:  "Test character mapping for exporting passwords",
			ShortcutHint:     "test",
			Executable:   ExecuteTestKeyboard,
			IsApplicable: application.AlwaysApplicable,
			CanHandleShortCut: application.CanHandleShortCutClosure("test")})
}