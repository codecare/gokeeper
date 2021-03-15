package commands

import (
	"github.com/codecare/gokeeper/internal/application"
)

func SetBucket(cmd []string) error {

	var newEntry = application.ActiveEntry
	newEntry.Name = application.ActiveEntry.Name
	newEntry.Description = application.ActiveEntry.Description
	newEntry.Login =  application.ActiveEntry.Login
	newEntry.CryptedPassword = application.ActiveEntry.CryptedPassword
	newEntry.Bucket =  readValue("bucket", newEntry.Bucket, extractCmd(cmd, 1))

	return ExecuteSave()
}

func RegisterSetBucket()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Set Bucket Name",
			Description:  "Set field 'bucket'",
			ShortcutHint:     "b",
			Executable:   SetBucket,
			IsApplicable: application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("b")})
}