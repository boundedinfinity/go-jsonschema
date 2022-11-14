package jsonschema

import (
	"net/url"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-marshaler"
)

func (t *System) Load(urls ...string) error {
	var realUrls []*url.URL

	for _, rawUrl := range urls {
		realUrl, err := url.Parse(rawUrl)

		if err != nil {
			return err
		}

		realUrls = append(realUrls, realUrl)
	}

	for _, realUrl := range realUrls {
		if realUrl.Scheme == "file" {
			err := t.loadPath(realUrl.Path)

			if err != nil {
				return err
			}
		}
	}

	return t.Resolve()
}

func (t *System) loadPath(path string) error {
	ok, err := pather.IsDirErr(path)

	if err != nil {
		return err
	}

	if ok {
		return t.loadDir(path)
	}

	return t.loadFile(path)
}

func (t *System) loadDir(path string) error {
	files, err := pather.GetFiles(path)

	if err != nil {
		return err
	}

	for _, file := range files {
		err = t.loadFile(file)

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *System) loadFile(path string) error {
	d, _, err := marshaler.UnmarshalFromFile[model.JsonSchemaCommon](path)

	if err != nil {
		return err
	}

	if d.Type.Empty() && d.Ref.Empty() {
		return model.ErrSchemaTypeOrRefNotFound
	}

	var schema model.JsonSchema

	if d.Type.Defined() {
		switch d.Type.Get() {
		case schematype.String:
			v, _, err := marshaler.UnmarshalFromFile[model.JsonSchemaString](path)

			if err != nil {
				return err
			}

			schema = v
		case schematype.Integer:
			v, _, err := marshaler.UnmarshalFromFile[model.JsonSchemaInteger](path)

			if err != nil {
				return err
			}

			schema = v
		case schematype.Number:
			v, _, err := marshaler.UnmarshalFromFile[model.JsonSchemaNumber](path)

			if err != nil {
				return err
			}

			schema = v
		default:
			return model.ErrSchemaTypeUnsupportedv(d.Type.Get())
		}
	}

	if d.Ref.Defined() {
		schema = d
	}

	return t.Register(schema)
}
