package jsonschema

import (
	"fmt"

	"github.com/boundedinfinity/jsonschema/mimetype"
)

func New() *System {
	return &System{
		schmeaMap: make(map[string]*JsonSchmea),
	}
}

type System struct {
	schmeaMap map[string]*JsonSchmea
}

func (t *System) LoadSchema(bss ...[]byte) error {

	for _, bs := range bss {
		var schemas []JsonSchmea
		var mimeType mimetype.MimeType

		if err := t.Unmarshal(&schemas, mimeType, bs); err != nil {
			return err
		}

		for _, schema := range schemas {
			if err := t.AddtoMap(&schema); err != nil {
				return err
			}
		}
	}

	for name, schema := range t.schmeaMap {
		fmt.Println(name)
		if err := t.resolveSchema(schema); err != nil {
			return err
		}
	}

	return nil
}
