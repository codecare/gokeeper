package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

type Encrypted struct {
	KeySalt []byte
	Nonce []byte
	Cipher[]byte
}

type Key struct {
	KeySalt []byte
	Key []byte
}

func encrypt(password, textToEncrypt []byte) (Encrypted, error) {

	var result Encrypted

	key, err := secureBytesFromPasswordStringNewSalt(password)
	if err != nil { return result, err }

	c, err := aes.NewCipher(key.Key)
	if err != nil { return result, err }

	gcm, err := cipher.NewGCM(c)
	if err != nil { return result, err }
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {return result, err}

	cipherBytes := gcm.Seal(nil, nonce, textToEncrypt, nil)

	result = Encrypted{
		key.KeySalt,
		nonce,
		cipherBytes }
	return result, nil
}

func decrypt(password []byte, encrypted Encrypted ) ([]byte, error) {

	key := secureBytesFromPasswordStringGivenSalt(password, encrypted.KeySalt)

	c, err := aes.NewCipher(key.Key)
	if err != nil {return nil, err}

	gcm, err := cipher.NewGCM(c)
	if err != nil {return nil, err}

	plaintext, err := gcm.Open(nil, encrypted.Nonce, encrypted.Cipher, nil)
	if err != nil {return nil, err}
	return plaintext, nil
}

func secureBytesFromPasswordStringNewSalt(password []byte) (Key, error) {
	var key Key
	key.KeySalt = make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, key.KeySalt)
	if err != nil {return key, err}

	key.Key = pbkdf2.Key(password, key.KeySalt, 16383, 32, sha512.New)
	return key, nil
}

func secureBytesFromPasswordStringGivenSalt(password, keySalt []byte) Key {
	var key Key
	key.KeySalt = keySalt
	key.Key = pbkdf2.Key(password, key.KeySalt, 16383, 32, sha512.New)
	return key
}

func algorithmName() string {
	return "PBKDF2:16383:256:SHA512:AES:GCM"
}

