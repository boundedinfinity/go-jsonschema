package idiomatic

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var _ JsonSchema = &JsonSchemaInteger{}

type JsonSchemaInteger struct {
	JsonSchemaCore
	ExclusiveMaximum int `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum int `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	Maximum          int `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Minimum          int `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	MultipleOf       int `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

func (t JsonSchemaInteger) GetId() string {
	return t.Id
}

func (t JsonSchemaInteger) GetType() string {
	return "integer"
}

func (t JsonSchemaInteger) Copy() JsonSchema {
	return &JsonSchemaInteger{
		JsonSchemaCore:   t.JsonSchemaCore.Copy(),
		ExclusiveMaximum: t.ExclusiveMaximum,
		ExclusiveMinimum: t.ExclusiveMinimum,
		Maximum:          t.Maximum,
		Minimum:          t.Minimum,
		MultipleOf:       t.MultipleOf,
	}
}

var (
	ErrJsonSchemaIntegerMultipleOfLessThan1 = errorer.New("multipleOf less than zero")
)

func (t JsonSchemaInteger) Validate() error {
	if err := t.JsonSchemaCore.Validate(); err != nil {
		return nil
	}

	if err := validateMultipleOf(t.MultipleOf, ErrJsonSchemaIntegerMultipleOfLessThan1); err != nil {
		return err
	}

	// if err := validateMaxMin(t.Maximum, t.Minimum); err != nil {
	// 	return err
	// }

	return nil
}

func (t *JsonSchemaInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type              string `json:"type"`
		JsonSchemaInteger `json:",inline"`
	}{
		Type:              t.GetType(),
		JsonSchemaInteger: *t,
	}
	return json.Marshal(dto)
}
