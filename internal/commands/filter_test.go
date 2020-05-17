package commands

import (
	"github.com/codecare/gokeeper/internal/application"
	"fmt"
	. "github.com/codecare/gokeeper/internal/passdata"
	"testing"
	"time"
)

func Test_applyFilter(t *testing.T) {
	application.AllEntries = factorisePassEntries()
	applyFilter("name")
	if application.NumberOfEntriesToSelect != 1 {
		t.Errorf("filter failed")
	}
}

func Test_preserveIndex(t *testing.T) {
	application.AllEntries = factorisePassEntries()
	resetFilter()

	application.ActiveIndex = 1
	application.ActiveEntry = &application.AllEntries[1]

	preserveSelectedIndex()
	if application.ActiveIndex < 0 {
		t.Errorf("index preservation failed")
	}
}

func factorisePassEntries() []PassEntry {

	dateString := "2019-05-14T11:45:26.371Z"
	timestamp, _ := time.Parse(time.RFC3339, dateString)
	crypted1 := CryptoContainer{ Algorithm: "PBKDF2:4096:256:AES:GCM", AlgorithmMeta: "cafe", EncryptedText: "babe"}
	passEntry1 := PassEntry{Name: "name", Description: "descr", Login: "us@codecare.de", LastUsage: timestamp, CryptedPassword: crypted1}
	passEntry2 := PassEntry{Name: "other", Description: "descr", Login: "us@codecare.de", LastUsage: timestamp, CryptedPassword: crypted1}

	var passEntries []PassEntry
	passEntries = append(passEntries, passEntry1)
	passEntries = append(passEntries, passEntry2)

	return passEntries
}

func Test_slicingSlice(t *testing.T) {
	resetFilter()

	application.SelectedFile = "b.va"

	printEntries("initial")

	var cmd = []string{"n", "name 1", "desc 1", "login 1", "pass 1"}
	err := ExecuteNewEntry(cmd)
	if err != nil { panic(err)}
	printEntries("new entry 1")

	cmd = []string{"n", "name 2", "desc 2 bbbb", "login 2", "pass 2"}
	err = ExecuteNewEntry(cmd)
	if err != nil { panic(err)}
	printEntries("new entry 2")

	cmd = []string{"n", "name 3", "desc 3 bbbb", "login 3", "pass 3"}
	err = ExecuteNewEntry(cmd)
	if err != nil { panic(err)}
	printEntries("new entry 3")

	cmd = []string{"f", "bbbb"}
	err = ExecuteFilter(cmd)
	if err != nil { panic(err)}
	printEntries("filter")

	cmd = []string{"1"}
	err = ExecuteSelect(cmd)
	if err != nil { panic(err)}
	printEntries("select")

	cmd = []string{"e", "name 3b", "desc 3 bbbb-b", "login 3b", "pass 3b"}
	err = ExecuteEditEntry(cmd)
	if err != nil { panic(err)}
	printEntries("edit entry")

	cmd = []string{"s"}
	err = ExecuteShow(cmd)
	if err != nil { panic(err)}
	printEntries("show")
}

func printEntries(title string) {
	fmt.Printf("----------- %s ------\n", title)
	for i:= 0; i < len(application.AllEntries); i++ {
		fmt.Printf("name all %d: %v\n", i, application.AllEntries[i].Name)
	}
	if application.FilteredEntries != nil {
		for i:= 0; i < len(application.FilteredEntries); i++ {
			fmt.Printf("name filtered %d: %v\n", i, application.FilteredEntries[i].Name)
		}
	}
	if application.ActiveEntry != nil {
		fmt.Printf("name active: %v\n", application.ActiveEntry.Name)
	}
	fmt.Printf("---------------------\n")
}