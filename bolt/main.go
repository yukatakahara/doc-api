package bolt

import (
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
)

// init and open
func Open(dbFile string) (*cayley.Handle, error) {
	graph.InitQuadStore("bolt", dbFile, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", dbFile, nil)

	if err != nil {
		log.Fatalln(err)
	}

	return store, nil
}

// AdminService represents a PostgreSQL implementation of myapp.UserService.
type AdminService struct {
	Store *cayley.Handle
}
