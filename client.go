package coremidi

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreMIDI/CoreMIDI.h>
#include <CoreServices/CoreServices.h>
*/
import "C"
import "errors"
import "fmt"

type Client struct {
	client C.MIDIClientRef
}

func NewClient(name string) (client Client, err error) {
	var clientRef C.MIDIClientRef

	stringToCFString(name, func(cfName C.CFStringRef) {
		osStatus := C.MIDIClientCreate(cfName, nil, nil, &clientRef)

		if osStatus != C.noErr {
			err = errors.New(fmt.Sprintf("%d: failed to create a client", int(osStatus)))
		} else {
			client = Client{clientRef}
		}
	})

	return
}
