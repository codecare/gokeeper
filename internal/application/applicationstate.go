package application

import (
	. "github.com/codecare/gokeeper/internal/passdata"
)

type LoadingState string

const (
	Loaded    LoadingState = "Loaded"
	NotLoaded LoadingState = "NotLoaded"
)

type ActiveIndex struct {
	Global int
	Filter int
}

func (activeIndex ActiveIndex) GetForDisplay() int {
	return activeIndex.Filter
}

func (activeIndex ActiveIndex) Reset() {
	activeIndex.Global = -1
	activeIndex.Filter = -1
}

var CurrentLoadingState = NotLoaded

var SelectedFile string

var AllEntries []PassEntry

// pointer to the active entry in all entries
var ActiveEntry *PassEntry

var FilteredEntries []*PassEntry

var CurrentActiveIndex = ActiveIndex{-1, -1}

var Key []byte

var LastGeneratedPassword []byte

var NumberOfEntriesToSelect = 0