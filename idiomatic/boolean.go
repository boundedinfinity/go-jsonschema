package idiomatic

import "github.com/boundedinfinity/go-jsonschema/model"

func NewBoolean() *JsonSchemaBoolean {
	schema := &JsonSchemaBoolean{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: model.SCHEMA_VERSION_2020_12,
			Type:   "boolean",
		},
	}

	return schema
}

type JsonSchemaBoolean struct {
	JsonSchemaCommon
}

var _ JsonSchema = &JsonSchemaBoolean{}

func (t JsonSchemaBoolean) TypeName() string {
	return "boolean"
}

func (t *JsonSchemaBoolean) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaBoolean) Copy() JsonSchema {
	return &JsonSchemaBoolean{
		JsonSchemaCommon: t.Common().Copy(),
	}
}

func (t JsonSchemaBoolean) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	return nil
}
