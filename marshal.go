package jsonschema

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func (t *System) unmarshal(s *JsonSchmea, ft JsonSchemaFileType, bs []byte) error {
	switch ft {
	case JsonSchemaFileType_Json:
		if err := json.Unmarshal(bs, s); err != nil {
			return err
		}
	case JsonSchemaFileType_Yaml:
		if err := yaml.Unmarshal(bs, s); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported file type")
	}

	return nil
}
