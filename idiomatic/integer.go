package idiomatic

import "github.com/boundedinfinity/go-jsonschema/model"

func NewInteger() *JsonSchemaInteger {
	schema := &JsonSchemaInteger{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: model.SCHEMA_VERSION_2020_12,
			Type:   "integer",
		},
	}

	return schema
}

type JsonSchemaInteger struct {
	JsonSchemaCommon
	ExclusiveMaximum int `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum int `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          int `json:"maximum" yaml:"maximum"`
	Minimum          int `json:"minimum" yaml:"minimum"`
	MultipleOf       int `json:"multipleOf" yaml:"multipleOf"`
}

var _ JsonSchema = &JsonSchemaInteger{}

func (t JsonSchemaInteger) TypeName() string {
	return "integer"
}

func (t JsonSchemaInteger) Copy() JsonSchema {
	return &JsonSchemaInteger{
		JsonSchemaCommon: t.Common().Copy(),
		ExclusiveMaximum: t.ExclusiveMaximum,
		ExclusiveMinimum: t.ExclusiveMinimum,
		Maximum:          t.Maximum,
		Minimum:          t.Minimum,
		MultipleOf:       t.MultipleOf,
	}
}

func (t *JsonSchemaInteger) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaInteger) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	if err := validateMultipleOf[int](t.MultipleOf); err != nil {
		return err
	}

	if err := validateMaxMin(t.Maximum, t.Minimum); err != nil {
		return err
	}

	return nil
}
