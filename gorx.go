package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version = "0.1.0"

	usage = `Usage: gorx [subcommand] (options) [args]
`

	help = `=================================================
 gorx
 Copyright 2018 Christopher Simpkins
 MIT License

 Source: https://github.com/go-rillas/gorx
=================================================
 Usage:
   $ gorx [subcommand] (options) [args]

 Options:
  -h, --help           Application help
      --usage          Application usage
  -v, --version        Application version
`
)

var versionShort, versionLong, helpShort, helpLong, usageLong *bool

func init() {
	// define available command line flags
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")
}

func main() {
	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		fmt.Printf("gorx v%s\n", version)
		os.Exit(0)
	case *helpShort, *helpLong:
		fmt.Println(help)
		os.Exit(0)
	case *usageLong:
		fmt.Println(usage)
		os.Exit(0)
	}

	args := flag.Args()

	switch args[0] {
	case "gor":
		fmt.Println("gor")
	case "go":
		fmt.Println("go")
	}
}
