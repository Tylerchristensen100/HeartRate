package parse

import (
	"encoding/json"
	"os"
)

func Directory(directory string) ([]Recording, error) {
	info, err := os.Stat(directory)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, os.ErrInvalid
	}
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var recordings []Recording
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := directory + "/" + file.Name()
		recording, err := File(filePath)
		if err != nil {
			return nil, err
		}
		recordings = append(recordings, recording...)
	}

	return recordings, nil
}

func File(filePath string) ([]Recording, error) {
	var recordings []Recording

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &recordings); err != nil {
		return nil, err
	}

	return recordings, nil
}
