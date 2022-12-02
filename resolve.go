package jsonschema

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
)

func (t *System) Resolve(schema model.JsonSchema) o.Option[model.JsonSchema] {
	if schema.IsConcrete() {
		return o.Some(schema)
	}

	id := schema.GetId().Get()
	return t.idMap.Get(id)
}
