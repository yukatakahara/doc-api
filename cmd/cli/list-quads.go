package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cayleygraph/cayley/quad"
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

	// var quads []interface{}
	var quads []quad.Quad
	adminService := &bolt.AdminService{Store: store}
	quads, err = adminService.Quads()

	if err != nil {
		log.Fatal(err)
	}

	printQuads(quads)
}

// func printQuads(quads []interface{}) {
func printQuads(quads []quad.Quad) {
	fmt.Println("\n==== All quads ====")

	for _, q := range quads {
		fmt.Println(q)
	}
}
