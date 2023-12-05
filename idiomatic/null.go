package idiomatic

import "github.com/boundedinfinity/go-jsonschema/model"

func NewNull() *JsonSchemaNull {
	schema := &JsonSchemaNull{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: model.SCHEMA_VERSION_2020_12,
			Type:   "null",
		},
	}

	return schema
}

type JsonSchemaNull struct {
	JsonSchemaCommon
}

var _ JsonSchema = &JsonSchemaNull{}

func (t JsonSchemaNull) TypeName() string {
	return "null"
}

func (t *JsonSchemaNull) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaNull) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	return nil
}
