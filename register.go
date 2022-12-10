package jsonschema

import (
	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
)

func (t *System) Register(schema model.JsonSchema, path optioner.Option[string]) error {
	if schema.GetId().Defined() {
		if t.idMap.Has(schema.GetId().Get()) {
			return model.ErrPathDuplicatev(schema.GetId().Get())
		}

		t.idMap[schema.GetId().Get()] = schema
	}

	if path.Defined() {
		t.pathMap[path.Get()] = schema
	}

	if schema.GetId().Defined() && path.Defined() {
		t.id2path[schema.GetId().Get()] = path.Get()
		t.path2id[path.Get()] = schema.GetId().Get()
	}

	return nil
}
