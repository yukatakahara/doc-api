package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

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
	// t := getTempfileName()
	// fmt.Printf("%v\n", t)

	// defer os.Remove(t)                 // clean up
	// store := initializeAndOpenGraph(t) // initialize the graph

	dbFile := flag.String("db", "db.boltdb", "BoltDB file")
	flag.Parse()

	store := initializeAndOpenGraph(dbFile)

	createAdmin(store) // add quads to the graph

	countOuts(store, "41234214")
	lookAtOuts(store, "41234214")
	lookAtIns(store, "41234214")

	// lookAtOuts(store, "jorgent")
	// lookAtIns(store, "jorgent")

	// lookAtFriendsOfFriends(store, "barakmich")

}

func initializeAndOpenGraph(dbFile *string) *cayley.Handle {
	graph.InitQuadStore("bolt", *dbFile, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", *dbFile, nil)
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

// countOuts ... well, counts Outs
func countOuts(store *cayley.Handle, to string) {
	p := cayley.StartPath(store, quad.Raw(to)).Out().Count()
	fmt.Printf("\n\ncountOuts for %s: ", to)
	p.Iterate(nil).EachValue(store, func(v quad.Value) {
		fmt.Printf("%d\n", quad.NativeOf(v))
	})
	fmt.Printf("============================================\n")
}

// lookAtOuts looks at the outbound links from the "to" node
func lookAtOuts(store *cayley.Handle, to string) {
	p := cayley.StartPath(store, quad.Raw(to)) // start from a single node, but we could start from multiple

	// this gives us a path with all the output predicates from our starting point
	p = p.Tag("subject").OutWithTags([]string{"predicate"}).Tag("object")

	fmt.Printf("\nlookAtOuts: subject (%s) -predicate-> object\n", to)
	fmt.Printf("============================================\n")

	p.Iterate(nil).TagValues(nil, func(m map[string]quad.Value) {
		fmt.Printf("%s `%s`-> %s\n", m["subject"], m["predicate"], m["object"])
	})
}

// lookAtIns looks at the inbound links to the "to" node
func lookAtIns(store *cayley.Handle, to string) {
	fmt.Printf("\nlookAtIns: object <-predicate- subject (%s)\n", to)
	fmt.Printf("=============================================\n")

	cayley.StartPath(store, quad.Raw(to)).Tag("object").InWithTags([]string{"predicate"}).Tag("subject").Iterate(nil).TagValues(nil, func(m map[string]quad.Value) {
		fmt.Printf("%s <-`%s` %s\n", m["object"], m["predicate"], m["subject"])
	})

}
