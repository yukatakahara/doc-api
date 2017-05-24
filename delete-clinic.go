package admin

import "github.com/cayleygraph/cayley/quad"

func (a *Admin) DeleteClinic(jwt string, id string) error {
	// id2 := quad.String(id)
	id2 := quad.StringToValue(id)

	err := store.RemoveNode(store.ValueOf(id2))
	if err != nil {
		return err
	}

	return nil
}
