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
	pathMap mapper.Mapper[string, model.JsonSchema]
	idMap   mapper.Mapper[string, model.JsonSchema]
	id2path mapper.Mapper[string, string]
	path2id mapper.Mapper[string, string]
}

func (t *System) Clear() {
	t.pathMap = make(mapper.Mapper[string, model.JsonSchema])
	t.idMap = make(mapper.Mapper[string, model.JsonSchema])
	t.id2path = make(mapper.Mapper[string, string])
	t.path2id = make(mapper.Mapper[string, string])
}
