package jsonschema

import (
	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/optional"
)

func (t *System) mapSchema(schema *JsonSchmea, file string) error {
	id := schema.Id.Get()

	if _, ok := t.i2l[id]; ok {
		return ErrDuplicateDefV(id)
	} else {
		t.i2l[id] = file
	}

	if _, ok := t.l2i[id]; ok {
		return ErrDuplicateDefV(id)
	} else {
		t.l2i[file] = id
	}

	if _, ok := t.schmeaMap[id]; ok {
		return ErrDuplicateDefV(id)
	} else {
		t.schmeaMap[id] = schema
	}

	for name, obj := range schema.Defs {
		qname, err := urlJoin(id, "$defs", name)

		if err != nil {
			return err
		}

		if _, ok := t.schmeaMap[qname]; ok {
			return ErrDuplicateDefV(qname)
		}

		t.schmeaMap[qname] = obj
	}

	switch schema.Type {
	case objecttype.String, objecttype.Number, objecttype.Integer, objecttype.Boolean, objecttype.Null, objecttype.Array:
		t.schmeaMap[id] = schema
	case objecttype.Object:
		for name, obj := range schema.Properties {
			qname, err := urlJoin(id, name)

			if err != nil {
				return err
			}

			if _, ok := t.schmeaMap[qname]; ok {
				return ErrDuplicateDefV(qname)
			}

			t.schmeaMap[qname] = obj
		}
	}

	return nil
}

func (t *System) resolveSchema(schema *JsonSchmea) error {
	if schema.Ref.IsDefined() {
		id := schema.Ref.Get()
		if ref, ok := t.schmeaMap[id]; ok {
			*schema = *ref
		} else {
			return ErrRefNotFoundV(schema.Ref.Get())
		}
	} else {
		switch schema.Type {
		case objecttype.Array:
			if err := t.resolveSchema(schema.Items); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *System) getRef(id optional.StringOptional, ref optional.StringOptional) (*JsonSchmea, error) {
	// var key string

	// if strings.HasPrefix(ref.String(), "#") {
	// 	idUrl, err := url.Parse(id.Get())

	// 	if err != nil {
	// 		return &JsonSchmea{}, err
	// 	}

	// 	idUrl.Path = path.Join(idUrl.Path, strings.TrimPrefix(ref.Get(), "#"))
	// 	key = idUrl.String()
	// } else {
	// 	key = ref.String()
	// }

	// if obj, ok := t.typeMap[key]; ok {
	// 	return obj, nil
	// }

	return &JsonSchmea{}, ErrRefNotFoundV(ref.Get())
}
