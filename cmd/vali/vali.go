package main

import (
	"log"
	"os"

	"github.com/twpayne/go-vali"
)

func validate(s *vali.Service, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return s.IGC(filename, f)
}

func main() {
	s := vali.NewService()
	errors := false
	for _, filename := range os.Args[1:] {
		if err := validate(s, filename); err != nil {
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
