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
	schemaMap mapper.Mapper[string, model.JsonSchema]
}

func (t *System) Clear() {
	t.schemaMap = make(mapper.Mapper[string, model.JsonSchema])
}
