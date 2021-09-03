# go-dtmf, the simplest way to decode DTMF audio in golang. 


This library provides two high level API to decode DTMF audio or byte slice. 
It uses the Goertzel algorithm. 

### Examples
The examples can be run from the examples folder.

``go run examples/dtmf_file.go``

## Docs

### dtmf

    import "github.com/Hallicopter/go-dtmf/dtmf"


### Usage

#### func  DecodeDTMFFromBytes

```go
func DecodeDTMFFromBytes(audioBytes []byte, rate float64) (string, error)
```
DecodeDTMFFromBytes This decodes the audio bytes and saves the value in
DTMF.DecodedValue

#### func  DecodeDTMFFromFile

```go
func DecodeDTMFFromFile(filepath string, rate float64) (string, error)
```
DecodeDTMFromFile Expects raw audio as the input, gives the decoded DTMF
string as output.



## Credits
It was initially built on [goertzel](https://github.com/CyCoreSystems/goertzel) by CyCoreSystems, but has been since ported to use a modified version of [go-dsp](https://github.com/samuel/go-dsp).
The underlying principle remains unchanged.
