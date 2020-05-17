package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"github.com/codecare/gokeeper/internal/crypt"
	"fmt"
	"github.com/codecare/gokeeper/internal/passdata"
)

func printEntrySecure(index int, passEntry passdata.PassEntry) {
	fmt.Printf("%d  ------------------------------------------\n", index)
	printValue("name:", passEntry.Name)
	printValue("description:", passEntry.Description)
	printValue("login:", passEntry.Login)
	printValue("passwordAlg:", passEntry.CryptedPassword.Algorithm)
}

func ExecutePrintSecure(cmd []string) error {
	printEntrySecure(application.ActiveIndex, *application.ActiveEntry)
	fmt.Printf("\n")
	return nil
}

func ExecutePrintInsecure(cmd []string) error {
	printEntrySecure(application.ActiveIndex, *application.ActiveEntry)
	bytes, err := crypt.DecryptFromContainer(application.ActiveEntry.CryptedPassword)
	if err != nil { return err }
	printValue("password:", string(bytes))
	fmt.Printf("\n")
	return nil
}

func RegisterPrintSecure()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:         "Print Secure",
			Description:  "Prints the active entry - not revealing the password",
			ShortcutHint:     "p",
			Executable:   ExecutePrintSecure,
			IsApplicable: application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("p")})
}

func printValue(parameterName, value string) {
	fmt.Printf("%-20s [%s]\n", parameterName, value)
}
func RegisterPrintInsecure()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Print Insecure",
			Description:       "Prints the active entry - revealing the password",
			ShortcutHint:      "px",
			Executable:        ExecutePrintInsecure,
			IsApplicable:      application.OnlyOnActiveEntryExists,
			CanHandleShortCut: application.CanHandleShortCutClosure("px")})
}