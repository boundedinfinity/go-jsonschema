package model

import o "github.com/boundedinfinity/go-commoner/optioner"

type jsonSchemaNumeric[T int64 | float64] struct {
	JsonSchemaCommon
	ExclusiveMaximum o.Option[T] `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum o.Option[T] `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          o.Option[T] `json:"maximum" yaml:"maximum"`
	Minimum          o.Option[T] `json:"minimum" yaml:"minimum"`
	MultipleOf       o.Option[T] `json:"multipleOf" yaml:"multipleOf"`
}

type JsonSchemaInteger struct {
	jsonSchemaNumeric[int64]
}

type JsonSchemaNumber struct {
	jsonSchemaNumeric[float64]
}
