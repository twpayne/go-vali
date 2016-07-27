package main

import (
	"log"
	"os"

	"github.com/twpayne/go-vali"
	"golang.org/x/net/context"
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
	errors := false
	ctx := context.Background()
	for _, filename := range os.Args[1:] {
		status, err := validate(ctx, s, filename)
		switch status {
		case vali.Valid:
			log.Printf("%s: %s", filename, status)
		case vali.Invalid, vali.Unknown:
			log.Printf("%s: %s: %s", filename, status, err)
			errors = true
		}
	}
	if errors {
		os.Exit(1)
	}
}
