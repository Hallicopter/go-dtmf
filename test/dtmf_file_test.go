package test

import (
	"github.com/Hallicopter/go-dtmf/dtmf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDTMFDecoding(t *testing.T) {
	fileName := "123456654321.raw"
	decodedValue, err := dtmf.DecodeDTMFFromFile(fileName, 8000.0)
	assert.Equal(t, "123456654321", decodedValue, "Decoded value should be 123456654321")
	assert.Nil(t, err, "no error should be returned")

	fileName = "147258369.raw"
	decodedValue, err = dtmf.DecodeDTMFFromFile(fileName, 8000.0)
	assert.Equal(t, "147258369", decodedValue, "Decoded value should be 123456654321")
	assert.Nil(t, err, "no error should be returned")
}

