package jsonschema

import (
	"github.com/boundedinfinity/mimetyper/file_extention"
	"github.com/boundedinfinity/optioner"
)

var (
	default_extentions = []string{
		mkExt(file_extention.Yaml),
		mkExt(file_extention.Yml),
		mkExt(file_extention.Json),
	}
)

func mkExt(ft file_extention.FileExtention) string {
	return ".schema" + ft.String()
}

func New() *System {
	sys := System{
		Schema:     JsonSchema{},
		Extentions: default_extentions,
	}

	sys.Clear()

	return &sys
}

type System struct {
	Extentions []string
	SourceMap  map[string]string
	IdMap      map[string]JsonSchema
	Schema     JsonSchema
}

func (t *System) ById(id string) optioner.Option[JsonSchema] {
	if s, ok := t.IdMap[id]; ok {
		return optioner.Some(s)
	} else {
		return optioner.None[JsonSchema]()
	}
}

func (t *System) ByFileKey(id string) optioner.Option[JsonSchema] {
	if id, ok := t.SourceMap[id]; ok {
		return t.ById(id)
	} else {
		return optioner.None[JsonSchema]()
	}
}

func (t *System) GetFileKey(id string) optioner.Option[string] {
	for id2, key := range t.SourceMap {
		if id == id2 {
			return optioner.Some(key)
		}
	}

	return optioner.None[string]()
}

func (t *System) Clear() {
	t.SourceMap = make(map[string]string)
	t.IdMap = make(map[string]JsonSchema)
}
