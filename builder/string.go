package builder

import (
	"github.com/boundedinfinity/go-commoner/mapper"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
)

type stringBuilder struct {
	baseBuilder
	Schema *model.JsonSchemaString
}

func String() *stringBuilder {
	s := model.NewString().(model.JsonSchemaString)

	return &stringBuilder{
		baseBuilder: baseBuilder{base: s.Base()},
		Schema:      &s,
	}
}

func (b *stringBuilder) Format(v stringformat.StringFormat) *stringBuilder {
	b.Schema.Format = o.Some(v)
	return b
}

func (b *stringBuilder) MaxLength(v int) *stringBuilder {
	b.Schema.MaxLength = o.Some(v)
	return b
}

func (b *stringBuilder) MinLength(v int) *stringBuilder {
	b.Schema.MinLength = o.Some(v)
	return b
}

func (b *stringBuilder) Pattern(v string) *stringBuilder {
	b.Schema.Pattern = o.Some(v)
	return b
}

func (b *stringBuilder) Enum(v []string) *stringBuilder {
	b.Schema.Enum = o.Some(v)
	return b
}

func (b *stringBuilder) EnumDescription(v mapper.Mapper[string, string]) *stringBuilder {
	b.Schema.EnumDescription = o.Some(v)
	return b
}
