package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Print("$ ")
		commandString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		strArr := strings.Split(commandString, " ")

		command := strings.TrimSpace(strArr[0])
		strArr = strArr[1:]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Print(strings.Join(strArr, " "))
		default:
			fmt.Println(command + ": command not found")
		}

	}

}
