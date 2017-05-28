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

	if cmd.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

		if *configPath == "" {
			*configPath = "/tmp/config.json"
		}

		configuration := config.ReadConf(*configPath)
		fmt.Println(configuration)

		var quads []quad.Quad
		quads, err = Admin.AllQuads()
		admin.CheckErr(err)
		printQuads(quads)
	}
}

func printQuads(quads []quad.Quad) {
	fmt.Println("\n==== All quads ====")

	for _, q := range quads {
		fmt.Println(q)
	}
}
