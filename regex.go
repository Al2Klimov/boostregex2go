package boostregex2go

import (
	"io"
	"reflect"
	"runtime"
	"unsafe"
)

/*
#include "libcxx/regex.h"
// CompileRegex
// FreeRegex
// MatchesSomewhere
*/
import "C"

type OOM struct {
}

var _ error = OOM{}

func (OOM) Error() string {
	return "out of memory"
}

type BadPattern struct {
}

var _ error = BadPattern{}

func (BadPattern) Error() string {
	return "bad pattern"
}

type Regex struct {
	rgx unsafe.Pointer
}

var _ io.Closer = (*Regex)(nil)

func (r *Regex) Close() error {
	C.FreeRegex(rgxPtr64(r))
	return nil
}

func (r *Regex) MatchesSomewhere(subject []byte) (bool, error) {
	defer runtime.KeepAlive(subject)
	start, end := bytesToCharRange(subject)

	switch C.MatchesSomewhere(rgxPtr64(r), start, end) {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, OOM{}
	}
}

func NewRegex(pattern []byte) (*Regex, error) {
	rgx := &Regex{}

	defer runtime.KeepAlive(pattern)
	start, end := bytesToCharRange(pattern)

	switch C.CompileRegex(start, end, rgxPtr64(rgx)) {
	case 0:
		return rgx, nil
	case 2:
		return nil, BadPattern{}
	default:
		return nil, OOM{}
	}
}

func bytesToCharRange(b []byte) (C.uint64_t, C.uint64_t) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return C.uint64_t(sh.Data), C.uint64_t(sh.Data + uintptr(sh.Len))
}

func rgxPtr64(p *Regex) C.uint64_t {
	return C.uint64_t(uintptr(unsafe.Pointer(p)))
}
