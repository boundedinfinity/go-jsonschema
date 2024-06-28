package json_schema

import "encoding/json"

//go:generate enumer -config=./string-format.enum.yaml

var _ JsonSchema = &JsonSchemaString{}

type JsonSchemaString struct {
	JsonSchemaCore
	MinLength        int          `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          string       `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxLength        int          `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Format           StringFormat `json:"format,omitempty" yaml:"format,omitempty"`
	ContentMediaType string       `json:"contentMediaType,omitempty" yaml:"contentMediaType,omitempty"`
	ContentSchema    string       `json:"contentSchema,omitempty" yaml:"contentSchema,omitempty"`
	ContentEncoding  string       `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`
}

func (t JsonSchemaString) GetId() string {
	return t.Id
}

func (t JsonSchemaString) GetType() string {
	return "string"
}

func (t JsonSchemaString) Copy() JsonSchema {
	return JsonSchemaString{
		JsonSchemaCore:   t.JsonSchemaCore.Copy(),
		MinLength:        t.MinLength,
		Pattern:          t.Pattern,
		MaxLength:        t.MaxLength,
		Format:           t.Format,
		ContentMediaType: t.ContentMediaType,
		ContentSchema:    t.ContentSchema,
		ContentEncoding:  t.ContentEncoding,
	}
}

func (t JsonSchemaString) Validate() error {
	if err := t.JsonSchemaCore.Validate(); err != nil {
		return err
	}

	return nil
}

func (t *JsonSchemaString) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type             string `json:"type"`
		JsonSchemaString `json:",inline"`
	}{
		Type:             t.GetType(),
		JsonSchemaString: *t,
	}
	return json.Marshal(dto)
}
