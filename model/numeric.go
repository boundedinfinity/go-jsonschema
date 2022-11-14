package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

type jsonSchemaNumeric[T int64 | float64] struct {
	JsonSchemaCommon `yaml:",inline"`
	ExclusiveMaximum o.Option[T] `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum o.Option[T] `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          o.Option[T] `json:"maximum" yaml:"maximum"`
	Minimum          o.Option[T] `json:"minimum" yaml:"minimum"`
	MultipleOf       o.Option[T] `json:"multipleOf" yaml:"multipleOf"`
}

func (t jsonSchemaNumeric[T]) Copy() JsonSchema {
	return jsonSchemaNumeric[T]{
		JsonSchemaCommon: t.JsonSchemaCommon.Copy().(JsonSchemaCommon),
		ExclusiveMaximum: t.ExclusiveMaximum,
		ExclusiveMinimum: t.ExclusiveMinimum,
		Maximum:          t.Maximum,
		Minimum:          t.Minimum,
		MultipleOf:       t.MultipleOf,
	}
}

func newNumeric[T int64 | float64](id string, t schematype.SchemaType) jsonSchemaNumeric[T] {
	return jsonSchemaNumeric[T]{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: o.Some(SCHEMA_VERSION_2020_12),
			Type:   o.Some(t),
			Id:     o.Some(id),
		},
	}
}

func (t jsonSchemaNumeric[T]) Merge(other JsonSchema) JsonSchema {
	internal := other.(jsonSchemaNumeric[T])

	return jsonSchemaNumeric[T]{
		JsonSchemaCommon: t.Merge(internal.JsonSchemaCommon).(JsonSchemaCommon),
		ExclusiveMaximum: internal.ExclusiveMaximum.OrFirst(t.ExclusiveMaximum),
		ExclusiveMinimum: internal.ExclusiveMinimum.OrFirst(t.ExclusiveMinimum),
		Maximum:          internal.Maximum.OrFirst(t.Maximum),
		Minimum:          internal.Minimum.OrFirst(t.Minimum),
		MultipleOf:       internal.MultipleOf.OrFirst(t.MultipleOf),
	}
}

type JsonSchemaInteger struct {
	jsonSchemaNumeric[int64]
}

func NewInteger(id string) JsonSchema {
	return newNumeric[int64](id, schematype.Integer)
}

type JsonSchemaNumber struct {
	jsonSchemaNumeric[float64]
}

func NewNumber(id string) JsonSchema {
	return newNumeric[float64](id, schematype.Number)
}
