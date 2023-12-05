package idiomatic

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/go-jsonschema/model"
)

type JsonSchemaTypeMap struct {
	m mapper.Mapper[string, JsonSchema]
}

func (t *JsonSchemaTypeMap) Register(schema JsonSchema) error {
	common := schema.Common()

	if common == nil || stringer.IsEmpty(common.Id) {
		return nil
	}

	if _, ok := t.m[common.Id]; !ok {
		return model.ErrJsonSchemaIdDuplicate
	}

	t.m[common.Id] = schema

	return nil
}

func (t JsonSchemaTypeMap) Lookup(schema JsonSchema) (JsonSchema, bool) {
	common := schema.Common()

	if common == nil {
		return nil, false
	}

	if !stringer.IsEmpty(common.Id) {
		return schema, true
	}

	if !stringer.IsEmpty(common.Ref) {
		if found, ok := t.m[common.Ref]; ok {
			return found, true
		}
	}

	return nil, false
}
