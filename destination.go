package coremidi

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation
#include <CoreMIDI/CoreMIDI.h>
*/
import "C"

type Destination struct {
	endpoint C.MIDIEndpointRef
	*Object
}
