package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"fmt"
	"github.com/codecare/gokeeper/internal/passdata"
)

func ExecuteNewEntry(cmd []string) error {

	var newEntry passdata.PassEntry
	newEntry.Name = readValue("name", newEntry.Name, extractCmd(cmd, 1))
	newEntry.Description = readValue("description", newEntry.Description, extractCmd(cmd, 2))
	newEntry.Login =  readValue("login", newEntry.Login, extractCmd(cmd, 3))
	password := readValue("password", "", extractCmd(cmd, 4))

	container, err := crypt.EncryptToContainer([]byte(password))
	if err != nil { return err }
	newEntry.CryptedPassword = container

	newEntry.Bucket =  readValue("bucket", newEntry.Bucket, extractCmd(cmd, 5))

	application.AllEntries = append(application.AllEntries, newEntry)

	fmt.Printf("crated new entry: %s\n", newEntry.Name)

	application.ActiveEntry = &newEntry
	resetFilter()

	return ExecuteSave()
}

func RegisterNewEntry()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "New Entry",
			Description:  "Create new entry",
			ShortcutHint:     "n",
			Executable:   ExecuteNewEntry,
			IsApplicable: application.OnlyOnStateLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("n")})
}
