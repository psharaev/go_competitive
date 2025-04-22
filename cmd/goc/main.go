package main

import (
	"github.com/psharaev/go_competitive/internal/actions/args_parser"
	"github.com/psharaev/go_competitive/internal/actions/generator"
	"log"
	"os"
)

const (
	success = 0
	fail    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	args, err := args_parser.ParseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
		return fail
	}

	err = generator.Generate(args)
	if err != nil {
		log.Fatal(err)
		return fail
	}

	return success
}
