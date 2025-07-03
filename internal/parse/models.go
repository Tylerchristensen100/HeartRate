package parse

import "time"

type Recording struct {
	DateTime string `json:"dateTime"`
	Value    struct {
		BPM        int `json:"bpm"`
		Confidence int `json:"confidence"`
	} `json:"value"`
}

type RecordingFlat struct {
	DateTime   time.Time `json:"dateTime"`
	BPM        int       `json:"bpm"`
	Confidence int       `json:"confidence"`
}

func ConvertRecording(recordings []Recording) ([]RecordingFlat, error) {
	var flatRecordings []RecordingFlat
	for _, recording := range recordings {
		t, err := time.Parse("01/02/06 15:04:05", recording.DateTime)
		if err != nil {
			panic(err)
		}

		flatRecording := RecordingFlat{
			DateTime:   t,
			BPM:        recording.Value.BPM,
			Confidence: recording.Value.Confidence,
		}
		flatRecordings = append(flatRecordings, flatRecording)
	}
	return flatRecordings, nil
}
