package dtmf

import (
	"bytes"
	"errors"
	"github.com/caicloud/nirvana/log"
	utils "go-dtmf/utils/dsp/dtmf"
	"io"
	"os"
)

// DecodeDTMFFromBytes
// This decodes the audio bytes and saves the value in DTMF.DecodedValue
func DecodeDTMFFromBytes(audioBytes []byte, rate float64) (string, error) {
	if len(audioBytes) == 0 {
		return "", errors.New("audio in the dtmf structure contains no bytes")
	}

	var dtmfOutput string
	sampleRate := 8000
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
			if keyCount == 12 {
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
func DecodeDTMFFromFile(filepath string, rate float64) (string, error) {
	audioBytes, err := os.ReadFile(filepath)
	if err != nil {
		return "N/A", err
	}

	decodedValue, err := DecodeDTMFFromBytes(audioBytes, rate)
	if err != nil {
		return "N/A", err
	}
	return decodedValue, nil
}
