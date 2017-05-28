package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cayleygraph/cayley/quad"
	"github.com/oren/doc-api"
	"github.com/oren/doc-api/bolt"
	"github.com/oren/doc-api/config"
)

func ListQuads(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	if *configPath == "" {
		*configPath = config.GetPathOfConfig()
	}

	configuration := config.ReadConf(*configPath)

	store, err := bolt.Open(configuration.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	Admin, err := admin.New(store)

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
