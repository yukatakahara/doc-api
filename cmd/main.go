package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
)

func main() {
	// Some globally applicable things
	graph.IgnoreMissing = true
	graph.IgnoreDuplicates = true

	// File for your new BoltDB. Use path to regular file and not temporary in the real world
	t := getTempfileName()
	fmt.Printf("%v\n", t)

	defer os.Remove(t)                 // clean up
	store := initializeAndOpenGraph(t) // initialize the graph

	createAdmin(store) // add quads to the graph

	// countOuts(store, "robertmeta")
	// lookAtOuts(store, "robertmeta")
	// lookAtIns(store, "robertmeta")

	// lookAtOuts(store, "jorgent")
	// lookAtIns(store, "jorgent")

	// lookAtFriendsOfFriends(store, "barakmich")

}

func initializeAndOpenGraph(atLoc string) *cayley.Handle {
	// Initialize the database
	graph.InitQuadStore("bolt", atLoc, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", atLoc, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return store
}

func getTempfileName() string {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	return tmpfile.Name()
}

func createAdmin(store *cayley.Handle) {
	store.AddQuad(quad.MakeRaw("41234214", "is_a", "admin", ""))
	store.AddQuad(quad.MakeRaw("41234214", "email", "foo@gmail.com", ""))
	store.AddQuad(quad.MakeRaw("41234214", "hashed_password", "aqjxkjx", ""))
}
