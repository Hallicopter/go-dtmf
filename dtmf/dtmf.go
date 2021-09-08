package dtmf

import (
	"bytes"
	"errors"
	utils "github.com/Hallicopter/go-dtmf/utils/dtmf"
	"github.com/caicloud/nirvana/log"
	"io"
	"os"
)

// DecodeDTMFFromBytes
// This decodes the audio bytes and saves the value in DTMF.DecodedValue
// The wiggleRoom value is recommended to be between 5-15.
// For shorter, sharper, faster DTMF audios, a wiggleRoom of 5 would be good.
// For longer, more continuous DTMF audios, a higher wiggleRoom will prevent false repeats.
func DecodeDTMFFromBytes(audioBytes []byte, rate float64, wiggleRoom int) (string, error) {
	if len(audioBytes) == 0 {
		return "", errors.New("audio in the dtmf structure contains no bytes")
	}

	var dtmfOutput string
	sampleRate := int(rate)
	blockSize := 205 * sampleRate / 8000
	window := blockSize / 4
	dt := utils.NewStandard(sampleRate, blockSize)
	lastKey := -1
	keyCount := 0
	samples := make([]float32, blockSize)

	rd := bytes.NewReader(audioBytes)

	buf := make([]byte, window*2)

	for {
		_, err := rd.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		copy(samples, samples[window:])

		si := len(samples) - window
		for i := 0; i < len(buf); i += 2 {
			s := float32(int16(buf[i])|(int16(buf[i+1])<<8)) / 32768.0
			samples[si] = s
			si++
		}

		if k, t := dt.Feed(samples); k == lastKey && t > 0.0 {
			keyCount++
			if keyCount == wiggleRoom {
				dtmfOutput += string(utils.Keypad[k])
			}
		} else {
			lastKey = k
			keyCount = 0
		}
	}

	return dtmfOutput, nil
}

// DecodeDTMFFromFile
// Expects raw audio as the input, gives the decoded DTMF string as output.
// The wiggleRoom value is recommended to be between 5-15.
// For shorter, sharper, faster DTMF audios, a wiggleRoom of 5 would be good.
// For longer, more continuous DTMF audios, a higher wiggleRoom will prevent false repeats.
func DecodeDTMFFromFile(filepath string, rate float64, wiggleRoom int) (string, error) {
	audioBytes, err := os.ReadFile(filepath)
	if err != nil {
		return "N/A", err
	}

	decodedValue, err := DecodeDTMFFromBytes(audioBytes, rate, wiggleRoom)
	if err != nil {
		return "N/A", err
	}
	return decodedValue, nil
}
