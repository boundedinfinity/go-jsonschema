package idiomatic

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func NewMarshaler() *JsonSchemaMarshaler {
	jsm := &JsonSchemaMarshaler{
		km: marshaler.NewKind(
			jsonSchemaDescriminator{},
			func(d jsonSchemaDescriminator) string {
				return d.Type
			},
		),
		source:   make(mapper.Mapper[string, JsonSchema]),
		resolved: make(mapper.Mapper[string, JsonSchema]),
	}

	jsm.km.RegisterNamer(&JsonSchemaArray{})
	jsm.km.RegisterNamer(&JsonSchemaBoolean{})
	jsm.km.RegisterNamer(&JsonSchemaInteger{})
	jsm.km.RegisterNamer(&JsonSchemaNull{})
	jsm.km.RegisterNamer(&JsonSchemaNumber{})
	jsm.km.RegisterNamer(&JsonSchemaObject{})
	jsm.km.RegisterNamer(&JsonSchemaRef{})
	jsm.km.RegisterNamer(&JsonSchemaString{})

	return jsm
}

type jsonSchemaDescriminator struct {
	Type string `json:"type" yaml:"type"`
	Ref  string `json:"$ref" yaml:"$ref"`
}

type JsonSchemaMarshaler struct {
	km       *marshaler.KindMarshaler[jsonSchemaDescriminator]
	source   mapper.Mapper[string, JsonSchema]
	resolved mapper.Mapper[string, JsonSchema]
}

func (t JsonSchemaMarshaler) Unmarshal(mimeType mime_type.MimeType, data []byte) (JsonSchema, error) {
	resolvedMimeType := mime_type.ResolveMimeType(mimeType)

	switch resolvedMimeType {
	case mime_type.ApplicationXYaml:
		return t.UnmarshalFromYAML(data)
	case mime_type.ApplicationJson:
		return t.UnmarshalFromJSON(data)
	default:
		return nil, model.ErrJsonSchemaMimeTypeUnsupportedv(mimeType)
	}
}

func (t JsonSchemaMarshaler) UnmarshalFromYAML(data []byte) (JsonSchema, error) {
	dataJson, err := yaml.YAMLToJSON(data)

	if err != nil {
		return nil, err
	}

	return t.UnmarshalFromJSON(dataJson)
}

func (t JsonSchemaMarshaler) UnmarshalFromJSON(data []byte) (JsonSchema, error) {
	var val any
	var err error

	val, err = t.km.Unmarshal(data)

	if err != nil && !errors.Is(err, marshaler.ErrKindMarshalerTypeNotRegistered) {
		return nil, err
	}

	schema, ok := val.(JsonSchema)

	if !ok {
		return nil, model.ErrJsonSchemaTypeInvalidv(val)
	}

	if err := t.process(schema); err != nil {
		return schema, err
	}

	return nil, nil
}

func (t JsonSchemaMarshaler) process(schema JsonSchema) error {
	switch val := schema.(type) {
	case *JsonSchemaRef:
	default:
		return model.ErrJsonSchemaTypeInvalidv(val)
	}

	return nil
}

// func UnmarshalSchema(data []byte) (JsonSchema, error) {
// 	var out JsonSchema
// 	var d jsonSchemaDescriminator

// 	if d.Type.Defined() {
// 		switch d.Type.Get() {
// 		case schematype.String:
// 			var v JsonSchemaString

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Integer:
// 			var v JsonSchemaInteger

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Number:
// 			var v JsonSchemaNumber

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Array:
// 			var v JsonSchemaArray

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Object:
// 			var v JsonSchemaObject

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			for name, raw := range d.Properties {
// 				if property, err := UnmarshalSchema(raw); err != nil {
// 					return nil, err
// 				} else {
// 					v.Properties[name] = property
// 				}
// 			}

// 			out = &v
// 		default:
// 			if d.Ref.Defined() {
// 				return nil, fmt.Errorf("something went wrong")
// 			}
// 		}
// 	}

// 	if d.Ref.Defined() {
// 		var v JsonSchemaRef

// 		if err := json.Unmarshal(data, &v); err != nil {
// 			return nil, err
// 		}

// 		out = &v
// 	}

// 	if out == nil {
// 		return nil, fmt.Errorf("something went wrong")
// 	}

// 	return out, nil
// }
