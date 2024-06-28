package json_schema

import "encoding/json"

var _ JsonSchema = &JsonSchemaNull{}

type JsonSchemaNull struct {
	JsonSchemaCore
}

func (t JsonSchemaNull) GetId() string {
	return t.Id
}

func (t JsonSchemaNull) GetType() string {
	return "null"
}

func (t JsonSchemaNull) Copy() JsonSchema {
	return &JsonSchemaNull{
		JsonSchemaCore: t.JsonSchemaCore.Copy(),
	}
}

func (t JsonSchemaNull) Validate() error {
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

func (t *JsonSchemaNull) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type           string `json:"type"`
		JsonSchemaNull `json:",inline"`
	}{
		Type:           t.GetType(),
		JsonSchemaNull: *t,
	}
	return json.Marshal(dto)
}
