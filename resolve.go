package jsonschema

import (
	"github.com/boundedinfinity/go-jsonschema/model"
)

func (t *System) Resolve() error {
	for k, v := range t.schemaMap {
		if v.GetType().Empty() && v.GetRef().Empty() {
			return model.ErrSchemaTypeOrRefNotFound
		}

		if v.GetType().Defined() {
			continue
		}

		ref := t.schemaMap.Get(v.GetRef().Get())

		if ref.Empty() {
			return model.ErrRefNotFoundv(v.GetRef().Get())
		}

		resolved := v.Merge(ref.Get())
		t.schemaMap[k] = resolved
	}

	return nil
}
