package internal

import (
	"log"
	"os"
)

type Scope struct {
	ips []string
}

func NewScope(ips []string, path string) (Scope, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return Scope{}, err
	}
	return Scope{ips: ips}, nil
}

func (s *Scope) WriteScope(file string) error {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if s.ips == nil {
		//accept all ips
		_, err := f.WriteString("10.10.10.10/0")
		if err != nil {
			return err
		}
	} else {
		for _, ip := range s.ips {
			_, err := f.WriteString(ip + "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
