package model

import (
	"github.com/boundedinfinity/go-commoner/mapper"
	o "github.com/boundedinfinity/go-commoner/optioner"
)

type JsonSchemaObject struct {
	Properties    o.Option[mapper.Mapper[string, JsonSchema]] `json:"properties" yaml:"properties"`
	MaxProperties o.Option[int]                               `json:"maxProperties" yaml:"maxProperties"`
	MinProperties o.Option[int]                               `json:"minProperties" yaml:"minProperties"`
}
