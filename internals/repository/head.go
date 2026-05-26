package repository

import (
	"os"
	"strings"
)

func GetHead() string {
	data, err := os.ReadFile(".TimeScape/HEAD")

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func UpdateHead(hash string) error {

	return os.WriteFile(".TimeScape/HEAD", []byte(hash), 0644)

}
