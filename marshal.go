package jsonschema

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/boundedinfinity/jsonschema/mimetype"
	"gopkg.in/yaml.v2"
)

func (t *System) Unmarshal(ss *[]JsonSchmea, ft mimetype.MimeType, bs []byte) error {
	switch ft {
	case mimetype.ApplicationJson:
		return t.unmarshalJson(ss, bs)
	case mimetype.ApplicationYaml:
		return t.unmarshalYaml(ss, bs)
	default:
		return ErrMimeTypeUnsupportedV(ft)
	}
}

func (t *System) unmarshalYaml(ss *[]JsonSchmea, bs []byte) error {
	d := yaml.NewDecoder(bytes.NewReader(bs))

	for {
		s := new(JsonSchmea)

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

func (t *System) unmarshalJson(ss *[]JsonSchmea, bs []byte) error {
	s := string(bs)
	s = strings.TrimSpace(s)
	f := s[0:1]

	switch f {
	case "{":
		var x JsonSchmea
		if err := json.Unmarshal(bs, &s); err != nil {
			return err
		}
		*ss = append(*ss, x)
	case "[":
		var xs []JsonSchmea

		if err := json.Unmarshal(bs, &xs); err != nil {
			return err
		}

		*ss = append(*ss, xs...)
	default:
		return fmt.Errorf("unsupported file type")
	}

	return nil
}
