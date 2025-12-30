package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/commands"
	. "github.com/codecare/gokeeper/internal/shell"
)

func main() {

	fmt.Printf("Built on %s from git %s\n", buildTime, gitRev)

	registerCommands()

	gameLoop()

}

func gameLoop() {

	err := commands.ExecuteHelp(nil)
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		// we got more than just the executable
		fmt.Printf("command line args: %v\n", os.Args[1:])
		executeCommand(os.Args[1:])
	}

	for {
		print("\n")
		printSelected()
		fmt.Print("please enter command or ?\n")
		cmd := ReadInput()

		fields := strings.Fields(cmd)

		executeCommand(fields)
	}
}

func printSelected() {
	if application.OnlyOnActiveEntryExists() {
		fmt.Printf("selected: %s\n", application.ActiveEntry.Title())
	}
}

func executeCommand(fields []string) {
	if len(fields) < 1 {
		return
	}
	command, err := application.FindCommandByShortCut(fields)
	if err != nil {
		fmt.Printf("unknown command %s\n", fields[0])
	} else {
		if command.IsApplicable() {
			err = command.Executable(fields)
			if err != nil {
				fmt.Printf("command failed %v\n", err)
			}

		} else {
			fmt.Printf("command %v is not applicable\n\n", command.Name)
		}
	}
}

func registerCommands() {

	commands.RegisterHelp()
	commands.RegisterExit()

	commands.RegisterOpenVault()
	commands.RegisterNewVault()
	commands.RegisterNewEntry()

	commands.RegisterShow()

	commands.RegisterFilter()

	commands.RegisterGeneratePassword()
	commands.RegisterGeneratePasswordWrite()
	commands.RegisterWriteLastGeneratedPassword()

	commands.RegisterSelect()
	commands.RegisterPrintSecure()
	commands.RegisterPrintInsecure()
	commands.RegisterEditEntry()

	commands.RegisterWriteSelectedLogin()
	commands.RegisterWriteSelectedPassword()

	commands.RegisterTestKeyboard()

	commands.RegisterExportBucket()

	commands.RegisterSetBucket()
	commands.RegisterDelete()

	commands.RegisterBucketFilter()
	commands.RegisterPrintBuckets()

	commands.RegisterClearScreen()
	commands.RegisterCopySelectedPasswordToClipboard()
	commands.RegisterCopySelectedLoginToClipboard()
	commands.RegisterCopyLastGeneratedPasswordToClipboard()
}

var (
	gitRev    = "undef"
	buildTime = "undef"
)
