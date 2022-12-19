package builder

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
)

type baseBuilder struct {
	base *model.JsonSchemaBase
}

func (b *baseBuilder) Comment(v string) *baseBuilder {
	b.base.Comment = o.Some(v)
	return b
}

func (b *baseBuilder) Deprecated(v bool) *baseBuilder {
	b.base.Deprecated = o.Some(v)
	return b
}

func (b *baseBuilder) Description(v string) *baseBuilder {
	b.base.Description = o.Some(v)
	return b
}

func (b *baseBuilder) Id(v string) *baseBuilder {
	b.base.Id = o.Some(v)
	return b
}

func (b *baseBuilder) ReadOnly(v bool) *baseBuilder {
	b.base.ReadOnly = o.Some(v)
	return b
}

func (b *baseBuilder) Title(v string) *baseBuilder {
	b.base.Title = o.Some(v)
	return b
}

func (b *baseBuilder) WriteOnly(v bool) *baseBuilder {
	b.base.WriteOnly = o.Some(v)
	return b
}
