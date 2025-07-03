package export

import (
	"encoding/json"
	"os"

	"org.freethegnomes.heartRate/internal/parse"
)

func JSON(recordings []parse.RecordingFlat, filePath string) error {
	err := createDirectoryIfNotExists(filePath)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(recordings, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
