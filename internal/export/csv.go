package export

import (
	"encoding/csv"
	"fmt"
	"os"

	"org.freethegnomes.heartRate/internal/parse"
)

func CSV(recordings []parse.RecordingFlat, filePath string) error {
	err := createDirectoryIfNotExists(filePath)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"DateTime", "BPM", "Confidence"}); err != nil {
		return err
	}

	for _, recording := range recordings {
		if err := writer.Write([]string{
			recording.DateTime.Format("01/02/06 15:04:05"),
			fmt.Sprintf("%d", recording.BPM),
			fmt.Sprintf("%d", recording.Confidence),
		}); err != nil {
			return err
		}
	}

	return nil
}
