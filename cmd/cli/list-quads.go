package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cayleygraph/cayley/quad"
	"github.com/oren/doc-api"
)

func ListQuads(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	if cmd.Parsed() {
		Admin, err := admin.New()

		if err != nil {
			panic(err)
		}

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
