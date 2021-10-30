package jsonschema

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/boundedinfinity/jsonschema/filetype"
	"gopkg.in/yaml.v2"
)

func (t *System) unmarshal(ss *[]JsonSchmea, ft filetype.FileType, bs []byte) error {
	switch ft {
	case filetype.Json:
		var s JsonSchmea
		if err := json.Unmarshal(bs, &s); err != nil {
			return err
		}
		*ss = append(*ss, s)
	case filetype.Yaml:
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

	default:
		return fmt.Errorf("unsupported file type")
	}

	return nil
}
