package jsonschema

import (
	"fmt"

	"github.com/boundedinfinity/mimetyper/mime_type"
)

func New() *System {
	return &System{
		Map: make(map[string]*JsonSchmea),
	}
}

type System struct {
	Map map[string]*JsonSchmea
}

func (t *System) LoadSchema(bss ...[]byte) error {
	for _, bs := range bss {
		var schemas []JsonSchmea
		var mimeType mime_type.MimeType

		if err := t.Unmarshal(&schemas, mimeType, bs); err != nil {
			return err
		}

		for _, schema := range schemas {
			if err := t.Add(&schema); err != nil {
				return err
			}
		}
	}

	for name, schema := range t.Map {
		fmt.Println(name)
		if err := t.Resolve(schema); err != nil {
			return err
		}
	}

	return nil
}
