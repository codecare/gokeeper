package passdata

import (
	"encoding/json"
	"strings"
	"time"
)

type PassEntry struct {
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Login           string          `json:"login"`
	RecoveryMail    string          `json:"recoveryMail"`
	LastUsage       time.Time       `json:"lastUsage"`
	CryptedPassword CryptoContainer `json:"cryptedPassword"`
	Bucket	        string          `json:"bucket"`
}

type CryptoContainer struct {
	Algorithm	  string `json:"algorithm"`
	Nonce string `json:"nonce"`
	KeySalt string `json:"keySalt"`
	EncryptedText string `json:"encryptedText"`
}


func PassEntriesToJson(passEntries []PassEntry) ([]byte, error) {

	return json.Marshal(passEntries)
}

func PassEntriesFromJson(bytes []byte) ([]PassEntry, error) {

	var deser []PassEntry
	err := json.Unmarshal(bytes, &deser)
	return deser, err
}

func CryptoContainerToJson(cryptoContainer CryptoContainer) ([]byte, error) {

	return json.Marshal(cryptoContainer)
}

func CryptoContainerFromJson(bytes []byte) (CryptoContainer, error) {

	var deser CryptoContainer
	err := json.Unmarshal(bytes, &deser)
	return deser, err
}

func (passEntry PassEntry) ContainsData() bool {
	if len(passEntry.Name) > 0 { return true }
	if len(passEntry.Description) > 0 { return true }
	if len(passEntry.Login) > 0 { return true }
	if len(passEntry.RecoveryMail) > 0 { return true }
	return passEntry.CryptedPassword.ContainsData()
}

func (passEntry PassEntry) Title() string {
	if len(passEntry.Name) > 0 { return passEntry.Name }
	if len(passEntry.Description) > 0 { return limit(passEntry.Description, 40) }
	if len(passEntry.Login) > 0 { return limit(passEntry.Login, 40) }
	if len(passEntry.RecoveryMail) > 0 { return limit(passEntry.RecoveryMail, 40) }
	return passEntry.CryptedPassword.Algorithm
}

func (passEntry PassEntry) MatchesFilter(s string) bool {
	s = strings.ToLower(s)
	if strings.Contains(strings.ToLower(passEntry.Name), s) { return true }
	if strings.Contains(strings.ToLower(passEntry.Description), s) { return true }
	if strings.Contains(strings.ToLower(passEntry.Login), s) { return true }
	if strings.Contains(strings.ToLower(passEntry.RecoveryMail), s) { return true }
	return false
}

func (passEntry PassEntry) Duplicate() PassEntry {
	return PassEntry{
		Name:            passEntry.Name,
		Description:     passEntry.Description,
		Login:           passEntry.Login,
		RecoveryMail:    passEntry.RecoveryMail,
		LastUsage:       passEntry.LastUsage,
		CryptedPassword: passEntry.CryptedPassword,
		Bucket:          passEntry.Bucket,
	}
}

func (passEntry PassEntry) MatchesBucketPrefix(bucketPrefix string) bool {
	bucketPrefix = strings.ToLower(bucketPrefix)
	if strings.HasPrefix(strings.ToLower(passEntry.Bucket), bucketPrefix) { return true }

	return false
}

func (cryptoContainer CryptoContainer) ContainsData() bool {
	return len(cryptoContainer.EncryptedText) > 0
}

func limit(input string, length int) string {
	if len(input) <= length { return input }
	return input[:length]
}
