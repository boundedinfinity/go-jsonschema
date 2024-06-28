package json_schema

var _ JsonSchema = &JsonSchemaObject{}

type JsonSchemaObject struct {
	JsonSchemaCore
	Properties            map[string]JsonSchema `json:"properties,omitempty" yaml:"properties,omitempty"`
	PatternProperties     map[string]JsonSchema `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
	AdditionalProperties  map[string]JsonSchema `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	PropertyNames         JsonSchema            `json:"propertyNames,omitempty" yaml:"propertyNames,omitempty"`
	DependentSchemas      map[string]JsonSchema `json:"dependentSchemas,omitempty" yaml:"dependentSchemas,omitempty"`
	MaxProperties         int                   `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties         int                   `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	DependentRequired     map[string][]string   `json:"dependentRequired,omitempty" yaml:"dependentRequired,omitempty"`
	Required              []string              `json:"required,omitempty" yaml:"required,omitempty"`
	UnevaluatedProperties JsonSchema            `json:"unevaluatedProperties,omitempty" yaml:"unevaluatedProperties,omitempty"`
}

func (t JsonSchemaObject) Copy() JsonSchema {
	dependentRequireds := map[string][]string{}
	for k, v := range t.DependentRequired {
		dependentRequireds[k] = v[:]
	}

	return JsonSchemaObject{
		JsonSchemaCore:        t.JsonSchemaCore.Copy(),
		Properties:            copySchemaMap(t.Properties),
		PatternProperties:     copySchemaMap(t.PatternProperties),
		AdditionalProperties:  copySchemaMap(t.AdditionalProperties),
		PropertyNames:         t.PropertyNames.Copy(),
		DependentSchemas:      copySchemaMap(t.DependentSchemas),
		MaxProperties:         t.MaxProperties,
		MinProperties:         t.MinProperties,
		DependentRequired:     dependentRequireds,
		Required:              t.Required[:],
		UnevaluatedProperties: t.UnevaluatedProperties.Copy(),
	}
}

func (t JsonSchemaObject) GetId() string {
	return t.Id
}

func (t JsonSchemaObject) GetType() string {
	return "object"
}

func (t JsonSchemaObject) Validate() error {
	if err := t.JsonSchemaCore.Validate(); err != nil {
		return err
	}

	return nil
}
