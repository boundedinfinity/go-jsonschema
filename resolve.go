package jsonschema

import (
	"fmt"

	"github.com/boundedinfinity/jsonschema/objecttype"
)

func (t *System) Add(schema *JsonSchmea) error {
	if schema.Id.IsEmpty() {
		return ErrIdEmpty
	}

	id := schema.Id.Get()

	if _, ok := t.Map[id]; ok {
		return ErrDuplicateDefV(id)
	} else {
		t.Map[id] = schema
	}

	for name, obj := range schema.Defs {
		qname, err := urlJoin(id, "$defs", name)

		if err != nil {
			return err
		}

		if _, ok := t.Map[qname]; ok {
			return ErrDuplicateDefV(qname)
		}

		t.Map[qname] = obj
	}

	switch schema.Type {
	case objecttype.String, objecttype.Number, objecttype.Integer, objecttype.Boolean, objecttype.Null, objecttype.Array:
		t.Map[id] = schema
	case objecttype.Object:
		for name, obj := range schema.Properties {
			qname, err := urlJoin(id, name)

			if err != nil {
				return err
			}

			if _, ok := t.Map[qname]; ok {
				return ErrDuplicateDefV(qname)
			}

			t.Map[qname] = obj
		}
	}

	return nil
}

func (t *System) Exist(id string) bool {
	if id == "" {
		return false
	}

	_, ok := t.Map[id]
	return ok
}

func (t *System) Assert(id string) error {
	if id == "" {
		return ErrIdEmpty
	}

	if _, ok := t.Map[id]; !ok {
		return ErrSchemaNotFoundV(id)
	}

	return nil
}

func (t *System) ResolveAll() error {
	for name, schema := range t.Map {
		fmt.Println(name)
		if err := t.Resolve(schema); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) Resolve(schema *JsonSchmea) error {
	if schema.Ref.IsDefined() {
		id := schema.Ref.Get()
		if ref, ok := t.Map[id]; ok {
			*schema = *ref
		} else {
			return ErrRefNotFoundV(schema.Ref.Get())
		}
	} else {
		switch schema.Type {
		case objecttype.Array:
			if err := t.Resolve(schema.Items); err != nil {
				return err
			}
		}
	}

	return nil
}
