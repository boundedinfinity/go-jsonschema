package jsonschema

import "github.com/boundedinfinity/go-jsonschema/model"

func (t *System) Check() error {
	if err := t.check1(); err != nil {
		return err
	}

	if err := t.check3(); err != nil {
		return err
	}

	if err := t.check2(); err != nil {
		return err
	}

	return nil
}

func (t *System) check1() error {
	for _, schema := range t.pathMap {
		id := schema.GetId()

		if id.Empty() {
			return model.ErrSchemaIdNotFound
		}

		if t.idMap.Has(id.Get()) {
			return model.ErrSchemaIdDuplicatev(id.Get())
		}

		t.idMap[id.Get()] = schema
	}

	return nil
}

func (t *System) check2() error {
	for _, schema := range t.idMap {
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

		if t.idMap.Get(c.Ref.Get()).Empty() {
			return model.ErrRefNotFoundv(c.Ref.Get())
		}
	case model.JsonSchemaRef:
		if c.Ref.Empty() {
			return model.ErrRefEmpty
		}

		if t.idMap.Get(c.Ref.Get()).Empty() {
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

func (t *System) check3() error {
	for _, schema := range t.idMap {
		if err := schema.Validate(); err != nil {
			return err
		}
	}

	return nil
}
