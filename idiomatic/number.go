package idiomatic

import "encoding/json"

var _ JsonSchema = &JsonSchemaNumber{}

type JsonSchemaNumber struct {
	JsonSchemaCore
	ExclusiveMaximum float64 `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum float64 `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	Maximum          float64 `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Minimum          float64 `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	MultipleOf       float64 `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

func (t JsonSchemaNumber) GetId() string {
	return t.Id
}

func (t JsonSchemaNumber) GetType() string {
	return "number"
}

func (t JsonSchemaNumber) Copy() JsonSchema {
	return &JsonSchemaNumber{
		JsonSchemaCore:   t.JsonSchemaCore.Copy(),
		ExclusiveMaximum: t.ExclusiveMaximum,
		ExclusiveMinimum: t.ExclusiveMinimum,
		Maximum:          t.Maximum,
		Minimum:          t.Minimum,
		MultipleOf:       t.MultipleOf,
	}
}

func (t JsonSchemaNumber) Validate() error {
	if err := t.JsonSchemaCore.Validate(); err != nil {
		return nil
	}

	// if err := validateMultipleOf(t.MultipleOf); err != nil {
	// 	return err
	// }

	// if err := validateMaxMin(t.Maximum, t.Minimum); err != nil {
	// 	return err
	// }

	return nil
}

func (t *JsonSchemaNumber) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type             string `json:"type"`
		JsonSchemaNumber `json:",inline"`
	}{
		Type:             t.GetType(),
		JsonSchemaNumber: *t,
	}
	return json.Marshal(dto)
}
