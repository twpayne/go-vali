package main

import (
	"context"
	"fmt"
	"os"

	vali "github.com/twpayne/go-vali"
)

func validate(ctx context.Context, s *vali.Service, filename string) (vali.Status, error) {
	f, err := os.Open(filename)
	if err != nil {
		return vali.Unknown, err
	}
	defer f.Close()
	return s.ValidateIGC(ctx, filename, f)
}

func main() {
	s := vali.New()
	worstStatus := vali.Valid
	ctx := context.Background()
	for _, filename := range os.Args[1:] {
		status, err := validate(ctx, s, filename)
		switch status {
		case vali.Valid:
			fmt.Printf("%s: %s\n", filename, status)
		case vali.Invalid, vali.Unknown:
			fmt.Printf("%s: %s: %s\n", filename, status, err)
		}
		if status > worstStatus {
			worstStatus = status
		}
	}
	os.Exit(int(worstStatus))
}
