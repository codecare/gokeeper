package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"fmt"
	"io/ioutil"
	"github.com/codecare/gokeeper/internal/passdata"
)

func ExecuteSave() error {
	fmt.Printf("------------------ saving -------------------\n")
	fmt.Printf("file: %v\n", application.SelectedFile)

	bytes, err := passdata.PassEntriesToJson(application.AllEntries)
	if err != nil { return err }

	cryptoContainer, err := crypt.EncryptToContainer(bytes)
	if err != nil { return err }

	json, err := passdata.CryptoContainerToJson(cryptoContainer)
	if err != nil { return err }

	err = ioutil.WriteFile(application.SelectedFile, json, 0600)
	if err != nil { return err }

	fmt.Printf("------------------ saved -------------------\n")
	return nil
}
