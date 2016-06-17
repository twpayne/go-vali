package vali

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
)

func ExampleService_ValidateIGC() {
	filename := "testdata/2006-06-10-XXX-3XI-01.IGC"
	igcFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer igcFile.Close()
	ctx := context.Background()
	if ok, err := NewService().ValidateIGC(ctx, filename, igcFile); ok {
		fmt.Println("OK")
	} else {
		fmt.Println(err)
	}
	// Output: OK
}
