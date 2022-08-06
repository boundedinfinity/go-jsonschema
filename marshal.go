package jsonschema

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/boundedinfinity/jsonschema/model"
	"github.com/boundedinfinity/mimetyper/mime_type"
	"gopkg.in/yaml.v2"
)

func (t *System) Unmarshal(ss *[]model.JsonSchema, mt mime_type.MimeType, bs []byte) error {
	realmt := mime_type.ResolveMimeType(mt)

	switch realmt {
	case mime_type.ApplicationJson:
		return t.unmarshalJson(ss, bs)
	case mime_type.ApplicationXYaml:
		return t.unmarshalYaml(ss, bs)
	default:
		return ErrMimeTypeUnsupportedV(mt)
	}
}

func (t *System) unmarshalYaml(ss *[]model.JsonSchema, bs []byte) error {
	d := yaml.NewDecoder(bytes.NewReader(bs))

	for {
		s := new(model.JsonSchema)
		err := d.Decode(&s)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		if s == nil {
			continue
		}

		*ss = append(*ss, *s)
	}

	return nil
}

func (t *System) unmarshalJson(ss *[]model.JsonSchema, bs []byte) error {
	s := string(bs)
	s = strings.TrimSpace(s)
	f := s[0:1]

	switch f {
	case "{":
		var x model.JsonSchema

		if err := json.Unmarshal(bs, &s); err != nil {
			return err
		}

		*ss = append(*ss, x)
	case "[":
		var xs []model.JsonSchema

		if err := json.Unmarshal(bs, &xs); err != nil {
			return err
		}

		*ss = append(*ss, xs...)
	default:
		return fmt.Errorf("unsupported JSON")
	}

	return nil
}
