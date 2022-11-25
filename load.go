package jsonschema

import (
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func (t *System) Load(paths ...string) error {
	for _, path := range paths {
		if err := t.loadPath(path); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) loadPath(path string) error {
	if t.pathMap.Has(path) {
		return model.ErrPathDuplicatev(path)
	}

	m, err := marshaler.ReadFromPath(path)

	if err != nil {
		return err
	}

	for _, content := range m {
		var bs []byte
		var err error

		switch content.MimeType {
		case mime_type.ApplicationXYaml:
			bs, err = yaml.YAMLToJSON(content.Data)

			if err != nil {
				return err
			}
		case mime_type.ApplicationJson:
			bs = content.Data
		default:
			return model.ErrMimeTypeUnsupportedv(content.MimeType)
		}

		schema, err := model.UnmarshalSchema(bs)

		if err != nil {
			return err
		}

		t.pathMap[path] = schema
	}

	return nil
}
