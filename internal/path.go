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

	p := strings.Split(s, string(os.PathListSeparator))
	return p
}

func (p PathArray) CheckIfCommandExists(comm string) (bool, string) {
	for _, dir := range p {
		de, err := os.ReadDir(dir)
		if err != nil {
			//fmt.Printf("Could not open file: %e", err)
			break
		}
		fmt.Printf("Check for %s in dir %s\n", comm, dir)

		for _, entry := range de {
			fmt.Printf("%s found in dir %s\n", entry.Name(), dir)
			if comm == entry.Name() {
				cdir := filepath.Join(dir, comm)

				return true, fmt.Sprintf("%s is a valid command", cdir)
			}
		}

	}
	return false, ""
}
