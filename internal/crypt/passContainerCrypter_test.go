package crypt

import (
	"github.com/codecare/gokeeper/internal/application"
	"testing"
)

func TestEncryptToContainer(t *testing.T) {

	application.Key = []byte("dasfseti4tnln")

	tobe := "dasfas5545feDSÜÜÄ"
	payload := []byte(tobe)
	cryptoContainer, err := EncryptToContainer(payload)
	if err != nil {
		t.Errorf("EncryptToContainer() error = %v", err)
		return
	}

	bytes, err := DecryptFromContainer(cryptoContainer)
	if err != nil {
		t.Errorf("DecryptFromContainer() error = %v", err)
		return
	}

	is := string(bytes)
	if !(tobe == is) {
		t.Errorf("roundtrip() tobe: %v + is: %v", tobe, is)
	}

}
