package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/codecare/gokeeper/internal/passdata"
	"github.com/codecare/gokeeper/internal/shell"
	"strings"
)

func ExecuteOpenVault(cmd []string) error {
	fmt.Println("-------- opening vault --------")

	file, err := getFilePath(cmd)
	if err != nil {return err}

	doesExist, err := fileDoesExist(file)
	if err != nil {return err}

	if !doesExist {
		return errors.New(fmt.Sprintf("file does not exist %s", file))
	} else 	{
		fmt.Printf("vault: %s\n", file)
	}

	err = ExecuteSelectPassword()
	if err != nil { return err }

	err = ExecuteLoad(file)

	if err != nil {
		fmt.Printf("file system error %v\n", err)
		return err
	}

	application.SelectedFile = file
	application.CurrentLoadingState = application.Loaded

	return nil
}

func getFilePath(cmd []string) (string, error) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("file system error %v", err)
		return "", err
	}

	var file string
	if len(cmd) >= 2 {
		file = cmd[1]
	} else {
		file, err = readFileName(currentDirectory)
		if err != nil {
			return "", err
		}
	}

	if !strings.HasPrefix(file, "/") {
		file = currentDirectory + "/" + file
	}
	return file, nil
}

func readFileName(currentDirectory string) (string, error) {
	fmt.Println("please select file: ")
	fmt.Printf("as absolute path starting with / or as relative path to %v\n", currentDirectory)

	file := shell.ReadInput()
	file = strings.TrimSpace(file)
	if len(file) < 1 {
		fmt.Printf("file name too short\n\n")
		return "", errors.New("file name too short")
	}
	return file, nil
}

func ExecuteLoad(file string) error {
	fmt.Printf("------------------ loading -------------------\n")
	fmt.Printf("file: %v\n", file)

	data, err := ioutil.ReadFile(file)
	if err != nil { return err }

	cryptoContainer, err := passdata.CryptoContainerFromJson(data)
	if err != nil { return err }

	jsonData, err := crypt.DecryptFromContainer(cryptoContainer)
	if err != nil { return err }

	passEntries, err := passdata.PassEntriesFromJson(jsonData)
	if err != nil { return err }

	application.AllEntries = passEntries
	application.NumberOfEntriesToSelect = len(application.AllEntries)

	resetFilter()

	fmt.Printf("------------------ loaded -------------------\n")
	return nil
}


func RegisterOpenVault()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Open Vault",
			Description:  "Open existing vault",
			ShortcutHint:     "o",
			Executable:   ExecuteOpenVault,
			IsApplicable: application.OnlyOnStateNotLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("o")})
}
