package main

import (
	"os"

	"github.com/trknhr/gbch"
)

func main() {
	os.Exit((&gbch.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
