package idiomatic

import "github.com/boundedinfinity/go-jsonschema/model"

// import (
// 	o "github.com/boundedinfinity/go-commoner/optioner"
// 	"github.com/boundedinfinity/go-jsonschema/schematype"
// )

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
	ExclusiveMaximum float64 `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum float64 `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          float64 `json:"maximum" yaml:"maximum"`
	Minimum          float64 `json:"minimum" yaml:"minimum"`
	MultipleOf       float64 `json:"multipleOf" yaml:"multipleOf"`
}

var _ JsonSchema = &JsonSchemaInteger{}

func (t JsonSchemaInteger) TypeName() string {
	return "integer"
}

func (t *JsonSchemaInteger) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaInteger) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	if err := validateMultipleOf(t.MultipleOf); err != nil {
		return err
	}

	if err := validateMaxMin(t.Maximum, t.Minimum); err != nil {
		return err
	}

	return nil
}
