package idiomatic

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-jsonschema/model"
)

func NewObject() *JsonSchemaObject {
	schema := &JsonSchemaObject{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: model.SCHEMA_VERSION_2020_12,
			Type:   "object",
		},
		Properties: mapper.Mapper[string, JsonSchema]{},
	}

	return schema
}

type JsonSchemaObject struct {
	JsonSchemaCommon
	Properties    mapper.Mapper[string, JsonSchema] `json:"properties" yaml:"properties"`
	MaxProperties int64                             `json:"maxProperties" yaml:"maxProperties"`
	MinProperties int64                             `json:"minProperties" yaml:"minProperties"`
}

var _ JsonSchema = &JsonSchemaObject{}

func (t JsonSchemaObject) TypeName() string {
	return "boolean"
}

func (t *JsonSchemaObject) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaObject) Validate() error {
	if err := validateMaxMinProperties(t.MaxProperties, t.MinProperties); err != nil {
		return err
	}

	if t.Properties == nil {
		return model.ErrObjectPropertiesMissing
	}

	for _, property := range t.Properties {
		if err := property.Validate(); err != nil {
			return err
		}
	}

	return nil
}
