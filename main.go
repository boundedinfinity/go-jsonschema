package jsonschema

func New() *System {
	return &System{
		i2l:    make(map[string]string),
		l2i:    make(map[string]string),
		uriMap: make(map[string]JsonSchmeaObject),
	}
}

type System struct {
	i2l    map[string]string
	l2i    map[string]string
	uriMap map[string]JsonSchmeaObject
}

func (t *System) LoadSchema(uris ...string) error {
	for _, uri := range uris {
		var schema JsonSchmea
		var uriType JsonSchemaUriType
		var fileType JsonSchemaFileType
		var bs []byte

		if err := t.detectUriType(uri, &uriType); err != nil {
			return err
		}

		if err := t.detectFileType(uri, &fileType); err != nil {
			return err
		}

		if err := t.readContent(uri, uriType, &bs); err != nil {
			return err
		}

		if err := t.unmarshal(&schema, fileType, bs); err != nil {
			return err
		}

		if schema.Id.IsDefined() {
			t.i2l[schema.Id.Get()] = uri
			t.l2i[uri] = schema.Id.Get()
		}

		if err := t.mapSchema(schema); err != nil {
			return err
		}

		if err := t.resolveSchema(&schema); err != nil {
			return err
		}
	}

	return nil
}
