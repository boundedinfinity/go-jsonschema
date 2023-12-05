package idiomatic

import (
	"github.com/boundedinfinity/go-jsonschema/model"
)

func validateMaxMin[T ~int64 | ~float64](max, min T) error {
	if max < min {
		return model.ErrJsonSchemaNumberInvalidMaxMinv("max: %v, min: %v", max, min)
	}

	return nil
}

func validateMultipleOf[T ~int64 | ~float64](v T) error {
	if v < 0 {
		return model.ErrJsonSchemaNumberMultipleOfInvalidv("%v", v)
	}

	return nil
}

func validateMaxMinContains[T ~int64](max, min T) error {
	if max < 0 {
		return model.ErrJsonSchemaArrayMaxContainsInvalidv("%v", max)
	}

	if min < 0 {
		return model.ErrJsonSchemaArrayMinContainsInvalidv("%v", min)
	}

	if max < min {
		return model.ErrJsonSchemaArrayMinContainsInvalidv("max: %v, min: %v", max, min)
	}

	return nil
}

func validateMaxMinLength[T ~int](max, min T) error {
	if max < 0 {
		return model.ErrJsonSchemaStringMaxLengthv("%v", max)
	}

	if min < 0 {
		return model.ErrJsonSchemaStringMinLengthv("%v", min)
	}

	if max < min {
		return model.ErrJsonSchemaStringMaxMinLengthv("max: %v, min: %v", max, min)
	}

	return nil
}

func validateMaxMinProperties[T ~int64](max, min T) error {
	if max < 0 {
		return model.ErrJsonSchemaObjectMaxPropertiesv("%v", max)
	}

	if min < 0 {
		return model.ErrJsonSchemaObjectMinPropertiesv("%v", min)
	}

	if max < min {
		return model.ErrJsonSchemaObjectMaxMinPropertiesv("max: %v, min: %v", max, min)
	}

	return nil
}
