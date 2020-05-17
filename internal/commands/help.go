package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"fmt"
)

func ExecuteHelp(cmd []string) error {
	fmt.Println("help called")
	fmt.Println("------------------ available commands -------------------")
	for i := range application.CommandDescriptions {
		var description = application.CommandDescriptions[i]
		if description.IsApplicable() {
			fmt.Printf("%-8s   %-30s    %s\n", description.ShortcutHint, description.Name, description.Description)
		}
	}
	fmt.Println()
	return nil
}

func RegisterHelp()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Help",
			Description:       "Show Help",
			ShortcutHint:      "?",
			Executable:        ExecuteHelp,
			IsApplicable:      application.AlwaysApplicable,
			CanHandleShortCut: application.CanHandleShortCutClosure("?")})
}
