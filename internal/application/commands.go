package application

import (
	"errors"
)

type ExecuteCommand func(cmd []string) error
type CanHandleShortCut func(shortcut string) bool
type CheckCommandApplicable func() bool

type CommandDescription struct {
	Name              string
	Description       string
	ShortcutHint      string
	Executable        ExecuteCommand
	IsApplicable      CheckCommandApplicable
	CanHandleShortCut CanHandleShortCut
}

var CommandDescriptions []CommandDescription

func RegisterCommand(description CommandDescription) {
	CommandDescriptions = append(CommandDescriptions, description)
}

func FindCommandByShortCut(cmd []string)(CommandDescription, error){

	if len(cmd) < 1 {
		var desc CommandDescription
		return desc, errors.New("no command found")
	}
	shortcut := cmd[0]
	var result CommandDescription

	for _, commandDescription := range CommandDescriptions {
		var check = commandDescription.CanHandleShortCut
		if check != nil && check(shortcut) {
			return commandDescription, nil
		}
	}

	return result, errors.New("unknown shortcut " + shortcut)
}


