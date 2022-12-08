package jsonschema

import "github.com/boundedinfinity/go-jsonschema/model"

func (t *System) Register(schema model.JsonSchema) error {
	if schema.GetId().Defined() {
		if t.pathMap.Has(schema.GetId().Get()) {
			return model.ErrPathDuplicatev(schema.GetId().Get())
		}

		t.pathMap[schema.GetId().Get()] = schema
	}

	return nil
}
