package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
)

func ExecuteEditEntry(cmd []string) error {

	var newEntry = application.ActiveEntry
	newEntry.Name = readValue("name", newEntry.Name, extractCmd(cmd, 1))
	newEntry.Description = readValue("description", newEntry.Description, extractCmd(cmd,2))
	newEntry.Login =  readValue("login", newEntry.Login, extractCmd(cmd, 3))
	bytes, err := crypt.DecryptFromContainer(application.ActiveEntry.CryptedPassword)
	if err != nil { return err }

	password := readValue("password", string(bytes), extractCmd(cmd, 4))

	container, err := crypt.EncryptToContainer([]byte(password))
	if err != nil { return err }
	newEntry.CryptedPassword = container

	newEntry.Bucket =  readValue("bucket", newEntry.Bucket, extractCmd(cmd, 5))

	return ExecuteSave()
}

func extractCmd(cmd []string, index int) string {
	if len(cmd) <= index { return "" }
	return cmd[index]
}

func RegisterEditEntry()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Edit Entry",
			Description:  "Edit selected entry",
			ShortcutHint:     "e",
			Executable:   ExecuteEditEntry,
			IsApplicable: application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("e")})
}
