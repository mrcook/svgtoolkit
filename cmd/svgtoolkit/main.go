package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"

	"github.com/mrcook/svgtoolkit"
	"github.com/mrcook/svgtoolkit/cmd/svgtoolkit/geopattern"
)

func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "Print current version")

	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&geopattern.PatternCmd{}, "")

	flag.Parse()

	if version {
		fmt.Println(svgtoolkit.Version)
		os.Exit(0)
	}

	ctx := context.Background()
	subcommands.Execute(ctx)
}
