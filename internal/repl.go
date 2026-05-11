package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		"pwd": {
			name:        "pwd",
			description: "Print current working directory",
			ctype:       "builtin",
			callback:    commandPwd,
		},
		"cd": {
			name:        "cd",
			description: "Change current working directory",
			ctype:       "builtin",
			callback:    commandCd,
		},
	}
	pathArray = GetPathArray()
}

func commandExit(sArr args) error {
	os.Exit(0)
	return nil
}

func commandEcho(sArr args) error {
	fmt.Print(strings.Join(sArr, " ") + "\n")
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

func commandPwd(sArr args) error {
	currDir, err := filepath.Abs("")
	if err != nil {
		return err
	}
	fmt.Println(currDir)
	return nil
}

func commandCd(sArr args) error {
	newDir := sArr[0]
	err := os.Chdir(newDir)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", newDir)
		return err
	}
	fmt.Println(newDir)
	return nil
}

func StartRepl() {
	for {
		fmt.Print("$ ")
		commandString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		strArr := strings.Split(strings.TrimSpace(commandString), " ")

		command := strings.TrimSpace(strArr[0])
		strArr = strArr[1:]

		if comm, ok := commands[command]; ok {
			comm.callback(strArr)
		} else if ok, _, _ := pathArray.CheckIfCommandExists(command); ok {
			//fmt.Println("Preparing command: ", commStr, strArr)
			cmd := exec.Command(command, strArr...)
			if cmd.Err == nil {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println("Command failed to run: ", err)
				}
				//fmt.Println("the process state is: ", cmd.ProcessState)
			} else {
				fmt.Println(cmd.Err.Error())
			}
		} else {
			fmt.Println(command + ": command not found")
		}

	}
}
