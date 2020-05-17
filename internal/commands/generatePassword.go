package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"fmt"
)

func ExecuteGeneratePassword(cmd []string) error {
	fmt.Println("------------------ generating password -------------------")
	// todo uli len as cmd[1] optional
	bytes, err := crypt.GeneratePassword(15)
	if err != nil { return err }
	fmt.Printf("%s\n\n", string(bytes))
	application.LastGeneratedPassword = bytes
	return nil
}

func RegisterGeneratePassword()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Generate",
			Description:       "Generate New Password",
			ShortcutHint:      "g",
			Executable:        ExecuteGeneratePassword,
			IsApplicable:      application.AlwaysApplicable,
			CanHandleShortCut: application.CanHandleShortCutClosure("g")})
}

func ExecuteGeneratePasswordWrite(cmd []string) error {
	err := ExecuteGeneratePassword(cmd)
	if err != nil { return err }
	err = ExecuteWriteLasteGeneratedPassword(cmd)
	return err
}

func RegisterGeneratePasswordWrite()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Generate and write",
			Description:       "Generate New Password and write to previous active application",
			ShortcutHint:      "gw",
			Executable:        ExecuteGeneratePasswordWrite,
			IsApplicable:      application.AlwaysApplicable,
			CanHandleShortCut: application.CanHandleShortCutClosure("gw")})
}