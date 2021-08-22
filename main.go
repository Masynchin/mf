package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// main takes command title from arguments, parses Makefile
// and executes matching command.
func main() {
	commandTitle := os.Args[1]

	commands := make(map[string]string)
	err := parseMakefile(commands)
	if err != nil {
		log.Fatal(err)
	}

	command, ok := commands[commandTitle]
	if !ok {
		log.Fatal(err)
	}

	err = execCommand(command)
	if err != nil {
		log.Fatal(err)
	}
}

// parseMakefile parses Makefile and store title - command
// pairs to given commands map.
func parseMakefile(commands map[string]string) error {
	content, err := ioutil.ReadFile("Makefile")
	if err != nil {
		return err
	}

	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines {
		parts := bytes.SplitN(line, []byte(": "), 2)
		if len(parts) != 2 {
			continue
		}
		title, command := parts[0], parts[1]
		commands[string(title)] = strings.TrimSpace(string(command))
	}
	return nil
}

// execCommand executes given command
// with all output sending to stdout and stderr.
func execCommand(command string) error {
	name, args := parseCommand(command)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// parseCommand parse given command to name and args.
func parseCommand(command string) (name string, args []string) {
	args = strings.Split(command, " ")
	name, args = args[0], args[1:]

	return
}
