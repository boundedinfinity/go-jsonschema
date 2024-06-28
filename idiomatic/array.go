package idiomatic

import "encoding/json"

var _ JsonSchema = &JsonSchemaArray{}

type JsonSchemaArray struct {
	JsonSchemaCore
	PrefixItems      []JsonSchema `json:"prefixItems,omitempty" yaml:"prefixItems,omitempty"`
	Items            JsonSchema   `json:"items,omitempty" yaml:"items,omitempty"`
	Contains         JsonSchema   `json:"contains,omitempty" yaml:"contains,omitempty"`
	MinItems         int          `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	MaxItems         int          `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MaxContains      int          `json:"maxContains,omitempty" yaml:"maxContains,omitempty"`
	MinContains      int          `json:"minContains,omitempty" yaml:"minContains,omitempty"`
	UniqueItems      bool         `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	UnevaluatedItems JsonSchema   `json:"unevaluatedItems,omitempty" yaml:"unevaluatedItems,omitempty"`
}

func (t JsonSchemaArray) Copy() JsonSchema {
	return JsonSchemaArray{
		JsonSchemaCore:   t.JsonSchemaCore.Copy(),
		PrefixItems:      copySchemaArray(t.PrefixItems),
		Items:            t.Items.Copy(),
		Contains:         t.Contains.Copy(),
		MinItems:         t.MinItems,
		MaxItems:         t.MaxItems,
		MaxContains:      t.MaxContains,
		MinContains:      t.MaxContains,
		UniqueItems:      t.UniqueItems,
		UnevaluatedItems: t.UnevaluatedItems.Copy(),
	}
}

func (t JsonSchemaArray) GetId() string {
	return t.Id
}

func (t JsonSchemaArray) GetType() string {
	return "array"
}

func (t JsonSchemaArray) Validate() error {
	if err := t.JsonSchemaCore.Validate(); err != nil {
		return err
	}

	return nil
}

func (t *JsonSchemaArray) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type            string `json:"type"`
		JsonSchemaArray `json:",inline"`
	}{
		Type:            t.GetType(),
		JsonSchemaArray: *t,
	}
	return json.Marshal(dto)
}
