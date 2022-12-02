package jsonschema

import (
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func (t *System) LoadPath(paths ...string) error {
	for _, path := range paths {
		if err := t.loadPath(path); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) loadPath(root string) error {
	m, err := marshaler.ReadFromPath(root)

	if err != nil {
		return err
	}

	for path, content := range m {
		if t.pathMap.Has(path) {
			return model.ErrPathDuplicatev(path)
		}

		if err := t.LoadSchema(content.Data, content.MimeType, path); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) LoadSchema(data []byte, mt mime_type.MimeType, path string) error {
	var bs []byte
	var err error

	switch mt {
	case mime_type.ApplicationXYaml:
		bs, err = yaml.YAMLToJSON(data)

		if err != nil {
			return err
		}
	case mime_type.ApplicationJson:
		bs = data
	default:
		return model.ErrMimeTypeUnsupportedv(mt)
	}

	schema, err := model.UnmarshalSchema(bs)

	if err != nil {
		return err
	}

	t.pathMap[path] = schema

	return nil
}
