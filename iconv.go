package iconv

// #cgo darwin LDFLAGS: -liconv
// #cgo freebsd LDFLAGS: -liconv
// #cgo windows LDFLAGS: -liconv
// #include <stdlib.h>
// char *do_iconv(char *to, char *from, char *text);
import "C"

import "errors"
import "unsafe"

func Iconv(to string, from string, text string)(string, error) {
	var c_from *C.char
	var c_to *C.char
	var c_text *C.char
	var c_out *C.char
	var out string

	c_from = C.CString(from)
	defer C.free(unsafe.Pointer(c_from))

	c_to = C.CString(to)
	defer C.free(unsafe.Pointer(c_to))

	c_text = C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_out = C.do_iconv(c_to, c_from, c_text)
	if c_out == nil {
		return "", errors.New("Can't convert text")
	}
	out = C.GoString(c_out)
	C.free(unsafe.Pointer(c_out))

	return out, nil
}
