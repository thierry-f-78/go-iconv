package iconv

// #cgo darwin LDFLAGS: -liconv
// #cgo freebsd LDFLAGS: -liconv
// #cgo windows LDFLAGS: -liconv
// #include <stdlib.h>
// char *do_iconv(char *to, char *from, char *text, int transliterate);
import "C"

import "errors"
import "unsafe"

// to is the target charset, from is the source charset text is the text to convert
// to. If transliterate is true, the function try to convert text even if the target
// charset has no corresponding char. Exemple, utf-8 "é" to ascii give "'e". If
// transliterate is false the conversion fails.
func Iconv(to string, from string, text string, transliterate bool)(string, error) {
	var c_from *C.char
	var c_to *C.char
	var c_text *C.char
	var c_out *C.char
	var c_transliterate C.int
	var out string

	c_from = C.CString(from)
	defer C.free(unsafe.Pointer(c_from))

	c_to = C.CString(to)
	defer C.free(unsafe.Pointer(c_to))

	c_text = C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	if transliterate {
		c_transliterate = 1
	} else {
		c_transliterate = 0
	}

	c_out = C.do_iconv(c_to, c_from, c_text, c_transliterate)
	if c_out == nil {
		return "", errors.New("Can't convert text")
	}
	out = C.GoString(c_out)
	C.free(unsafe.Pointer(c_out))

	return out, nil
}
