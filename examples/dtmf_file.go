package main

import (
	"fmt"
	"go-dtmf/dtmf"
	"os"
	"time"
)

func main() {
	// Testing using file mono, 8000hz file for DTMF tone of 1
	fileName := "test/8_filtered.raw"
	char, err := getSingleDTMFValueFromFile(fileName, 8000.0, 50*time.Millisecond)
	if err != nil && err != dtmf.NoDigitFoundError {
		fmt.Println("There is an error", err)
		return
	}
	fmt.Println("Decoded character is ", string(char))

	// Testing using file mono, 8000hz file for DTMF tone of 1
	fileName = "test/dtmf8.raw"
	char, err = getSingleDTMFValueFromFile(fileName, 8000.0, 70*time.Millisecond)
	if err != nil {
		fmt.Println("There is an error", err)
		return
	}
	fmt.Println("Decoded character is ", string(char))

	//fileName = "test/8197284321.raw"
	//valueString, err := getMultipleDTMFValueFromFile(fileName, 8000.0, 50*time.Millisecond, 2000)
	//if err != nil {
	//	fmt.Println("There is an error", err)
	//	return
	//}
	//fmt.Println("Decoded character is ", valueString)
}

func getSingleDTMFValueFromFile(filepath string, rate float64, minDur time.Duration) (rune, error) {
	audioBytes, err := os.ReadFile(filepath)
	if err != nil {
		return '-', err
	}

	d := dtmf.NewDTMFStruct(rate, minDur, audioBytes)

	char, err := d.DecodeDTMFFromBytes(20)
	if err != nil {
		return '-', err
	}
	return char, nil
}

func getMultipleDTMFValueFromFile(filepath string, rate float64, minDur time.Duration, resolution int) (string, error) {
	bytesRead := 0
	valueString := ""

	audioBytes, err := os.ReadFile(filepath)
	if err != nil {
		return "NA", err
	}

	for bytesRead <= len(audioBytes) && bytesRead+resolution < len(audioBytes) {
		d := dtmf.NewDTMFStruct(rate, minDur, audioBytes[bytesRead:bytesRead+resolution])
		bytesRead += resolution
		char, err := d.DecodeDTMFFromBytes(10)
		if err != nil && err != dtmf.NoDigitFoundError {
			return "NA", err
		}
		valueString += string(char)
	}
	return valueString, nil
}