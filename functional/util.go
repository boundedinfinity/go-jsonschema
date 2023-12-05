package functional

// import (
// 	"fmt"

// 	o "github.com/boundedinfinity/go-commoner/optioner"
// )

// func validateMaxMin[T ~int64 | ~float64](max, min o.Option[T]) error {
// 	if max.Defined() && min.Defined() {
// 		if max.Get() < min.Get() {
// 			message := fmt.Sprintf("max: %v, min: %v", max.Get(), min.Get())
// 			return ErrNumberInvalidMaxMinv(message)
// 		}
// 	}

// 	return nil
// }

// func validateMultipleOf[T ~int64 | ~float64](v o.Option[T]) error {
// 	if v.Defined() && v.Get() < 0 {
// 		return ErrNumberMultipleOfInvalidv(v.Get())
// 	}

// 	return nil
// }

// func validateMaxMinContains[T ~int64](max, min o.Option[T]) error {
// 	if max.Defined() && max.Get() < 0 {
// 		return ErrArrayMaxContainsInvalidv(max.Get())
// 	}

// 	if min.Defined() && min.Get() < 0 {
// 		return ErrArrayMinContainsInvalidv(min.Get())
// 	}

// 	if max.Defined() && min.Defined() {
// 		if max.Get() < min.Get() {
// 			message := fmt.Sprintf("max: %v, min: %v", max.Get(), min.Get())
// 			return ErrArrayMinContainsInvalidv(message)
// 		}
// 	}

// 	return nil
// }

// func validateMaxMinLength[T ~int](max, min o.Option[T]) error {
// 	if max.Defined() && max.Get() < 0 {
// 		return ErrStringMaxLengthv(max.Get())
// 	}

// 	if min.Defined() && min.Get() < 0 {
// 		return ErrStringMinLengthv(min.Get())
// 	}

// 	if max.Defined() && min.Defined() {
// 		if max.Get() < min.Get() {
// 			message := fmt.Sprintf("max: %v, min: %v", max.Get(), min.Get())
// 			return ErrStringMaxMinLengthv(message)
// 		}
// 	}

// 	return nil
// }

// func validateMaxMinProperties[T ~int64](max, min o.Option[T]) error {
// 	if max.Defined() && max.Get() < 0 {
// 		return ErrObjectMaxPropertiesv(max.Get())
// 	}

// 	if min.Defined() && min.Get() < 0 {
// 		return ErrObjectMinPropertiesv(min.Get())
// 	}

// 	if max.Defined() && min.Defined() {
// 		if max.Get() < min.Get() {
// 			message := fmt.Sprintf("max: %v, min: %v", max.Get(), min.Get())
// 			return ErrObjectMaxMinPropertiesv(message)
// 		}
// 	}

// 	return nil
// }
