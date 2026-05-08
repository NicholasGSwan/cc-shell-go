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
		cdir := filepath.Join(dir, comm)
		if _, err := os.Stat(cdir); err == nil {
			return true, fmt.Sprintf("%s is %s\n", comm, cdir)
		}

	}
	return false, ""
}
