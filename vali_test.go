package vali_test

import (
	"context"
	"fmt"
	"os"

	"github.com/twpayne/go-vali"
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
	if status, err := vali.New().ValidateIGC(ctx, filename, igcFile); status == vali.Valid {
		fmt.Println("OK")
	} else {
		fmt.Println(err)
	}
	// Output: OK
}
