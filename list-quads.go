package admin

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
)

// get all quads from the db
func (a *Admin) AllQuads() ([]quad.Quad, error) {
	results, err := readAllQuads(store)

	if err != nil {
		return results, err
	}

	return results, nil
}

func readAllQuads(store *cayley.Handle) ([]quad.Quad, error) {
	var results []quad.Quad
	it := store.QuadsAllIterator()
	defer it.Close()

	for it.Next() {
		results = append(results, store.Quad(it.Result()))
	}

	return results, nil
}
