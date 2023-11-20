package main

import (
	"context"
	"fmt"
	"os"

	vali "github.com/twpayne/go-vali"
)

func validate(ctx context.Context, s *vali.Service, filename string) (vali.Status, *vali.Response, error) {
	f, err := os.Open(filename)
	if err != nil {
		return vali.StatusUnknown, nil, err
	}
	defer f.Close()
	return s.ValidateIGC(ctx, filename, f)
}

func main() {
	s := vali.New()
	worstStatus := vali.StatusValid
	ctx := context.Background()
	for _, filename := range os.Args[1:] {
		status, _, err := validate(ctx, s, filename)
		switch status {
		case vali.StatusValid:
			fmt.Printf("%s: %s\n", filename, status)
		case vali.StatusInvalid:
			fmt.Printf("%s: %s: %s\n", filename, status, err)
			if worstStatus < 1 {
				worstStatus = 1
			}
		case vali.StatusUnknown:
			fmt.Printf("%s: %s: %s\n", filename, status, err)
			if worstStatus < 2 {
				worstStatus = 2
			}
		}
	}
	os.Exit(int(worstStatus))
}
