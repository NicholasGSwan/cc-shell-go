package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type args []string

type cliCommand struct {
	name        string
	description string
	ctype       string
	callback    func(args) error
}

var commands map[string]cliCommand
var pathArray PathArray

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the shell",
			ctype:       "builtin",
			callback:    commandExit,
		},
		"echo": {
			name:        "echo",
			description: "Return the text following the command",
			ctype:       "builtin",
			callback:    commandEcho,
		},
		"type": {
			name:        "type",
			description: "Describe the type of a command",
			ctype:       "builtin",
			callback:    commandType,
		},
	}
	pathArray = GetPathArray()
}

func commandExit(sArr args) error {
	os.Exit(0)
	return nil
}

func commandEcho(sArr args) error {
	fmt.Print(strings.Join(sArr, " "))
	return nil
}

func commandType(sArr args) error {
	c := strings.TrimSpace(sArr[0])
	if comm, ok := commands[c]; ok {
		fmt.Printf("%s is a shell %s\n", comm.name, comm.ctype)
		return nil
	} else if ok, s := pathArray.CommandTypeFunc(c); ok {
		fmt.Print(s)
		return nil
	} else {
		fmt.Printf("%s: not found\n", c)
		return nil
	}
}

func StartRepl() {
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
		if strArr[len(strArr)-1] == "" {
			strArr = strArr[:len(strArr)-1]
		}
		if comm, ok := commands[command]; ok {
			comm.callback(strArr)
		} else if ok, _, commStr := pathArray.CheckIfCommandExists(command); ok {

			cmd := exec.Command(commStr, strArr...)
			if cmd.Err == nil {
				err := cmd.Run()
				if err != nil {
					fmt.Println("Command failed to run: ", err)
				}
			} else {
				fmt.Println(cmd.Err.Error())
			}
		} else {
			fmt.Println(command + ": command not found")
		}

	}
}
