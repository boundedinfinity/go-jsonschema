package idiomatic

import (
	"github.com/boundedinfinity/go-jsonschema/model"
	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

func validateMaxMin[T number](max, min T) error {
	if max < min {
		return model.ErrJsonSchemaNumberInvalidMaxMinv(max, min)
	}

	return nil
}

func validateMultipleOf[T number](v int) error {
	if v < 0 {
		return model.ErrJsonSchemaNumberMultipleOfInvalidv(v)
	}

	return nil
}

func validateMaxMinContains[T number](max, min T) error {
	if max < 0 {
		return model.ErrJsonSchemaArrayMaxContainsInvalidv(max)
	}

	if min < 0 {
		return model.ErrJsonSchemaArrayMinContainsInvalidv(min)
	}

	if max < min {
		return model.ErrJsonSchemaArrayInvalidMaxMinv(min, max)
	}

	return nil
}

func validateMaxMinLength[T number](max, min T) error {
	if max < 0 {
		return model.ErrJsonSchemaStringMaxLengthv(max)
	}

	if min < 0 {
		return model.ErrJsonSchemaStringMinLengthv(min)
	}

	if max < min {
		return model.ErrJsonSchemaStringMaxMinLengthv(max, min)
	}

	return nil
}

func validateMaxMinProperties[T number](max, min T) error {
	if max < 0 {
		return model.ErrJsonSchemaObjectMaxPropertiesv(max)
	}

	if min < 0 {
		return model.ErrJsonSchemaObjectMinPropertiesv(min)
	}

	if max < min {
		return model.ErrJsonSchemaObjectMaxMinPropertiesv(max, min)
	}

	return nil
}
