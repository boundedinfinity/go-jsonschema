package jsonschema

import "github.com/boundedinfinity/go-jsonschema/model"

func (t *System) Check() error {
	for _, schema := range t.source2schema {
		id := schema.Base().Id

		if id.Empty() {
			return model.ErrSchemaIdNotFound
		}

		if t.id2schema.Has(id.Get()) {
			return model.ErrSchemaIdDuplicatev(id.Get())
		}

		t.id2schema[id.Get()] = schema
	}

	for _, schema := range t.id2schema {
		if err := schema.Validate(); err != nil {
			return err
		}
	}

	for _, schema := range t.id2schema {
		if err := t.checkResolve(schema); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) checkResolve(schema model.JsonSchema) error {
	switch c := schema.(type) {
	case *model.JsonSchemaRef:
		if c.Ref.Empty() {
			return model.ErrRefEmpty
		}

		if t.id2schema.Get(c.Ref.Get()).Empty() {
			return model.ErrRefNotFoundv(c.Ref.Get())
		}
	case model.JsonSchemaRef:
		if c.Ref.Empty() {
			return model.ErrRefEmpty
		}

		if t.id2schema.Get(c.Ref.Get()).Empty() {
			return model.ErrRefNotFoundv(c.Ref.Get())
		}
	case *model.JsonSchemaArray:
		if c.Items.Empty() {
			return model.ErrArrayItemsEmpty
		}

		if err := t.checkResolve(c.Items.Get()); err != nil {
			return err
		}
	case model.JsonSchemaArray:
		if c.Items.Empty() {
			return model.ErrArrayItemsEmpty
		}

		if err := t.checkResolve(c.Items.Get()); err != nil {
			return err
		}
	case *model.JsonSchemaObject:
		for _, property := range c.Properties {
			if err := t.checkResolve(property); err != nil {
				return err
			}
		}
	case model.JsonSchemaObject:
		for _, property := range c.Properties {
			if err := t.checkResolve(property); err != nil {
				return err
			}
		}
	}

	return nil
}
