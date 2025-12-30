package commands

import (
	"fmt"

	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"golang.design/x/clipboard"
)

func ExecuteCopySelectedPasswordToClipboard(cmd []string) error {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	var activeEntry = application.ActiveEntry
	bytes, err := crypt.DecryptFromContainer(activeEntry.CryptedPassword)
	if err != nil {
		return err
	}
	clipboard.Write(clipboard.FmtText, bytes)
	fmt.Println("Copied password to clipboard: " + activeEntry.Name)

	return nil
}

func ExecuteCopySelectedLoginToClipboard(cmd []string) error {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	var activeEntry = application.ActiveEntry

	clipboard.Write(clipboard.FmtText, ([]byte)(activeEntry.Login))
	fmt.Println("Copied login to clipboard: " + activeEntry.Name)

	return nil
}

func ExecuteCopyLastGeneratedPasswordToClipboard(cmd []string) error {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, application.LastGeneratedPassword)
	fmt.Println("Copied last generated password to clipboard ")

	return nil
}

func RegisterCopySelectedPasswordToClipboard() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Copy Password to Clipboard",
			Description:       "Copy Selected Password Value to Clipboard",
			ShortcutHint:      "cp",
			Executable:        ExecuteCopySelectedPasswordToClipboard,
			IsApplicable:      application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("cp")})
}

func RegisterCopySelectedLoginToClipboard() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Copy Login to Clipboard",
			Description:       "Copy Selected Login Value to Clipboard",
			ShortcutHint:      "cl",
			Executable:        ExecuteCopySelectedLoginToClipboard,
			IsApplicable:      application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("cl")})
}

func RegisterCopyLastGeneratedPasswordToClipboard() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Copy last generated password",
			Description:       "Copies the last generated password to the clipboard",
			ShortcutHint:      "cg",
			Executable:        ExecuteCopyLastGeneratedPasswordToClipboard,
			IsApplicable:      application.LastGeneratedPasswordExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("cg")})
}
