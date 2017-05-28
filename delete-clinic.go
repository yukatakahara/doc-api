package admin

import "github.com/cayleygraph/cayley/quad"

func (a *Admin) DeleteClinic(jwt string, id string) error {
	idIRI := quad.IRI(id)

	err := a.Store.RemoveNode(idIRI)
	if err != nil {
		return err
	}

	return nil
}
