package jsonschema

import (
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-jsonschema/model"
)

func (t *System) Register(root, source string, schema model.JsonSchema) error {
	t.All = append(t.All, schema)
	t.source2root[source] = root
	t.source2schema[source] = schema
	add(t.root2source, root, source)

	id := schema.Base().Id
	if id.Defined() {
		if t.id2schema.Has(schema.Base().Id.Get()) {
			return model.ErrPathDuplicatev(schema.Base().Id.Get())
		}

		t.id2schema[id.Get()] = schema
		t.id2root[id.Get()] = root
		add(t.root2id, root, id.Get())
		t.id2source[id.Get()] = source
		t.source2id[source] = id.Get()
	}

	return nil
}

func add[T any](m mapper.Mapper[string, []T], k string, v T) {
	if !m.Has(k) {
		m[k] = make([]T, 0)
	}

	m[k] = append(m[k], v)
}
