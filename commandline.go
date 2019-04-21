package main

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// CommandLine declare command line parameters
type CommandLine struct {
	FilePath     string
	TemplatePath string
	IsCarters    bool
}

// ReadCommandLine retrieve command line parameters
func ReadCommandLine() CommandLine {
	flag.String("cart", "", "path to file with cart's HTML")
	flag.String("template", "", "path to file with Qwintry HTML template")
	flag.Bool("carters", false, "set to parse Carters cart")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	return CommandLine{
		FilePath:     viper.GetString("cart"),
		TemplatePath: viper.GetString("template"),
		IsCarters:    viper.GetBool("carters"),
	}
}
