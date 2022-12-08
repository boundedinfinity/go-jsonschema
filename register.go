package jsonschema

import "github.com/boundedinfinity/go-jsonschema/model"

func (t *System) Register(schema model.JsonSchema) error {
	if schema.GetId().Defined() {
		t.pathMap[schema.GetId().Get()] = schema
	}

	return nil
}
