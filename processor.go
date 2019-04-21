package main

import (
	"os"

	"github.com/Toshik1978/cart_to_qwintry/cartparsers"
	"github.com/pkg/errors"
)

// Process process command line and do all actions
func Process(cmdLine CommandLine) error {
	file, err := os.Open(cmdLine.FilePath)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}

	var parser cartparsers.CartParser
	if cmdLine.IsCarters {
		parser = cartparsers.NewCartersParser()
	}
	if parser == nil {
		return errors.New("failed get parser")
	}

	result, err := parser.Parse(file)
	if err != nil {
		return errors.Wrap(err, "failed to parse cart")
	}

	return SaveTemplate(cmdLine.TemplatePath, result)
}
