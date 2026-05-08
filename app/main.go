package main

import (
	"fmt"

	repl "github.com/codecrafters-io/shell-starter-go/internal"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	// p := repl.GetPathArray()
	// for _, v := range p {
	// 	fmt.Println(v)
	// }

	// fmt.Println(p.CheckIfCommandExists("badcommand"))
	// fmt.Println(p.CheckIfCommandExists("exercism.exe"))

	//fmt.Println(os.Getenv("PATH"))
	repl.StartRepl()

}
