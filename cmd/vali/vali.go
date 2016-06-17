package main

import (
	"log"
	"os"

	"github.com/twpayne/go-vali"
	"golang.org/x/net/context"
)

func validate(ctx context.Context, s *vali.Service, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return s.ValidateIGC(ctx, filename, f)
}

func main() {
	s := vali.NewService()
	errors := false
	ctx := context.Background()
	for _, filename := range os.Args[1:] {
		if err := validate(ctx, s, filename); err != nil {
			log.Printf("%s: %v", filename, err)
			errors = true
			continue
		}
		log.Printf("%s: OK", filename)
	}
	if errors {
		os.Exit(1)
	}
}
