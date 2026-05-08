package repl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PathArray []string

func GetPathArray() PathArray {
	s := os.Getenv("PATH")
	p := strings.Split(s, ";")
	return p
}

func (p PathArray) CheckIfCommandExists(comm string) bool {
	for _, dir := range p {
		de, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("Could not open file: %e", err)
			break
		}
		for _, entry := range de {
			if comm == entry.Name() && !entry.Type().IsDir() {
				cdir := filepath.Join(dir, comm)
				fmt.Printf("%s is a valid command", cdir)
				return true
			}
		}

	}
	return false
}
