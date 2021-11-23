package jsonschema

import (
	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/optional"
)

//go:generate enumer -standalone=true -package=objecttype -name=ObjectType -items=string,number,integer,object,array,boolean,null
//go:generate enumer -standalone=true -package=stringformat -name=StringFormat -items-from=json-string-format.txt

type JsonSchmea struct {
	Anchor               optional.StringOptional  `json:"$anchor" yaml:"$anchor"`
	AdditionalProperties optional.BoolOptional    `json:"additionalProperties" yaml:"additionalProperties"`
	Comment              optional.StringOptional  `json:"$comment" yaml:"$comment"`
	Const                optional.StringOptional  `json:"const" yaml:"const"`
	Default              interface{}              `json:"default" yaml:"default"`
	Defs                 map[string]*JsonSchmea   `json:"$defs" yaml:"$defs"`
	Deprecated           optional.BoolOptional    `json:"deprecated" yaml:"deprecated"`
	Description          optional.StringOptional  `json:"description" yaml:"description"`
	Enum                 []string                 `json:"enum" yaml:"enum"`
	Examples             []interface{}            `json:"examples" yaml:"examples"`
	ExclusiveMaximum     optional.Float64Optional `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum     optional.Float64Optional `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Id                   optional.StringOptional  `json:"$id" yaml:"$id"`
	Items                *JsonSchmea              `json:"items" yaml:"items"`
	MaxContains          optional.IntOptional     `json:"maxContains" yaml:"maxContains"`
	MaxItems             optional.IntOptional     `json:"maxItems" yaml:"maxItems"`
	Maximum              optional.Float64Optional `json:"maximum" yaml:"maximum"`
	MaxLength            optional.IntOptional     `json:"maxLength" yaml:"maxLength"`
	MaxProperties        optional.IntOptional     `json:"maxProperties" yaml:"maxProperties"`
	MinContains          optional.IntOptional     `json:"minContains" yaml:"minContains"`
	MinItems             optional.IntOptional     `json:"minItems" yaml:"minItems"`
	Minimum              optional.Float64Optional `json:"minimum" yaml:"minimum"`
	MinLength            optional.IntOptional     `json:"minLength" yaml:"minLength"`
	MinProperties        optional.IntOptional     `json:"minProperties" yaml:"minProperties"`
	MultipleOf           optional.Float64Optional `json:"multipleOf" yaml:"multipleOf"`
	Pattern              optional.StringOptional  `json:"pattern" yaml:"pattern"`
	Properties           map[string]*JsonSchmea   `json:"properties" yaml:"properties"`
	PropertyNames        map[string]string        `json:"propertyNames" yaml:"pr  opertyNames"`
	PatternProperties    map[string]*JsonSchmea   `json:"patternProperties" yaml:"patternProperties"`
	ReadOnly             optional.BoolOptional    `json:"readOnly" yaml:"readOnly"`
	Ref                  optional.StringOptional  `json:"$ref" yaml:"$ref"`
	Required             []string                 `json:"required" yaml:"required"`
	Schema               optional.StringOptional  `json:"$schema" yaml:"$schema"`
	Title                optional.StringOptional  `json:"title" yaml:"title"`
	Type                 objecttype.ObjectType    `json:"type" yaml:"type"`
	UniqueItems          optional.BoolOptional    `json:"uniqueItems" yaml:"uniqueItems"`
	WriteOnly            optional.BoolOptional    `json:"writeOnly" yaml:"writeOnly"`
}
