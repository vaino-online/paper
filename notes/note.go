package notes

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SaveToDisk saves a single note into
func SaveToDisk(n *Note, folder string) error {
	filename := filepath.Join(folder, n.Title) // TODO: Sanitize title
	return os.WriteFile(filename, n.Body, 0600)
}

func LoadFromDisk(keyword string, folder string) (*Note, error) {
	filename, err := searchForKeywordInFilename(folder, keyword)
	if err != nil {
		return nil, err
	}
	body, err := os.ReadFile(filepath.Join(folder, filename))
	if err != nil {
		return nil, err
	}
	return &Note{Title: filename, Body: body}, nil
}

func searchForKeywordInFilename(folder string, keyword string) (string, error) {
	filesInFolder, _ := ioutil.ReadDir(folder)
	for _, file := range filesInFolder {
		// FIXME: This is inefficient because it reads the whole file at once
		fileBytes, err := ioutil.ReadFile(filepath.Join(folder, file.Name()))
		if err != nil {
			// This is not normal but we can safeuly ignore it.
			log.Printf("Could not read file at %v", file.Name())
		}
		fileContents := string(fileBytes)
		if strings.Contains(fileContents, keyword) {
			return file.Name(), nil
		}
	}
	return "", fmt.Errorf("No notes found for keyword: %v", keyword)
}
