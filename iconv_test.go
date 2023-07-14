package iconv

import "testing"

func Test(t *testing.T) {
	var err error
	var out string
	var expect_01 string = "hello, éé, hello éé"

	/* test sio => utf8 conversion */
	out, err = Iconv("utf-8", "iso8859-1", "hello, \xe9\xe9, hello \xe9\xe9")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if out != expect_01 {
		t.Errorf("Expect %q, got %q", expect_01, out)
	}

	/* Test wrong utf8 sequence, expect error */
	_, err = Iconv("iso8859-1", "utf-8", string([]byte{0b11100010, 0b10100011, 0b10100011}))
	if err == nil {
		t.Errorf("Expect an error, got success")
	}
}
