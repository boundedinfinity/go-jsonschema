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
	Anchor               optioner.Optioner[string]       `json:"$anchor" yaml:"$anchor"`
	AdditionalProperties optioner.Optioner[bool]         `json:"additionalProperties" yaml:"additionalProperties"`
	Comment              optioner.Optioner[string]       `json:"$comment" yaml:"$comment"`
	Const                optioner.Optioner[string]       `json:"const" yaml:"const"`
	Default              interface{}                     `json:"default" yaml:"default"`
	Defs                 map[string]*JsonSchmea          `json:"$defs" yaml:"$defs"`
	Deprecated           optioner.Optioner[bool]         `json:"deprecated" yaml:"deprecated"`
	Description          optioner.Optioner[string]       `json:"description" yaml:"description"`
	Enum                 []string                        `json:"enum" yaml:"enum"`
	Examples             []interface{}                   `json:"examples" yaml:"examples"`
	ExclusiveMaximum     optioner.Optioner[float64]      `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum     optioner.Optioner[float64]      `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Format               stringformat.StringFormatOption `json:"format" yaml:"format"`
	Id                   optioner.Optioner[string]       `json:"$id" yaml:"$id"`
	Items                *JsonSchmea                     `json:"items" yaml:"items"`
	MaxContains          optioner.Optioner[int]          `json:"maxContains" yaml:"maxContains"`
	MaxItems             optioner.Optioner[int]          `json:"maxItems" yaml:"maxItems"`
	Maximum              optioner.Optioner[float64]      `json:"maximum" yaml:"maximum"`
	MaxLength            optioner.Optioner[int]          `json:"maxLength" yaml:"maxLength"`
	MaxProperties        optioner.Optioner[int]          `json:"maxProperties" yaml:"maxProperties"`
	MinContains          optioner.Optioner[int]          `json:"minContains" yaml:"minContains"`
	MinItems             optioner.Optioner[int]          `json:"minItems" yaml:"minItems"`
	Minimum              optioner.Optioner[float64]      `json:"minimum" yaml:"minimum"`
	MinLength            optioner.Optioner[int]          `json:"minLength" yaml:"minLength"`
	MinProperties        optioner.Optioner[int]          `json:"minProperties" yaml:"minProperties"`
	MultipleOf           optioner.Optioner[float64]      `json:"multipleOf" yaml:"multipleOf"`
	Pattern              optioner.Optioner[string]       `json:"pattern" yaml:"pattern"`
	Properties           map[string]*JsonSchmea          `json:"properties" yaml:"properties"`
	PropertyNames        map[string]string               `json:"propertyNames" yaml:"pr  opertyNames"`
	PatternProperties    map[string]*JsonSchmea          `json:"patternProperties" yaml:"patternProperties"`
	ReadOnly             optioner.Optioner[bool]         `json:"readOnly" yaml:"readOnly"`
	Ref                  optioner.Optioner[string]       `json:"$ref" yaml:"$ref"`
	Required             []string                        `json:"required" yaml:"required"`
	Schema               optioner.Optioner[string]       `json:"$schema" yaml:"$schema"`
	Title                optioner.Optioner[string]       `json:"title" yaml:"title"`
	Type                 objecttype.ObjectTypeOption     `json:"type" yaml:"type"`
	UniqueItems          optioner.Optioner[bool]         `json:"uniqueItems" yaml:"uniqueItems"`
	WriteOnly            optioner.Optioner[bool]         `json:"writeOnly" yaml:"writeOnly"`
}

func (t JsonSchmea) HasValidation() bool {
	if t.Type.IsDefined() {
		switch t.Type.Get() {
		case objecttype.Integer, objecttype.Number:
			if t.Minimum.Defined() || t.ExclusiveMinimum.Defined() ||
				t.Maximum.Defined() || t.ExclusiveMaximum.Defined() ||
				t.MultipleOf.Defined() {
				return true
			}
		case objecttype.String:
			if t.MinLength.Defined() || t.MaxLength.Defined() ||
				t.Pattern.Defined() || t.Format.IsDefined() {
				return true
			}
		default:
			return false
		}
	}

	return false
}
