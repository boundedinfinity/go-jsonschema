package jsonschema

import "fmt"

func New() *System {
	return &System{
		i2l:       make(map[string]string),
		l2i:       make(map[string]string),
		schmeaMap: make(map[string]*JsonSchmea),
	}
}

type System struct {
	i2l       map[string]string
	l2i       map[string]string
	schmeaMap map[string]*JsonSchmea
}

func (t *System) LoadSchema(uris ...string) error {
	files, err := t.uris2files(uris...)

	if err != nil {
		return err
	}

	for _, file := range files {
		var schema JsonSchmea
		var uriType JsonSchemaUriType
		var fileType JsonSchemaFileType
		var bs []byte

		if err := t.detectUriType(file, &uriType); err != nil {
			return err
		}

		if err := t.detectFileType(file, &fileType); err != nil {
			return err
		}

		if err := t.readContent(file, uriType, &bs); err != nil {
			return err
		}

		if err := t.unmarshal(&schema, fileType, bs); err != nil {
			return err
		}

		if err := t.mapSchema(&schema, file); err != nil {
			return err
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
