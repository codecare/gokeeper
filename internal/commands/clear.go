package commands

import (
	"fmt"
	"github.com/codecare/gokeeper/internal/application"
)

func ExecuteClearScreen(cmd []string) error {
	fmt.Println("clear called")
	fmt.Printf(string([]byte{0x1b, '[', '3', 'J'}))
	fmt.Println("cleared")
	return nil
}

func RegisterClearScreen() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Clear",
			Description:       "Clear screen",
			ShortcutHint:      "c",
			Executable:        ExecuteClearScreen,
			IsApplicable:      application.AlwaysApplicable,
			CanHandleShortCut: application.CanHandleShortCutClosure("c")})
}
