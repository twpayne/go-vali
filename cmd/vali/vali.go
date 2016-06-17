package main

import (
	"log"
	"os"

	"github.com/twpayne/go-vali"
	"golang.org/x/net/context"
)

func validate(ctx context.Context, s *vali.Service, filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()
	return s.ValidateIGC(ctx, filename, f)
}

func main() {
	s := vali.New()
	errors := false
	ctx := context.Background()
	for _, filename := range os.Args[1:] {
		if ok, err := validate(ctx, s, filename); !ok {
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
