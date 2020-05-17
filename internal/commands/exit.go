package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"fmt"
	"os"
)

func ExecuteExit(cmd []string) error {
	fmt.Println("exit called")
	os.Exit(0)
	return nil
}


func RegisterExit()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Exit",
			Description:       "Exit the application",
			ShortcutHint:      "x",
			Executable:        ExecuteExit,
			IsApplicable:      application.AlwaysApplicable,
			CanHandleShortCut: application.CanHandleShortCutClosure("x")})
}
