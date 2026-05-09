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

func (p PathArray) CheckIfCommandExists(comm string) (bool, string, string) {
	for _, dir := range p {
		cdir := filepath.Join(dir, comm)
		if fi, err := os.Stat(cdir); err == nil && !fi.IsDir() && fi.Mode()%2 != 0 {
			//fmt.Printf("File perms for %s is %#o\n", cdir, )
			return true, comm, cdir
		}

	}
	return false, "", ""
}

func (p PathArray) CommandTypeFunc(comm string) (bool, string) {
	exists, cmd, cmdDir := p.CheckIfCommandExists(comm)
	return exists, fmt.Sprintf("%s is %s\n", cmd, cmdDir)
}

func (p PathArray) GetCommandString(comm string) (bool, string) {
	exists, comm1, cdir := p.CheckIfCommandExists(comm)
	return exists, filepath.Join(cdir, comm1)
}
