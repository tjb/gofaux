package writer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func WriteToFile(filename string, data interface{}) {
	// Ensure the filename has the .json extension
	if filepath.Ext(filename) != ".json" {
		filename = filename + ".json"
	}

	// Create the JSON data with indentation for readability
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON for saving: %v", err)
	}

	// Ensure the directory exists before saving
	dir := "tmp"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Build the full file path
	filePath := filepath.Join(dir, filename)

	// Write the JSON to a file
	err = os.WriteFile(filePath, jsonBytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write JSON to file: %v", err)
	}

	fmt.Printf("JSON successfully saved to %s\n", filePath)
}

func ParseFiles(directory string) (map[string]string, error) {
	fileMap := make(map[string]string)

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fileName := file.Name()
			name := fileName[:len(fileName)-len(filepath.Ext(fileName))]
			fileMap[name] = filepath.Join(directory, fileName)
		}
	}

	return fileMap, nil
}
