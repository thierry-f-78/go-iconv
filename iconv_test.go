package iconv

import "testing"

func Test(t *testing.T) {
	var err error
	var out string
	var expect_01 string = "hello, éé, hello éé"
	var expect_02 string = "a'eb'ec"

	/* test sio => utf8 conversion */
	out, err = Iconv("utf-8", "iso8859-1", "hello, \xe9\xe9, hello \xe9\xe9", false)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if out != expect_01 {
		t.Errorf("Expect %q, got %q", expect_01, out)
	}

	/* Test wrong utf8 sequence, expect error */
	_, err = Iconv("iso8859-1", "utf-8", string([]byte{0b11100010, 0b10100011, 0b10100011}), false)
	if err == nil {
		t.Errorf("Expect an error, got success")
	}

	/* Test without transliteration */
	_, err = Iconv("ascii", "iso8859-1", "a\xe9b\xe9c", false)
	if err == nil {
		t.Errorf("Expect an error, got success")
	}

	/* Test with transliteration */
	out, err = Iconv("ascii", "iso8859-1", "a\xe9b\xe9c", true)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if out != expect_02 {
		t.Errorf("Expect %q, got %q", expect_02, out)
	}
}
