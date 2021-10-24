package jsonschema

import (
	"github.com/boundedinfinity/optional"
)

//go:generate enumer -package=jsonschema -name=JsonSchemaFileType -items=unknown,json,yaml
//go:generate enumer -package=jsonschema -name=JsonSchemaUriType -items=unknown,file,http
//go:generate enumer -package=jsonschema -name=JsonSchemaType -items=string,number,integer,object,array,boolean,null
//go:generate enumer -package=jsonschema -name=JsonSchemaStringFormat -items=date-time,time,date,duration,email,idn-email,hostname,idn-hostname,ipv4,ipv6,uuid,uri,uri-reference,iri,iri-reference,uri-template,json-pointer,regex
//go:generate enumer -package=jsonschema -name=JsonSchemaMimeType -items=xxx

type JsonSchmea struct {
	Type                 JsonSchemaType           `json:"type" yaml:"type"`
	Schema               optional.StringOptional  `json:"$schema" yaml:"$schema"`
	Id                   optional.StringOptional  `json:"$id" yaml:"$id"`
	Comment              optional.StringOptional  `json:"$comment" yaml:"$comment"`
	Defs                 map[string]JsonSchmea    `json:"$defs" yaml:"$defs"`
	Properties           map[string]JsonSchmea    `json:"properties" yaml:"properties"`
	PatternProperties    map[string]JsonSchmea    `json:"patternProperties" yaml:"patternProperties"`
	AdditionalProperties optional.BoolOptional    `json:"additionalProperties" yaml:"additionalProperties"`
	Required             []string                 `json:"required" yaml:"required"`
	PropertyNames        map[string]string        `json:"propertyNames" yaml:"propertyNames"`
	MinProperties        optional.IntOptional     `json:"minProperties" yaml:"minProperties"`
	MaxProperties        optional.IntOptional     `json:"maxProperties" yaml:"maxProperties"`
	Title                optional.StringOptional  `json:"title" yaml:"title"`
	Description          optional.StringOptional  `json:"description" yaml:"description"`
	Items                *JsonSchmea              `json:"items" yaml:"items"`
	MinContains          optional.IntOptional     `json:"minContains" yaml:"minContains"`
	MaxContains          optional.IntOptional     `json:"maxContains" yaml:"maxContains"`
	MinItems             optional.IntOptional     `json:"minItems" yaml:"minItems"`
	MaxItems             optional.IntOptional     `json:"maxItems" yaml:"maxItems"`
	UniqueItems          optional.BoolOptional    `json:"uniqueItems" yaml:"uniqueItems"`
	Default              interface{}              `json:"default" yaml:"default"`
	Examples             []interface{}            `json:"examples" yaml:"examples"`
	Deprecated           optional.BoolOptional    `json:"deprecated" yaml:"deprecated"`
	ReadOnly             optional.BoolOptional    `json:"readOnly" yaml:"readOnly"`
	WriteOnly            optional.BoolOptional    `json:"writeOnly" yaml:"writeOnly"`
	Enum                 []string                 `json:"enum" yaml:"enum"`
	Const                optional.StringOptional  `json:"const" yaml:"const"`
	Ref                  optional.StringOptional  `json:"$ref" yaml:"$ref"`
	Anchor               optional.StringOptional  `json:"$anchor" yaml:"$anchor"`
	MinLength            optional.IntOptional     `json:"minLength" yaml:"minLength"`
	MaxLength            optional.IntOptional     `json:"maxLength" yaml:"maxLength"`
	Pattern              optional.StringOptional  `json:"pattern" yaml:"pattern"`
	MultipleOf           optional.Float64Optional `json:"multipleOf" yaml:"multipleOf"`
	Minimum              optional.Float64Optional `json:"minimum" yaml:"minimum"`
	ExclusiveMinimum     optional.Float64Optional `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum              optional.Float64Optional `json:"maximum" yaml:"maximum"`
	ExclusiveMaximum     optional.Float64Optional `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
}
