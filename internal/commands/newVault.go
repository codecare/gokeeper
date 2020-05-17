package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"github.com/codecare/gokeeper/internal/shell"
	"strings"
)

func ExecuteNewVault(cmd []string) error {
	fmt.Println("-------- creating new vault --------")
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("file system error %v", err)
		return err
	}
	fmt.Println("please select file: ")
	fmt.Printf("as absolute path starting with / or as relative path to %v\n", currentDirectory)

	file := shell.ReadInput()
	file = strings.TrimSpace(file)
	if len(file) < 1 {
		fmt.Printf("file name too short\n\n")
		return errors.New("file name too short")
	}

	if !strings.HasPrefix(file, "/") {
		file = currentDirectory + "/" + file
	}

	doesExist, err := fileDoesExist(file)
	if err != nil {
		fmt.Printf("file system error checking exists %v\n", err)
		return err
	}
	if doesExist {
		fmt.Printf("file does already exist\n\n")
		return errors.New("file does already exist")
	}

	fmt.Printf("file does not exist: %s\n", file)
	fmt.Printf("create file and parent folders?\n")
	err = ExecuteConfirm()
	if err != nil {
		return errors.New("aborted by user")
	}
	dirName := filepath.Dir(file)
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Printf("file system error %v\n", err)
		return err
	}

	err = ExecuteSelectPassword()
	if err != nil { return err }

	application.SelectedFile = file
	err = ExecuteSave()

	if err != nil {
		fmt.Printf("file system error %v\n", err)
		return err
	}

	fmt.Printf("Created file: %v", file)

	application.CurrentLoadingState = application.Loaded

	return nil
}


func RegisterNewVault()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "New Vault",
			Description:  "Create new vault",
			ShortcutHint:     "v",
			Executable:   ExecuteNewVault,
			IsApplicable: application.OnlyOnStateNotLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("v")})
}