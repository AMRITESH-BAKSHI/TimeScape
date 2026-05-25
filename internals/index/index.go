package index

import (
	"fmt"

	"os"

	"path/filepath"
	"strings"
)

func AddToIndex(filename string, hash string) error {

	indexPath := filepath.Join(".TimeScape", "index")

	data, err := os.ReadFile(indexPath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")

	found := false
	var updatedLines []string

	for _, line := range lines {
		parts := strings.Split(line, " ")

		if len(parts) > 1 && parts[0] == filename {
			updatedLines = append(updatedLines, fmt.Sprintf("%s %s", filename, hash))
			found = true

		} else if line != "" {
			updatedLines = append(updatedLines, line)

		}
	}

	if !found {

		updatedLines = append(updatedLines, fmt.Sprintf("%s %s", filename, hash))
	}

	finalData := strings.Join(updatedLines,"\n");

	os.WriteFile(indexPath,[]byte(finalData),0644)
return nil;
}
