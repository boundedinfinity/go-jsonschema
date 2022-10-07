package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/mimetyper/mime_type"
)

type typeSelector struct {
	Type o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
}

func Unmarshal(data []byte, schemas *[]JsonSchema, mt mime_type.MimeType) error {
	var selectors []typeSelector

	if err := marshaler.Unmarshal(&selectors, mt, data); err != nil {
		return err
	}

	for _, selector := range selectors {
		if selector.Type.Empty() {
			return ErrSchemaTypeMissing
		}

		var schema JsonSchema

		switch selector.Type.Get() {
		case schematype.String:
			var concrete []JsonSchemaString

			if err := marshaler.Unmarshal(&concrete, mt, data); err != nil {
				return err
			}

			schema = concrete[0]
		case schematype.Integer:
			var concrete []JsonSchemaInteger

			if err := marshaler.Unmarshal(&concrete, mt, data); err != nil {
				return err
			}

			schema = concrete[0]
		case schematype.Number:
			var concrete []JsonSchemaNumber

			if err := marshaler.Unmarshal(&concrete, mt, data); err != nil {
				return err
			}

			schema = concrete[0]
		default:
			return ErrSchemaTypeUnsupportedV(selector.Type.Get())
		}

		*schemas = append(*schemas, schema)
	}

	return nil
}
