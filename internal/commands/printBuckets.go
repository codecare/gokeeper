package commands

import (
	"fmt"
	"github.com/codecare/gokeeper/internal/application"
)

func ExecutePrintBuckets(cmd []string) error {

	var bucketNames []string
	for index, _ := range application.AllEntries {
		bucketNames = append(bucketNames, application.AllEntries[index].Bucket)
	}

	bucketNames = unique(bucketNames)
	fmt.Printf("differnt buckets: %d\n", len(bucketNames))
	for _, bucketName := range bucketNames {
		fmt.Printf("'%v'\n", bucketName)
	}
	return nil
}

func RegisterPrintBuckets()() {
	application.RegisterCommand(
		application.CommandDescription{
			Name:              "Print Buckets",
			Description:       "Print all bucket names",
			ShortcutHint:      "pb",
			Executable:        ExecutePrintBuckets,
			IsApplicable:      application.OnlyOnStateLoaded,
			CanHandleShortCut: application.CanHandleShortCutClosure("pb")})
}

func unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
