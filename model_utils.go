package jsonschema

import (
	"github.com/boundedinfinity/commons/slices"
	"github.com/boundedinfinity/go-trier"
	"github.com/boundedinfinity/optioner"
)

func (js JsonSchema) Copy(name string) JsonSchema {
	copy := JsonSchema{
		system: js.system,
		parent: js.parent,
		Anchor: js.Anchor,
	}

	return copy
}

func (js JsonSchema) GetProperty(name string) optioner.Option[JsonSchema] {
	if prop, ok := js.Properties[name]; ok {
		return optioner.Some(*prop)
	}

	return optioner.None[JsonSchema]()
}

func (js JsonSchema) IsRequired() trier.Try[bool] {
	return trier.Success(false)
}

func (js JsonSchema) IsPropRequired(propname string) trier.Try[bool] {

	if js.Type.Empty() {
		return trier.Failure[bool](ErrObjectTypeEmpty)
	}

	if len(js.Properties) <= 0 {
		return trier.Success(false)
	}

	if len(js.Required) <= 0 {
		return trier.Success(false)
	}

	_, ok := js.Properties[propname]

	if !ok {
		trier.Failure[bool](ErrPropertyNotFoundV(propname))
	}

	if slices.Contains(js.Required, propname) {
		return trier.Success(true)
	}

	return trier.Success(false)
}
