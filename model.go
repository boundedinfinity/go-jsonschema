package jsonschema

import (
	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/jsonschema/stringformat"
	"github.com/boundedinfinity/optioner"
)

type JsonSchemaSelector struct {
	Id     optioner.Option[string]                `json:"$id" yaml:"$id"`
	Schema optioner.Option[string]                `json:"$schema" yaml:"$schema"`
	Type   optioner.Option[objecttype.ObjectType] `json:"type" yaml:"type"`
}

type JsonSchemaGeneric[T ~string | ~int64 | ~float64] struct {
	JsonSchemaSelector
	Comment     optioner.Option[string] `json:"$comment" yaml:"$comment"`
	Default     optioner.Option[T]      `json:"default" yaml:"default"`
	Deprecated  optioner.Option[bool]   `json:"deprecated" yaml:"deprecated"`
	Description optioner.Option[string] `json:"description" yaml:"description"`
	Examples    optioner.Option[[]T]    `json:"examples" yaml:"examples"`
	Title       optioner.Option[string] `json:"title" yaml:"title"`
	ReadOnly    optioner.Option[bool]   `json:"readOnly" yaml:"readOnly"`
	WriteOnly   optioner.Option[bool]   `json:"writeOnly" yaml:"writeOnly"`
}
type JsonSchema interface {
	Validate() error
}

type JsonSchemaNumeric[T ~int64 | ~float64] struct {
	JsonSchemaGeneric[T]
	ExclusiveMaximum optioner.Option[T] `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum optioner.Option[T] `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          optioner.Option[T] `json:"maximum" yaml:"maximum"`
	Minimum          optioner.Option[T] `json:"minimum" yaml:"minimum"`
	MultipleOf       optioner.Option[T] `json:"multipleOf" yaml:"multipleOf"`
}

type JsonSchemaString[T ~string] struct {
	JsonSchemaGeneric[T]
	Format    optioner.Option[stringformat.StringFormat] `json:"format" yaml:"format"`
	Enum      optioner.Option[[]T]                       `json:"enum" yaml:"enum"`
	MaxLength optioner.Option[T]                         `json:"maxLength" yaml:"maxLength"`
	MinLength optioner.Option[T]                         `json:"minLength" yaml:"minLength"`
	Pattern   optioner.Option[T]                         `json:"pattern" yaml:"pattern"`
}

type JsonSchemaStringExtended[T ~string] struct {
	JsonSchemaString[T]
	EnumDescription optioner.Option[map[T]string] `json:"enum-description" yaml:"enum-description"`
}

type JsonSchemaObject struct {
	Properties map[string]*JsonSchema `json:"properties" yaml:"properties"`
}

type JsonSchema2 struct {
	system               *System
	parent               optioner.Option[JsonSchema]
	Anchor               optioner.Option[string] `json:"$anchor" yaml:"$anchor"`
	AdditionalProperties optioner.Option[bool]   `json:"additionalProperties" yaml:"additionalProperties"`
	Const                optioner.Option[string] `json:"const" yaml:"const"`
	Defs                 map[string]*JsonSchema  `json:"$defs" yaml:"$defs"`

	Items         *JsonSchema          `json:"items" yaml:"items"`
	MaxContains   optioner.Option[int] `json:"maxContains" yaml:"maxContains"`
	MaxItems      optioner.Option[int] `json:"maxItems" yaml:"maxItems"`
	MaxProperties optioner.Option[int] `json:"maxProperties" yaml:"maxProperties"`
	MinContains   optioner.Option[int] `json:"minContains" yaml:"minContains"`
	MinItems      optioner.Option[int] `json:"minItems" yaml:"minItems"`
	MinProperties optioner.Option[int] `json:"minProperties" yaml:"minProperties"`

	PropertyNames     map[string]string       `json:"propertyNames" yaml:"pr  opertyNames"`
	PatternProperties map[string]*JsonSchema  `json:"patternProperties" yaml:"patternProperties"`
	Ref               optioner.Option[string] `json:"$ref" yaml:"$ref"`
	Required          []string                `json:"required" yaml:"required"`
	UniqueItems       optioner.Option[bool]   `json:"uniqueItems" yaml:"uniqueItems"`
}
