package jsonschema

import (
	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/jsonschema/stringformat"
	"github.com/boundedinfinity/optioner"
)

//go:generate enumer -standalone=true -package=objecttype -name=ObjectType -items=string,number,integer,object,array,boolean,null
//go:generate optioner -package=objecttype -type=ObjectType
//go:generate enumer -standalone=true -package=stringformat -name=StringFormat -items-from=json-string-format.txt
//go:generate optioner -package=stringformat -type=StringFormat

type JsonSchmea struct {
	Anchor               optioner.StringOption           `json:"$anchor" yaml:"$anchor"`
	AdditionalProperties optioner.BoolOption             `json:"additionalProperties" yaml:"additionalProperties"`
	Comment              optioner.StringOption           `json:"$comment" yaml:"$comment"`
	Const                optioner.StringOption           `json:"const" yaml:"const"`
	Default              interface{}                     `json:"default" yaml:"default"`
	Defs                 map[string]*JsonSchmea          `json:"$defs" yaml:"$defs"`
	Deprecated           optioner.BoolOption             `json:"deprecated" yaml:"deprecated"`
	Description          optioner.StringOption           `json:"description" yaml:"description"`
	Enum                 []string                        `json:"enum" yaml:"enum"`
	Examples             []interface{}                   `json:"examples" yaml:"examples"`
	ExclusiveMaximum     optioner.Float64Option          `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum     optioner.Float64Option          `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Format               stringformat.StringFormatOption `json:"format" yaml:"format"`
	Id                   optioner.StringOption           `json:"$id" yaml:"$id"`
	Items                *JsonSchmea                     `json:"items" yaml:"items"`
	MaxContains          optioner.IntOption              `json:"maxContains" yaml:"maxContains"`
	MaxItems             optioner.IntOption              `json:"maxItems" yaml:"maxItems"`
	Maximum              optioner.Float64Option          `json:"maximum" yaml:"maximum"`
	MaxLength            optioner.IntOption              `json:"maxLength" yaml:"maxLength"`
	MaxProperties        optioner.IntOption              `json:"maxProperties" yaml:"maxProperties"`
	MinContains          optioner.IntOption              `json:"minContains" yaml:"minContains"`
	MinItems             optioner.IntOption              `json:"minItems" yaml:"minItems"`
	Minimum              optioner.Float64Option          `json:"minimum" yaml:"minimum"`
	MinLength            optioner.IntOption              `json:"minLength" yaml:"minLength"`
	MinProperties        optioner.IntOption              `json:"minProperties" yaml:"minProperties"`
	MultipleOf           optioner.Float64Option          `json:"multipleOf" yaml:"multipleOf"`
	Pattern              optioner.StringOption           `json:"pattern" yaml:"pattern"`
	Properties           map[string]*JsonSchmea          `json:"properties" yaml:"properties"`
	PropertyNames        map[string]string               `json:"propertyNames" yaml:"pr  opertyNames"`
	PatternProperties    map[string]*JsonSchmea          `json:"patternProperties" yaml:"patternProperties"`
	ReadOnly             optioner.BoolOption             `json:"readOnly" yaml:"readOnly"`
	Ref                  optioner.StringOption           `json:"$ref" yaml:"$ref"`
	Required             []string                        `json:"required" yaml:"required"`
	Schema               optioner.StringOption           `json:"$schema" yaml:"$schema"`
	Title                optioner.StringOption           `json:"title" yaml:"title"`
	Type                 objecttype.ObjectTypeOption     `json:"type" yaml:"type"`
	UniqueItems          optioner.BoolOption             `json:"uniqueItems" yaml:"uniqueItems"`
	WriteOnly            optioner.BoolOption             `json:"writeOnly" yaml:"writeOnly"`
}
