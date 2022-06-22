package jsonschema

import (
	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/jsonschema/stringformat"
	"github.com/boundedinfinity/optioner"
)

//go:generate enumer -path=./objecttype/main.go
//go:generate enumer -path=./stringformat/main.go

type JsonSchema struct {
	system               *System
	parent               optioner.Option[JsonSchema]
	Anchor               optioner.Option[string]                    `json:"$anchor" yaml:"$anchor"`
	AdditionalProperties optioner.Option[bool]                      `json:"additionalProperties" yaml:"additionalProperties"`
	Comment              optioner.Option[string]                    `json:"$comment" yaml:"$comment"`
	Const                optioner.Option[string]                    `json:"const" yaml:"const"`
	Default              interface{}                                `json:"default" yaml:"default"`
	Defs                 map[string]*JsonSchema                     `json:"$defs" yaml:"$defs"`
	Deprecated           optioner.Option[bool]                      `json:"deprecated" yaml:"deprecated"`
	Description          optioner.Option[string]                    `json:"description" yaml:"description"`
	Enum                 []string                                   `json:"enum" yaml:"enum"`
	Examples             []interface{}                              `json:"examples" yaml:"examples"`
	ExclusiveMaximum     optioner.Option[float64]                   `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum     optioner.Option[float64]                   `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Format               optioner.Option[stringformat.StringFormat] `json:"format" yaml:"format"`
	Id                   optioner.Option[string]                    `json:"$id" yaml:"$id"`
	Items                *JsonSchema                                `json:"items" yaml:"items"`
	MaxContains          optioner.Option[int]                       `json:"maxContains" yaml:"maxContains"`
	MaxItems             optioner.Option[int]                       `json:"maxItems" yaml:"maxItems"`
	Maximum              optioner.Option[float64]                   `json:"maximum" yaml:"maximum"`
	MaxLength            optioner.Option[int]                       `json:"maxLength" yaml:"maxLength"`
	MaxProperties        optioner.Option[int]                       `json:"maxProperties" yaml:"maxProperties"`
	MinContains          optioner.Option[int]                       `json:"minContains" yaml:"minContains"`
	MinItems             optioner.Option[int]                       `json:"minItems" yaml:"minItems"`
	Minimum              optioner.Option[float64]                   `json:"minimum" yaml:"minimum"`
	MinLength            optioner.Option[int]                       `json:"minLength" yaml:"minLength"`
	MinProperties        optioner.Option[int]                       `json:"minProperties" yaml:"minProperties"`
	MultipleOf           optioner.Option[float64]                   `json:"multipleOf" yaml:"multipleOf"`
	Pattern              optioner.Option[string]                    `json:"pattern" yaml:"pattern"`
	Properties           map[string]*JsonSchema                     `json:"properties" yaml:"properties"`
	PropertyNames        map[string]string                          `json:"propertyNames" yaml:"pr  opertyNames"`
	PatternProperties    map[string]*JsonSchema                     `json:"patternProperties" yaml:"patternProperties"`
	ReadOnly             optioner.Option[bool]                      `json:"readOnly" yaml:"readOnly"`
	Ref                  optioner.Option[string]                    `json:"$ref" yaml:"$ref"`
	Required             []string                                   `json:"required" yaml:"required"`
	Schema               optioner.Option[string]                    `json:"$schema" yaml:"$schema"`
	Title                optioner.Option[string]                    `json:"title" yaml:"title"`
	Type                 optioner.Option[objecttype.ObjectType]     `json:"type" yaml:"type"`
	UniqueItems          optioner.Option[bool]                      `json:"uniqueItems" yaml:"uniqueItems"`
	WriteOnly            optioner.Option[bool]                      `json:"writeOnly" yaml:"writeOnly"`
}
