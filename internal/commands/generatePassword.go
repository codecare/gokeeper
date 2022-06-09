package commands

import (
	"fmt"
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"strconv"
	"strings"
)

func ExecuteGeneratePassword(cmd []string) error {
	fmt.Println("------------------ generating password -------------------")

	passwordLength := parseLengthParameter(cmd)
	allowedChars := determineAllowedChars(cmd)
	bytes, err := crypt.GeneratePassword(passwordLength, allowedChars)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n\n", string(bytes))
	application.LastGeneratedPassword = bytes
	return nil
}

func determineAllowedChars(cmd []string) string {
	if len(cmd) >= 3 {
		var charSet = strings.TrimSpace(cmd[2])
		cleaned := strings.TrimSpace(charSet)
		if cleaned == "s" || cleaned == "sh" || cleaned == "shell" {
			return crypt.AllowedCharsShellFriendly
		}
	}
	return crypt.AllowedCharsAll
}

func parseLengthParameter(cmd []string) int {
	var passwordLength = 15
	if len(cmd) >= 2 {
		var lenStr = strings.TrimSpace(cmd[1])
		intVar, err := strconv.Atoi(lenStr)
		if err == nil {
			if intVar < 8 {
				passwordLength = 8
			} else if intVar > 128 {
				passwordLength = 128
			} else {
				passwordLength = intVar
			}
		}
	}
	return passwordLength
}

func RegisterGeneratePassword() {
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
	if err != nil {
		return err
	}
	err = ExecuteWriteLasteGeneratedPassword(cmd)
	return err
}

func RegisterGeneratePasswordWrite() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Generate and write",
			Description:       "Generate New Password and write to previous active application",
			ShortcutHint:      "gw",
			Executable:        ExecuteGeneratePasswordWrite,
			IsApplicable:      application.AlwaysApplicable,
			CanHandleShortCut: application.CanHandleShortCutClosure("gw")})
}
