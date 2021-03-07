package crypt

import (
	"github.com/codecare/gokeeper/internal/application"
	"encoding/hex"
	"github.com/codecare/gokeeper/internal/passdata"
)

func EncryptToContainer(payload []byte) (passdata.CryptoContainer, error) {

	return EncryptToContainerWithPassword(application.Key, payload)
}

func EncryptToContainerWithPassword(password []byte, payload []byte) (passdata.CryptoContainer, error) {

	var result passdata.CryptoContainer

	encrypted, err := encrypt(password, payload)
	if err != nil {return result, err}

	result =  passdata.CryptoContainer{
		Algorithm:	  algorithmName(),
		Nonce: hex.EncodeToString(encrypted.Nonce),
		KeySalt: hex.EncodeToString(encrypted.KeySalt),
		EncryptedText: hex.EncodeToString(encrypted.Cipher)}

	return result, nil
}

func DecryptFromContainer(container passdata.CryptoContainer) ([]byte, error) {

	var result []byte

	if algorithmName() != container.Algorithm {
		return result, nil
	}

	cipherBytes, err := hex.DecodeString(container.EncryptedText)
	if err != nil {return result, err}

	nonce, err := hex.DecodeString(container.Nonce)
	if err != nil {return result, err}

	keySalt, err := hex.DecodeString(container.KeySalt)
	if err != nil {return result, err}

	var encrypted = Encrypted{ KeySalt: keySalt, Nonce: nonce, Cipher: cipherBytes	}
	decrypted, err := decrypt(application.Key, encrypted)
	if err != nil {return result, err}

	return decrypted, nil
}

