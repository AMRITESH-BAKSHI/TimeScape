package repository

import (
	"TimeScape/internals/objects"
	"TimeScape/internals/storage"
	"fmt"
)

func Loghistory() error {

	currentHash := GetHead()

	for currentHash != "" {
		data, err := storage.ReadObject(currentHash)

		if err != nil {
			return err
		}

		commit := objects.ParseCommitData(data)

		fmt.Println("commit: ", currentHash)
		fmt.Println("message: ", commit.MSG)
		fmt.Println("Tree: ", commit.TreeHash)
		fmt.Println()

		currentHash = commit.ParentHash
	}
	return nil
}
