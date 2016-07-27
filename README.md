# go-vali

[![Build Status](https://travis-ci.org/twpayne/go-vali.svg?branch=master)](https://travis-ci.org/twpayne/go-vali)
[![GoDoc](https://godoc.org/github.com/twpayne/go-vali?status.svg)](https://godoc.org/github.com/twpayne/go-vali)
[![Report Card](https://goreportcard.com/badge/github.com/twpayne/go-vali)](https://goreportcard.com/report/github.com/twpayne/go-vali)

Package vali provides a client interface to CIVL's Open Validation Server.
See http://vali.fai-civl.org/webservice.html.

Example:

```go
func ExampleNew_ValidateIGC() {
	filename := "testdata/2006-06-10-XXX-3XI-01.IGC"
	igcFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer igcFile.Close()
	ctx := context.Background()
	if status, err := New().ValidateIGC(ctx, filename, igcFile); status == Valid {
		fmt.Println("OK")
	} else {
		fmt.Println(err)
	}
	// Output: OK
}
```

[License](LICENSE)
