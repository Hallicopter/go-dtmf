package main

import (
	"fmt"

	"github.com/speechcatch/go-dtmf/dtmf"
)

func main() {
	// Testing using file mono, 8000hz files

	fileName := "test/123456654321.raw"
	valueString, err := dtmf.DecodeDTMFFromFile(fileName, 8000.0, 12, 0.0)
	if err != nil {
		fmt.Println("There is an error", err)
		return
	}
	fmt.Println("Decoded character is", valueString)

	fileName = "test/147258369.raw"
	valueString, err = dtmf.DecodeDTMFFromFile(fileName, 8000.0, 12, 0.0)
	if err != nil {
		fmt.Println("There is an error", err)
		return
	}
	fmt.Println("Decoded character is", valueString)
}
