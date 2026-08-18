package crypt

import (
	"crypto/rand"
	"io"
)

const AllowedCharsAll = "_!#$%&*+,-.0123456789:;<=>?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const AllowedCharsShellFriendly = "_%+-.0123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const GroupingChars = "_+-."
const NonGroupingChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GeneratePassword(length int, allowedChars string) ([]byte, error) {

	nonce := make([]byte, 1)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i], err = validRandomByte(allowedChars)
		if err != nil {
			return nil, err
		}
	}
	return result, nil

}

func GenerateGroupedPassword(length int, groupSpacing int, allowedChars string, groupingChars string) ([]byte, error) {

	nonce := make([]byte, 1)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	result := make([]byte, 0, length)
	for range length {
		if (len(result)+1)%(groupSpacing+1) == 0 && len(result) > 0 {
			b, err := validRandomByte(groupingChars)
			if err != nil {
				return nil, err
			}
			result = append(result, b)
		}
		b, err := validRandomByte(allowedChars)
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

func validRandomByte(allowedChars string) (byte, error) {

	for {
		nonce := make([]byte, 1)
		_, err := io.ReadFull(rand.Reader, nonce)
		if err != nil {
			return 0, err
		}

		index := int(nonce[0])
		if index < len(allowedChars) {
			return allowedChars[index], nil
		}
	}
}
