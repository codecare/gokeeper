package commands

import "github.com/codecare/gokeeper/internal/application"

func ExecuteSelectPassword() error {
	p1 := readPassword()
	application.Key = []byte(p1)
	return nil
}
