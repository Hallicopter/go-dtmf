package dtmf

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/caicloud/nirvana/log"
	"goertzel"
	"time"
)

// Standard frequencies for DTMF as per RFC.
var (
	Keypad = []rune{
		'1', '2', '3', 'A',
		'4', '5', '6', 'B',
		'7', '8', '9', 'C',
		'*', '0', '#', 'D',
	}
	StdLowFreq        = []float64{697, 770, 852, 941}
	StdHighFreq       = []float64{1209, 1336, 1477, 1633}
	NoDigitFoundError = errors.New("no digit found")
)

type DTMF struct {
	audioBytes      []byte
	sampleRate      float64
	DecodedValue    rune
	minimumDuration time.Duration
}

func NewDTMFStruct(sampleRate float64, minDur time.Duration, audioBytes []byte) DTMF {
	return DTMF{
		audioBytes:      audioBytes,
		sampleRate:      sampleRate,
		minimumDuration: minDur,
	}
}

func (dtmf *DTMF) DecodeDTMFFromBytes(tolerance int) (dtmfChar rune, err error) {
	var found bool
	var row, col int
	var freq float64
	var foundLowFreq, foundHighFreq bool
	var toleranceSlice = make([]float64, tolerance*2+1)

	// Looks for the dual tones of the Dual Tones Multi Frequency fame.

	for i := -tolerance; i<tolerance; i++ {
		toleranceSlice[i+tolerance] = float64(i)
	}


	for _, tolerance := range toleranceSlice {
		for row, freq = range StdLowFreq {
			found, err = goertzel.DetectTone(context.Background(), freq+tolerance, dtmf.sampleRate, dtmf.minimumDuration, bytes.NewReader(dtmf.audioBytes))
			if found {
				fmt.Println("Founder lower freq", freq+tolerance, "row", row)
				foundLowFreq = true
				break
			}
		}
		if foundLowFreq {
			break
		}
	}

	for _, tolerance := range toleranceSlice {
		for col, freq = range StdHighFreq {
			found, err = goertzel.DetectTone(context.Background(), freq+tolerance, dtmf.sampleRate, dtmf.minimumDuration, bytes.NewReader(dtmf.audioBytes))
			if found {
				fmt.Println("Founder higher freq", freq+tolerance, "col", col)
				foundHighFreq = true
				break
			}
		}
		if foundHighFreq {
			break
		}
	}

	if foundLowFreq && foundHighFreq {
		return Keypad[row*4+col], nil
	}

	log.Infof("No character found in DTMF")
	return '-', NoDigitFoundError
}
