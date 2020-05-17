package crypt

import (
	"crypto/rand"
	"io"
)

var allowedChars = "!#$%&*+,-.0123456789:;<=>?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GeneratePassword(length int) ([]byte, error) {

	nonce := make([]byte, 1)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {return nil, err}

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i], err = validRandomByte()
		if err != nil {return nil, err}
	}
	return result, nil

}

func validRandomByte() (byte, error) {

	for {
		nonce := make([]byte, 1)
		_, err := io.ReadFull(rand.Reader, nonce)
		if err != nil {return 0, err}

		index := int(nonce[0])
		if index < len(allowedChars) {
			return allowedChars[index], nil
		}
	}
}