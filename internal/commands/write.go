package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"github.com/codecare/gokeeper/internal/shell"
)

func ExecuteWriteLasteGeneratedPassword(cmd  []string) error {
	err := shell.SendCharacters(string(application.LastGeneratedPassword), true)

	return err
}

func RegisterWriteLastGeneratedPassword()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Writes last generated password",
			Description:  "Writes the last generated password to the previous active application",
			ShortcutHint:     "wg",
			Executable:   ExecuteWriteLasteGeneratedPassword,
			IsApplicable: application.LastGeneratedPasswordExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("wg")})
}

func ExecuteWriteSelectedPassword(cmd  []string) error {
	bytes, err := crypt.DecryptFromContainer(application.ActiveEntry.CryptedPassword)
	if err != nil { return err }

	err = shell.SendCharacters(string(bytes), true)

	return err
}

func RegisterWriteSelectedPassword()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Writes selected password",
			Description:  "Writes the password of the selected entry to the previous active application",
			ShortcutHint:     "wp",
			Executable:   ExecuteWriteSelectedPassword,
			IsApplicable: application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("wp")})
}

func ExecuteWriteSelectedLogin(cmd  []string) error {
	err := shell.SendCharacters(application.ActiveEntry.Login, true)

	return err
}

func RegisterWriteSelectedLogin()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Writes selected login",
			Description:  "Writes the login of the selected entry to the previous active application",
			ShortcutHint:     "wl",
			Executable:   ExecuteWriteSelectedLogin,
			IsApplicable: application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("wl")})
}

