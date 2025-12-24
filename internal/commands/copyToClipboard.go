package commands

import (
	"fmt"
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"golang.design/x/clipboard"
)

func ExecuteCopyToClipboard(cmd []string) error {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	var activeEntry = application.ActiveEntry
	fmt.Println("Copied password to clipboard: " + activeEntry.Name)
	bytes, err := crypt.DecryptFromContainer(activeEntry.CryptedPassword)
	if err != nil { return err }
	clipboard.Write(clipboard.FmtText, bytes)

	return nil
}

func RegisterCopyToClipboard() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Copy to Clipboard",
			Description:       "Copy Password Value to Clipboard",
			ShortcutHint:      "cp",
			Executable:        ExecuteCopyToClipboard,
			IsApplicable:      application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("cp")})
}
