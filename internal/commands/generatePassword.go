package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
)

func ExecuteGeneratePassword(cmd []string) error {
	fmt.Println("------------------ generating password -------------------")

	passwordLength := parseLengthParameter(cmd)
	mode := determineMode(cmd)

	var bytes []byte
	var err error

	switch mode {
	case Normal:
		bytes, err = crypt.GeneratePassword(passwordLength, crypt.AllowedCharsAll)
	case Shell:
		bytes, err = crypt.GeneratePassword(passwordLength, crypt.AllowedCharsShellFriendly)
	case Grouped:
		bytes, err = crypt.GenerateGroupedPassword(passwordLength, 4, crypt.NonGroupingChars, crypt.GroupingChars)
	}

	if err != nil {
		return err
	}
	fmt.Printf("%s\n\n", string(bytes))
	application.LastGeneratedPassword = bytes
	return nil
}

type Mode string

const (
	Normal  Mode = "Normal"
	Shell   Mode = "Shell"
	Grouped Mode = "Grouped"
)

func determineMode(cmd []string) Mode {
	if len(cmd) >= 3 {
		cleaned := strings.TrimSpace(cmd[2])
		switch cleaned {
		case "s", "sh", "shell":
			return Shell
		case "gr", "g", "group":
			return Grouped
		}
	}
	return Normal
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
			Description:       "Generate New Password - use 'g 20' for 20 chars - use 'g 20 sh' for shell friendly 20 chars - use g 16 gr for grouped chars",
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
