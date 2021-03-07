package commands

import (
	"errors"
	"fmt"
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"github.com/codecare/gokeeper/internal/passdata"
	"github.com/codecare/gokeeper/internal/shell"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ExecuteExportBucket(cmd []string) error {
	fmt.Println("-------- exporting bucket to new vault --------")

	fmt.Printf("please select bucket prefix: ")
	bucketPrefix := shell.ReadInput()

	var filteredEntries []*passdata.PassEntry

	filteredEntries = filterByBucketPrefix(bucketPrefix)
	if len(filteredEntries) == 0 {
		fmt.Printf("no entries found for bucket prefix %v\n", bucketPrefix)
		return nil
	}

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

	password, err := ExecuteSelectReturnPassword()
	if err != nil { return err }

	// pass entries have to be recoded with new password!
	err = executeSaveNewFileWithNewPassword(file, filteredEntries, password)

	if err != nil {
		fmt.Printf("file system error %v\n", err)
		return err
	}

	fmt.Printf("Exported file: %v", file)

	return nil
}

func executeSaveNewFileWithNewPassword(file string, entries []*passdata.PassEntry, password []byte) error {
	// do not override existing entries!
	var rekeyedEntries []passdata.PassEntry

	for index, _ := range entries {
		entryToRekey := entries[index]

		passwordFromContainer, err := crypt.DecryptFromContainer(entryToRekey.CryptedPassword)
		if err != nil { return err }

		// two passwords
		container, err := crypt.EncryptToContainerWithPassword(password, passwordFromContainer)
		if err != nil { return err }

		var rekeyedEntry = entryToRekey.Duplicate()
		rekeyedEntry.CryptedPassword = container
		rekeyedEntries = append(rekeyedEntries, rekeyedEntry)
	}

	fmt.Printf("------------------ saving -------------------\n")
	fmt.Printf("file: %v\n", file)

	bytes, err := passdata.PassEntriesToJson(rekeyedEntries)
	if err != nil { return err }

	cryptoContainer, err := crypt.EncryptToContainerWithPassword(password, bytes)
	if err != nil { return err }

	json, err := passdata.CryptoContainerToJson(cryptoContainer)
	if err != nil { return err }

	err = ioutil.WriteFile(file, json, 0600)
	if err != nil { return err }

	fmt.Printf("------------------ saved -------------------\n")
	return nil
}

func filterByBucketPrefix(bucketPrefix string) []*passdata.PassEntry {
	passEntries := make([]*passdata.PassEntry, 0)

	for index, _ := range application.AllEntries {
		if application.AllEntries[index].MatchesBucketPrefix(bucketPrefix) {
			passEntries = append(passEntries, &application.AllEntries[index])
		}
	}
	return passEntries
}


func RegisterExportBucket()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Export Bucket",
			Description:  "Exports a bucket to a new vault",
			ShortcutHint:     "xb",
			Executable:   ExecuteExportBucket,
			IsApplicable: application.OnlyOnStateLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("xb")})
}
