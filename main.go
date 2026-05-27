package main

import (
	"TimeScape/internals/hashing"
	"TimeScape/internals/objects"
	"TimeScape/internals/repository"
	"TimeScape/internals/storage"
	"TimeScape/utils"

	"fmt"
	"os"
)

func main() {

	command := os.Args[1]

	switch command {

	case "hash-code":
		{
			filename := os.Args[2]
			_, err := os.Stat(filename)
			if err != nil {
				fmt.Println("file does not exists")
				return
			}

			content, err := utils.ReadFile(filename)

			if err != nil {
				fmt.Println("Error reading file")
				return
			}

			blob := objects.CreateBlobObject(content)

			hash := hashing.GenerateSHA1(blob)
			err = storage.StoreObject(hash, blob)

			if err != nil {
				fmt.Println("Error storing file:", err)
				return
			}
			fmt.Println(hash)
		}

	case "cat-file":
		{

			hash := os.Args[2]

			data, err := storage.ReadObject(hash)

			if err != nil {
				fmt.Println("Cannot read object")
				return
			}

			original_data := objects.ExtractBlobContent(data)

			fmt.Println(string(original_data))

		}

	case "init":
		{
			err := repository.InitRepository()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

	case "add":
		{
			filename := os.Args[2]

			err := repository.AddFile(filename)

			if err != nil {
				return
			}

		}
	// case "write-tree":
	// 	{
	// 		hash, err := repository.WriteTree()
	// 		if err != nil {
	// 			fmt.Println("error", err)
	// 			return
	// 		}

	// 		fmt.Println(hash)

	// 	}

	case "commit":
		{

			hash, err := repository.Commit(os.Args[2])

			if err != nil {
				fmt.Println("error", err)
				return
			}

			fmt.Println(hash)

			return

		}
	case "log":
		{
			err := repository.Loghistory()
			if err != nil {
				return
			}
		}

	case "checkout":
		{

			commithash := os.Args[2]
			err := repository.Checkout(commithash)

			if err != nil {
				fmt.Println("error", err)
				return
			}

			return
		}

	}

}
