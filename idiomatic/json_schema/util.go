package json_schema

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"golang.org/x/exp/constraints"
)

func copySchemaMap(m map[string]JsonSchema) map[string]JsonSchema {
	results := map[string]JsonSchema{}

	for k, v := range m {
		results[k] = v.Copy()
	}

	return results
}

func copySchemaArray(elems []JsonSchema) []JsonSchema {
	var results []JsonSchema

	for _, elem := range elems {
		results = append(results, elem.Copy())
	}

	return results
}

type number interface {
	constraints.Integer | constraints.Float
}

func validateMaxMin[T number](max, min T, err *errorer.Errorer) error {
	// if max < min {
	// 	return model.ErrJsonSchemaNumberInvalidMaxMinv(max, min)
	// }

	return nil
}

func validateMultipleOf[T number](v T, err *errorer.Errorer) error {
	if v < 0 {
		return err.WithValue(v)
	}

	return nil
}

func validateMaxMinContains[T number](max, min T, err *errorer.Errorer) error {
	// if max < 0 {
	// 	return model.ErrJsonSchemaArrayMaxContainsInvalidv(max)
	// }

	// if min < 0 {
	// 	return model.ErrJsonSchemaArrayMinContainsInvalidv(min)
	// }

	// if max < min {
	// 	return model.ErrJsonSchemaArrayInvalidMaxMinv(min, max)
	// }

	return nil
}

// func validateMaxMinLength[T number](max, min T) error {
// 	if max < 0 {
// 		return model.ErrJsonSchemaStringMaxLengthv(max)
// 	}

// 	if min < 0 {
// 		return model.ErrJsonSchemaStringMinLengthv(min)
// 	}

// 	if max < min {
// 		return model.ErrJsonSchemaStringMaxMinLengthv(max, min)
// 	}

// // 	if min < 0 {
// // 		return model.ErrJsonSchemaObjectMinPropertiesv("%v", min)
// // 	}

// func validateMaxMinProperties[T number](max, min T) error {
// 	if max < 0 {
// 		return model.ErrJsonSchemaObjectMaxPropertiesv(max)
// 	}

// 	if min < 0 {
// 		return model.ErrJsonSchemaObjectMinPropertiesv(min)
// 	}

// 	if max < min {
// 		return model.ErrJsonSchemaObjectMaxMinPropertiesv(max, min)
// 	}

// 	return nil
// }
