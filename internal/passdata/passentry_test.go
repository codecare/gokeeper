package passdata

import (
	"reflect"
	"testing"
	"time"
)

func Test_toJson(t *testing.T) {

	passEntries := factorisePassEntries()

	json, err := PassEntriesToJson(passEntries)
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}

	var tobe = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2019-05-14T11:45:26.371Z\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"}}]"

	is := string(json)
	if !(tobe == is) {
		t.Errorf("toJson() tobe: %v + is: %v", tobe, is)
	}
}

// [{"name":"name","description":"descr","login":"us@codecare.de","recoveryMail":"","lastUsage":"2019-05-14T11:45:26.371Z","cryptedPassword":{"salt":"cafe","encryptedText":"babe","initVector":"0000"}}] + is:
// [{"name":"name","description":"descr","login":"us@codecare.de","recoveryMail":"","lastUsage":"2019-05-14T11:45:26.371Z","cryptedPassword":{"algorithm":"PBKDF2:4096:256:AES:GCM","salt":"cafe","encryptedText":"babe","initVector":"0000"}}]
//

func factorisePassEntries() []PassEntry {

	dateString := "2019-05-14T11:45:26.371Z"
	timestamp, _ := time.Parse(time.RFC3339, dateString)
	crypted1 := CryptoContainer{ Algorithm: "PBKDF2:4096:256:AES:GCM", KeySalt: "cafe", Nonce: "affe", EncryptedText: "babe"}
	passEntry1 := PassEntry{Name: "name", Description: "descr", Login: "us@codecare.de", LastUsage: timestamp, CryptedPassword: crypted1}
	var passEntries []PassEntry
	passEntries = append(passEntries, passEntry1)
	return passEntries
}

func Test_fromJson(t *testing.T) {

	tobe := factorisePassEntries()

	var inputData = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2019-05-14T11:45:26.371Z\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"}}]"

	passEntries, err := PassEntriesFromJson([]byte(inputData))
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(tobe, passEntries) {
		t.Errorf("fromJson() tobe: %v + is: %v", tobe, passEntries)
	}
}

func Test_fromJson2(t *testing.T) {

	tobe := factorisePassEntries()

	var inputData = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2019-04-06T15:35:29.612+02:00\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"}}]"

	passEntries, err := PassEntriesFromJson([]byte(inputData))
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(tobe, passEntries) {
		t.Errorf("fromJson() tobe: %v + is: %v", tobe, passEntries)
	}
}

func Test_fromJson3(t *testing.T) {

	tobe := factorisePassEntries()

	var inputData = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2016-12-08T17:06:22.437Z\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"}}]"

	passEntries, err := PassEntriesFromJson([]byte(inputData))
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(tobe, passEntries) {
		t.Errorf("fromJson() tobe: %v + is: %v", tobe, passEntries)
	}
}

func Test_filter(t *testing.T) {

	entries := factorisePassEntries()
	matchesFilter := entries[0].MatchesFilter("name")
	if !matchesFilter {
		t.Errorf("filter not working")
	}
}


