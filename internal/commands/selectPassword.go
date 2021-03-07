package commands

import "github.com/codecare/gokeeper/internal/application"

func ExecuteSelectPassword() error {
	p1 := readPassword()
	application.Key = []byte(p1)
	return nil
}

func ExecuteSelectReturnPassword() ([]byte, error) {
	p1 := readPassword()
	return []byte(p1), nil
}
