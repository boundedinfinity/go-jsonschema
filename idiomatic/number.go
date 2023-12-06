package idiomatic

import "github.com/boundedinfinity/go-jsonschema/model"

// import (
// 	o "github.com/boundedinfinity/go-commoner/optioner"
// 	"github.com/boundedinfinity/go-jsonschema/schematype"
// )

func NewNumber() *JsonSchemaNumber {
	schema := &JsonSchemaNumber{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: model.SCHEMA_VERSION_2020_12,
			Type:   "number",
		},
	}

	return schema
}

type JsonSchemaNumber struct {
	JsonSchemaCommon
	ExclusiveMaximum float64 `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum float64 `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          float64 `json:"maximum" yaml:"maximum"`
	Minimum          float64 `json:"minimum" yaml:"minimum"`
	MultipleOf       float64 `json:"multipleOf" yaml:"multipleOf"`
}

var _ JsonSchema = &JsonSchemaNumber{}

func (t JsonSchemaNumber) TypeName() string {
	return "number"
}

func (t *JsonSchemaNumber) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaNumber) Copy() JsonSchema {
	return &JsonSchemaNumber{
		JsonSchemaCommon: t.Common().Copy(),
		ExclusiveMaximum: t.ExclusiveMaximum,
		ExclusiveMinimum: t.ExclusiveMinimum,
		Maximum:          t.Maximum,
		Minimum:          t.Minimum,
		MultipleOf:       t.MultipleOf,
	}
}

func (t JsonSchemaNumber) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	// if err := validateMultipleOf(t.MultipleOf); err != nil {
	// 	return err
	// }

	if err := validateMaxMin(t.Maximum, t.Minimum); err != nil {
		return err
	}

	return nil
}
