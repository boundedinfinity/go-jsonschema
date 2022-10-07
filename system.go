package jsonschema

import (
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-jsonschema/model"
)

func New() *System {
	sys := System{}
	sys.Clear()

	return &sys
}

type System struct {
	PathMap   mapper.Mapper[string, string]
	SchemaMap mapper.Mapper[string, model.JsonSchema]
	Schema    model.JsonSchema
}

func (t *System) Clear() {
	t.PathMap = make(mapper.Mapper[string, string])
	t.SchemaMap = make(mapper.Mapper[string, model.JsonSchema])
}
