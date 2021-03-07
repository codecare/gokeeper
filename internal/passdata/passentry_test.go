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

	var tobe = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2019-05-14T11:45:26.371Z\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"},\"bucket\":\"b\"}]"

// Test_toJson: passentry_test.go:23: toJson() tobe:
//[{"name":"name","description":"descr","login":"us@codecare.de","recoveryMail":"","lastUsage":"2019-05-14T11:45:26.371Z","cryptedPassword":{"algorithm":"PBKDF2:4096:256:AES:GCM","nonce":"affe","keySalt":"cafe","encryptedText":"babe"}}] + is:
//[{"name":"name","description":"descr","login":"us@codecare.de","recoveryMail":"","lastUsage":"2019-05-14T11:45:26.371Z","cryptedPassword":{"algorithm":"PBKDF2:4096:256:AES:GCM","nonce":"affe","keySalt":"cafe","encryptedText":"babe"},"bucket":""}]

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
	passEntry1 := PassEntry{Name: "name", Description: "descr", Login: "us@codecare.de", LastUsage: timestamp, CryptedPassword: crypted1, Bucket: "b"}
	var passEntries []PassEntry
	passEntries = append(passEntries, passEntry1)
	return passEntries
}

func Test_fromJson(t *testing.T) {

	tobe := factorisePassEntries()

	var inputData = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2019-05-14T11:45:26.371Z\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"},\"bucket\":\"b\"}]"

	passEntries, err := PassEntriesFromJson([]byte(inputData))
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(tobe, passEntries) {
		t.Errorf("fromJson() tobe: \n\t%v + is: \n\t%v", tobe, passEntries)
	}
}

func Test_fromJsonBackwardCompatabilityBucket(t *testing.T) {

	tobe := factorisePassEntries()
	tobe[0].Bucket = ""

	var inputData = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2019-05-14T11:45:26.371Z\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"}}]"

	passEntries, err := PassEntriesFromJson([]byte(inputData))
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(tobe, passEntries) {
		t.Errorf("fromJson() \n\ttobe: %v + \n\tis  : %v", tobe, passEntries)
	}
}

func Test_fromJsonTimeZones(t *testing.T) {

	tobe := factorisePassEntries()
	tobe[0].Bucket = ""
	dateString := "2019-05-14T11:45:26.371+02:00"
	timestamp, _ := time.Parse(time.RFC3339, dateString)
	tobe[0].LastUsage = timestamp

	var inputData = "[{\"name\":\"name\",\"description\":\"descr\",\"login\":\"us@codecare.de\",\"recoveryMail\":\"\",\"lastUsage\":\"2019-05-14T11:45:26.371+02:00\",\"cryptedPassword\":{\"algorithm\":\"PBKDF2:4096:256:AES:GCM\",\"nonce\":\"affe\",\"keySalt\":\"cafe\",\"encryptedText\":\"babe\"}}]"

	passEntries, err := PassEntriesFromJson([]byte(inputData))
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(tobe, passEntries) {
		t.Errorf("fromJson() \n\ttobe: %v + \n\tis  : %v", tobe, passEntries)
	}
}



func Test_filter(t *testing.T) {

	entries := factorisePassEntries()
	matchesFilter := entries[0].MatchesFilter("name")
	if !matchesFilter {
		t.Errorf("filter not working")
	}
}


