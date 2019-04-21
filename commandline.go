package main

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// CommandLine declare command line parameters
type CommandLine struct {
	FilePath  string
	IsCarters bool
}

// ReadCommandLine retrieve command line parameters
func ReadCommandLine() CommandLine {
	flag.String("file", "", "path to file with HTML")
	flag.Bool("carters", false, "set to parse Carters cart")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	return CommandLine{
		FilePath:  viper.GetString("file"),
		IsCarters: viper.GetBool("carters"),
	}
}
