package jsonschema

import (
	"github.com/boundedinfinity/go-jsonschema/model"
)

func (t *System) Register(schema model.JsonSchema) error {
	if schema.GetId().Empty() {
		return model.ErrSchemaIdNotFound
	}

	if t.schemaMap.Has(schema.GetId().Get()) {
		return model.ErrSchemaIdDuplicatev(schema.GetId().Get())
	}

	t.schemaMap[schema.GetId().Get()] = schema

	return nil
}
