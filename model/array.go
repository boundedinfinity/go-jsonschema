package model

type JsonSchemaArray struct {
	JsonSchemaCommon `yaml:",inline"`
	Items            JsonSchema `json:"items" yaml:"items"`
}
