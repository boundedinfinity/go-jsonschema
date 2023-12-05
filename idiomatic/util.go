package idiomatic

import (
	"github.com/boundedinfinity/go-jsonschema/model"
)

func validateMaxMin[T ~int64 | ~float64](max, min T) error {
	if max < min {
		return model.ErrNumberInvalidMaxMinv("max: %v, min: %v", max, min)
	}

	return nil
}

func validateMultipleOf[T ~int64 | ~float64](v T) error {
	if v < 0 {
		return model.ErrNumberMultipleOfInvalidv("%v", v)
	}

	return nil
}

func validateMaxMinContains[T ~int64](max, min T) error {
	if max < 0 {
		return model.ErrArrayMaxContainsInvalidv("%v", max)
	}

	if min < 0 {
		return model.ErrArrayMinContainsInvalidv("%v", min)
	}

	if max < min {
		return model.ErrArrayMinContainsInvalidv("max: %v, min: %v", max, min)
	}

	return nil
}

func validateMaxMinLength[T ~int](max, min T) error {
	if max < 0 {
		return model.ErrStringMaxLengthv("%v", max)
	}

	if min < 0 {
		return model.ErrStringMinLengthv("%v", min)
	}

	if max < min {
		return model.ErrStringMaxMinLengthv("max: %v, min: %v", max, min)
	}

	return nil
}

func validateMaxMinProperties[T ~int64](max, min T) error {
	if max < 0 {
		return model.ErrObjectMaxPropertiesv("%v", max)
	}

	if min < 0 {
		return model.ErrObjectMinPropertiesv("%v", min)
	}

	if max < min {
		return model.ErrObjectMaxMinPropertiesv("max: %v, min: %v", max, min)
	}

	return nil
}
