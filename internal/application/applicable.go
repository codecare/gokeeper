package application

func AlwaysApplicable() bool {
	return true
}

func OnlyOnStateLoaded() bool {
	return CurrentLoadingState == Loaded
}

func OnlyOnStateNotLoaded() bool {
	return CurrentLoadingState == NotLoaded
}

func LastGeneratedPasswordExists() bool {
	return LastGeneratedPassword != nil
}
func OnlyOnSelectableExists() bool {
	return NumberOfEntriesToSelect > 0
}

func CanHandleShortCutClosure(shortCut string) func(cmd string) bool {
	return func(cmd string) bool {
		return shortCut == cmd
	}
}

func OnlyOnActiveEntryExists() bool {
	return ActiveEntry != nil
}

