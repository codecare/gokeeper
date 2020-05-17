package commands


import "C"
import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"github.com/codecare/gokeeper/internal/shell"
	"strings"
	"syscall"
)

func ExecuteConfirm() error {
	fmt.Println("please confirm (y/n)")
	confirmation := shell.ReadInput()
	confirmation = strings.TrimSpace(confirmation)
	if strings.EqualFold("y", confirmation) {
		return nil
	}
	return errors.New("not confiremd by user")
}

func readValue(parameterName, defaultValue, newValue string) string {
	// use from command
	if len(newValue) > 0 {
		fmt.Printf("%-20s ->[%s]\n", parameterName + ":", newValue)
		return newValue
	}
	fmt.Printf("%-20s [%s]\n", parameterName + ":", defaultValue)
	inputValue := shell.ReadInput()
	if len(inputValue) < 1 {
		return defaultValue
	}
	return inputValue
}

func readValueNoDefault(parameterName string) string {
	fmt.Printf("%-20s \n", parameterName + ":")
	return shell.ReadInput()
}





func readPassword() string {

	if (terminal.IsTerminal(int(syscall.Stdin))) {

		fmt.Print("Enter Password (terminal): ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return ""
		}
		password := string(bytePassword)

		return strings.TrimSpace(password)

	} else {

		fmt.Print("Enter Password (no terminal): ")
		inputValue := shell.ReadInput()
		return strings.TrimSpace(inputValue)
	}
}



func fileDoesExist(file string) (bool, error) {

	_, err := os.Stat(file)

	if err != nil {

		if os.IsNotExist(err) {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

