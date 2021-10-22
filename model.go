package jsonschema

import (
	"fmt"

	"github.com/boundedinfinity/optional"
)

//go:generate enumer -package=jsonschema -name=JsonSchemaFileType -items=unknown,json,yaml
//go:generate enumer -package=jsonschema -name=JsonSchemaUriType -items=unknown,file,http
//go:generate enumer -package=jsonschema -name=JsonSchemaType -items=string,number,integer,object,array,boolean,null
//go:generate enumer -package=jsonschema -name=JsonSchemaStringFormat -items=date-time,time,date,duration,email,idn-email,hostname,idn-hostname,ipv4,ipv6,uuid,uri,uri-reference,iri,iri-reference,uri-template,json-pointer,regex
//go:generate enumer -package=jsonschema -name=JsonSchemaMimeType -items=xxx

var (
	ErrUnknownUriType        = fmt.Errorf("unknown uri type")
	ErrUnknownFileType       = fmt.Errorf("unknown file type")
	ErrDuplicateDef          = fmt.Errorf("duplicate $def")
	ErrUnsupportedSchemaType = fmt.Errorf("unsupported schema type")
	ErrRefNotFound           = fmt.Errorf("reference not found")
)

func ErrUnknownUriTypeV(v string) error  { return errV(v, ErrUnknownUriType) }
func ErrUnknownFileTypeV(v string) error { return errV(v, ErrUnknownFileType) }
func ErrDuplicateDefV(v string) error    { return errV(v, ErrDuplicateDef) }
func ErrRefNotFoundV(v string) error     { return errV(v, ErrRefNotFound) }

func ErrUnsupportedSchemaTypeV(v JsonSchemaType) error {
	return errV(v.String(), ErrUnsupportedSchemaType)
}

func errV(v string, err error) error {
	return fmt.Errorf("%v : %w", v, err)
}

type JsonSchmea struct {
	Schema       optional.StringOptional     `json:"$schema" yaml:"$schema"`
	Id           optional.StringOptional     `json:"$id" yaml:"$id"`
	Comment      optional.StringOptional     `json:"$comment" yaml:"$comment"`
	Title        optional.StringOptional     `json:"title" yaml:"title"`
	Description  optional.StringOptional     `json:"description" yaml:"description"`
	Properties   map[string]JsonSchmeaObject `json:"properties" yaml:"properties"`
	Defs         map[string]JsonSchmeaObject `json:"$defs" yaml:"$defs"`
	BaseUri      string
	RetrievalUri string
}

type JsonSchmeaObject struct {
	Type                 JsonSchemaType              `json:"type" yaml:"type"`
	Properties           map[string]JsonSchmeaObject `json:"properties" yaml:"properties"`
	PatternProperties    map[string]JsonSchmeaObject `json:"patternProperties" yaml:"patternProperties"`
	AdditionalProperties optional.BoolOptional       `json:"additionalProperties" yaml:"additionalProperties"`
	Required             []string                    `json:"required" yaml:"required"`
	PropertyNames        map[string]string           `json:"propertyNames" yaml:"propertyNames"`
	MinProperties        optional.IntOptional        `json:"minProperties" yaml:"minProperties"`
	MaxProperties        optional.IntOptional        `json:"maxProperties" yaml:"maxProperties"`
	Title                optional.StringOptional     `json:"title" yaml:"title"`
	Description          optional.StringOptional     `json:"description" yaml:"description"`
	Items                *JsonSchmeaObject           `json:"items" yaml:"items"`
	MinContains          optional.IntOptional        `json:"minContains" yaml:"minContains"`
	MaxContains          optional.IntOptional        `json:"maxContains" yaml:"maxContains"`
	MinItems             optional.IntOptional        `json:"minItems" yaml:"minItems"`
	MaxItems             optional.IntOptional        `json:"maxItems" yaml:"maxItems"`
	UniqueItems          optional.BoolOptional       `json:"uniqueItems" yaml:"uniqueItems"`
	Default              interface{}                 `json:"default" yaml:"default"`
	Examples             []interface{}               `json:"examples" yaml:"examples"`
	Deprecated           optional.BoolOptional       `json:"deprecated" yaml:"deprecated"`
	ReadOnly             optional.BoolOptional       `json:"readOnly" yaml:"readOnly"`
	WriteOnly            optional.BoolOptional       `json:"writeOnly" yaml:"writeOnly"`
	Enum                 []string                    `json:"enum" yaml:"enum"`
	Const                optional.StringOptional     `json:"const" yaml:"const"`
	Ref                  optional.StringOptional     `json:"$ref" yaml:"$ref"`
	Anchor               optional.StringOptional     `json:"$anchor" yaml:"$anchor"`
}

type JsonSchmeaString struct {
	MinLength optional.IntOptional    `json:"minLength" yaml:"minLength"`
	MaxLength optional.IntOptional    `json:"maxLength" yaml:"maxLength"`
	Pattern   optional.StringOptional `json:"pattern" yaml:"pattern"`
}

type JsonSchmeaInteger struct {
	MultipleOf       optional.IntOptional  `json:"multipleOf" yaml:"multipleOf"`
	Minimum          optional.IntOptional  `json:"minimum" yaml:"minimum"`
	ExclusiveMinimum optional.BoolOptional `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          optional.IntOptional  `json:"maximum" yaml:"maximum"`
	ExclusiveMaximum optional.BoolOptional `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
}

type Number struct {
	JsonSchmeaMultipleOf optional.IntOptional  `json:"multipleOf" yaml:"multipleOf"`
	Minimum              optional.IntOptional  `json:"minimum" yaml:"minimum"`
	ExclusiveMinimum     optional.BoolOptional `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum              optional.IntOptional  `json:"maximum" yaml:"maximum"`
	ExclusiveMaximum     optional.BoolOptional `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
}

type JsonSchmeaBoolean struct {
}

type JsonSchmeaNull struct {
}
