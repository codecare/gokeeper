package application

import (
	. "github.com/codecare/gokeeper/internal/passdata"
)

type LoadingState string

const (
	Loaded    LoadingState = "Loaded"
	NotLoaded LoadingState = "NotLoaded"
)

type DisplayState struct {
	DisplayIndex int
	hidden bool
}

var CurrentLoadingState = NotLoaded

var SelectedFile string

var AllEntries []PassEntry

// pointer to the active entry in all entries
var ActiveEntry *PassEntry

var FilteredEntries []*PassEntry

var ActiveIndex int

var Key []byte

var LastGeneratedPassword []byte

var NumberOfEntriesToSelect = 0