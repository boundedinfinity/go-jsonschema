package model

import o "github.com/boundedinfinity/go-commoner/optioner"

type JsonSchema2 struct {
	Anchor               o.Option[string]       `json:"$anchor" yaml:"$anchor"`
	AdditionalProperties o.Option[bool]         `json:"additionalProperties" yaml:"additionalProperties"`
	Const                o.Option[string]       `json:"const" yaml:"const"`
	Defs                 map[string]*JsonSchema `json:"$defs" yaml:"$defs"`

	MaxContains o.Option[int] `json:"maxContains" yaml:"maxContains"`
	MaxItems    o.Option[int] `json:"maxItems" yaml:"maxItems"`
	MinContains o.Option[int] `json:"minContains" yaml:"minContains"`
	MinItems    o.Option[int] `json:"minItems" yaml:"minItems"`

	PropertyNames     map[string]string      `json:"propertyNames" yaml:"pr  opertyNames"`
	PatternProperties map[string]*JsonSchema `json:"patternProperties" yaml:"patternProperties"`
	Ref               o.Option[string]       `json:"$ref" yaml:"$ref"`
	Required          []string               `json:"required" yaml:"required"`
	UniqueItems       o.Option[bool]         `json:"uniqueItems" yaml:"uniqueItems"`
}
