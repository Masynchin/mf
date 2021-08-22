package main

import (
	"testing"
)

func TestParseMakefile(t *testing.T) {
	commands := make(map[string]string)
	err := parseMakefile(commands)
	if err != nil {
		t.Fatalf("Couldn't parse Makefile commands")
	}

	command, ok := commands["echonumber"]
	if !ok {
		t.Fatalf("Couldn't find `echonumber` command")
	} else if command != "echo 123" {
		t.Fatalf("Wrong `echonumber` command content")
	}
}

func TestExecCommand(t *testing.T) {
	err := execCommand("echo PASS ECHO 123")
	if err != nil {
		t.Fatalf("Error while executing `echo PASS ECHO 123` command")
	}
}
