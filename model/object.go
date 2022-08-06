package model

import o "github.com/boundedinfinity/optioner"

type JsonSchemaObject struct {
	Properties    o.Option[map[string]JsonSchema] `json:"properties" yaml:"properties"`
	MaxProperties o.Option[int]                   `json:"maxProperties" yaml:"maxProperties"`
	MinProperties o.Option[int]                   `json:"minProperties" yaml:"minProperties"`
}
