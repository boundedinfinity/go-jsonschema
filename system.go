package jsonschema

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/mimetyper/file_extention"
)

var (
	default_extentions = []string{
		mkExt(file_extention.Yaml),
		mkExt(file_extention.Yml),
		mkExt(file_extention.Json),
	}
)

func New() *System {
	sys := System{
		Extentions: default_extentions,
	}

	return &sys
}

type System struct {
	Extentions []string
	SourceMap  map[string]string
	IdMap      map[string]model.JsonSchema
	Schema     model.JsonSchema
}

func (t *System) HasId(id string) bool {
	_, ok := t.IdMap[id]
	return ok
}

func (t *System) ById(id string) o.Option[model.JsonSchema] {
	if s, ok := t.IdMap[id]; ok {
		return o.Some(s)
	} else {
		return o.None[model.JsonSchema]()
	}
}

// func (t *System) ByFileKey(id string) optioner.Option[JsonSchema] {
// 	if id, ok := t.SourceMap[id]; ok {
// 		return t.ById(id)
// 	} else {
// 		return optioner.None[JsonSchema]()
// 	}
// }

// func (t *System) GetFileKey(id string) optioner.Option[string] {
// 	for id2, key := range t.SourceMap {
// 		if id == id2 {
// 			return optioner.Some(key)
// 		}
// 	}

// 	return optioner.None[string]()
// }

// func (t *System) Clear() {
// 	t.SourceMap = make(map[string]string)
// 	t.IdMap = make(map[string]JsonSchema)
// }
