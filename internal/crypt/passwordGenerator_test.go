package crypt

import (
	"fmt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {

	for i:=0; i < 100; i++ {
		bytes, err := GeneratePassword(16)
		if err != nil {
			t.Errorf("save() error = %v", err);
			return
		}

		fmt.Printf("pass: %s\n", string(bytes))
	}
}

