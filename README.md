# go-iconv

[![GoDoc](https://pkg.go.dev/badge/github.com/thierry-f-78/go-iconv)](https://pkg.go.dev/github.com/thierry-f-78/go-iconv)

This library convert text from charset to another one using standard libc iconv. It contains only one function. There's an example:

```go
// convert ISO8859-1 to UTF-8, without transliteration
out, err = Iconv("utf-8", "iso8859-1", "hello, \xe9\xe9, hello \xe9\xe9", false)
if err != nil {
	println("error:", err.Error())
} else {
	println(out)
}
```

For more information about iconv:

`man 3 iconv`
