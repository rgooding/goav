// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/error.h>
//#include <stdlib.h>
//int avStrError(int errnum, char* errbuf, int errbuf_size) { return av_strerror(errnum, errbuf, errbuf_size); }
import "C"
import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

const (
	AvErrorEOF    = -('E' | ('O' << 8) | ('F' << 16) | (' ' << 24))
	AvErrorEAGAIN = -35
)

func ErrorFromCode(code int) error {
	const bufSize = 64
	myStr := strings.Repeat(" ", bufSize)
	cStr := C.CString(myStr)
	defer C.free(unsafe.Pointer(cStr))

	if code >= 0 {
		return nil
	}
	res := C.avStrError(C.int(code), cStr, bufSize)
	if res < 0 {
		return fmt.Errorf("unknown error %d", code)
	}
	return errors.New(C.GoString(cStr))
}
