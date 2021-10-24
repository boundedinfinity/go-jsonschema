package jsonschema

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"path"
	"strings"
)

func (t *System) readContent(uri string, ut JsonSchemaUriType, bs *[]byte) error {
	switch ut {
	case JsonSchemaUriType_File:
		if err := t.readFile(uri, bs); err != nil {
			return err
		}
	case JsonSchemaUriType_Http:
		if err := t.readHttp(uri, bs); err != nil {
			return err
		}
	default:
		return ErrUnknownUriTypeV(uri)
	}

	return nil
}

func (t *System) readFile(uri string, bs *[]byte) error {
	p := uri
	p = strings.TrimPrefix(p, "file://")

	bs2, err := ioutil.ReadFile(p)

	if err != nil {
		return err
	}

	*bs = bs2

	return nil
}

func (t *System) readHttp(uri string, bs *[]byte) error {
	return fmt.Errorf("TODO")
}

func (t *System) detectFileType(uri string, ft *JsonSchemaFileType) error {
	*ft = JsonSchemaFileType_Unknown

	ext := path.Ext(uri)

	switch ext {
	case ".yaml":
		*ft = JsonSchemaFileType_Yaml
	case ".yml":
		*ft = JsonSchemaFileType_Yaml
	case ".json":
		*ft = JsonSchemaFileType_Json
	default:
		return ErrUnknownFileTypeV(uri)
	}

	return nil
}

func (t *System) detectUriType(uri string, ut *JsonSchemaUriType) error {
	*ut = JsonSchemaUriType_Unknown

	parsed, err := url.Parse(uri)

	if err != nil {
		return err
	}

	if parsed.Scheme == "http" || parsed.Scheme == "https" {
		*ut = JsonSchemaUriType_Http
	}

	if parsed.Scheme == "file" {
		*ut = JsonSchemaUriType_File
	}

	if *ut == JsonSchemaUriType_Unknown {
		return ErrUnknownUriTypeV(uri)
	}

	return nil
}
