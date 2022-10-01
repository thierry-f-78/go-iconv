package iconv

import "testing"

func Test(t *testing.T) {
	var err error
	var out string

	out, err = Iconv("utf-8", "iso8859-1", "\xe9\xe9\xe9\xe9\xe9\xe9\xe9\xe9\xe9\xe9\xe9\xe9\xe9")
	if err != nil {
		println(err.Error())
	}
	println(out)
}
