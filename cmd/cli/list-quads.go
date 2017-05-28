package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cayleygraph/cayley/quad"
	"github.com/oren/doc-api"
	"github.com/oren/doc-api/config"
)

func ListQuads(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	if *configPath == "" {
		*configPath = getPathOfConfig()
	}

	configuration := config.ReadConf(*configPath)
	fmt.Println(configuration)

	// TODO: pass configuration.DbPath
	Admin, err := admin.New()

	if err != nil {
		panic(err)
	}

	var quads []quad.Quad
	quads, err = Admin.AllQuads()
	admin.CheckErr(err)
	printQuads(quads)
}

func printQuads(quads []quad.Quad) {
	fmt.Println("\n==== All quads ====")

	for _, q := range quads {
		fmt.Println(q)
	}
}
