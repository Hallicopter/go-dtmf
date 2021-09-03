# go-dtmf, the simplest way to decode DTMF audio in golang. 

### About
This library provides two high level API to decode DTMF audio or byte slice. 
It uses the Goertzer algorithm. 
It was initially built on [goertzel](https://github.com/CyCoreSystems/goertzel) by CyCoreSystems, but has been since ported to use [go-dsp](https://github.com/samuel/go-dsp).
The underlying principle remains unchanged.

### Examples
The examples can be run from the examples folder.

``go run examples/dtmf_file.go``

### Docs

### dtmf

    import "go-dtmf/dtmf"


### Usage

#### func  DecodeDTMFValueFromFile

```go
func DecodeDTMFValueFromFile(filepath string, rate float64) (string, error)
```
DecodeDTMFValueFromFile Expects raw audio as the input, gives the decoded DTMF
string as output.

#### type DTMF

```go
type DTMF struct {
	DecodedValue string
}
```


#### func  NewDTMFStruct

```go
func NewDTMFStruct(sampleRate float64, audioBytes []byte) DTMF
```
NewDTMFStruct Creates and initialises a struct that can be used to call the
decoding method.

#### func (*DTMF) DecodeDTMFFromBytes

```go
func (dtmf *DTMF) DecodeDTMFFromBytes() (err error)
```
DecodeDTMFFromBytes This decodes the audio bytes and saves the value in
DTMF.DecodedValue
