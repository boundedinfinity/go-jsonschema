package json_schema

import "encoding/json"

var _ JsonSchema = &JsonSchemaBoolean{}

type JsonSchemaBoolean struct {
	JsonSchemaCore
}

func (t JsonSchemaBoolean) GetId() string {
	return t.Id
}

func (t JsonSchemaBoolean) GetType() string {
	return "boolean"
}

func (t JsonSchemaBoolean) Copy() JsonSchema {
	return &JsonSchemaBoolean{
		JsonSchemaCore: t.JsonSchemaCore.Copy(),
	}
}

func (t JsonSchemaBoolean) Validate() error {
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

func (t *JsonSchemaBoolean) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type              string `json:"type"`
		JsonSchemaBoolean `json:",inline"`
	}{
		Type:              t.GetType(),
		JsonSchemaBoolean: *t,
	}
	return json.Marshal(dto)
}
