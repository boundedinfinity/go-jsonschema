package jsonschema

import (
	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/optional"
)

func validateInteger(v optional.Float64Optional) error {
	if v.IsEmpty() {
		return nil
	}

	x := float64(int64(v.Get()))

	if x != v.Get() {
		return ErrNotInteger
	}

	return nil
}

func validateNonNegative(v optional.IntOptional) error {
	if v.IsEmpty() {
		return nil
	}

	if v.Get() < 0 {
		return ErrNotNonNegative
	}

	return nil
}

func (t JsonSchmea) Validate() error {
	switch t.Type {
	case objecttype.Array:
		if t.Items == nil {
			return ErrInvalidArrayMissingItems
		}

		if err := validateNonNegative(t.MinItems); err != nil {
			return err
		}

		if err := validateNonNegative(t.MaxItems); err != nil {
			return err
		}

		if err := validateNonNegative(t.MaxContains); err != nil {
			return err
		}

		if err := t.Items.Validate(); err != nil {
			return err
		}
	case objecttype.Integer:
		if t.MultipleOf.IsDefined() {
			if t.MultipleOf.Get() <= 0 {
				return ErrInvalidMultipleOf
			}

			if err := validateInteger(t.MultipleOf); err != nil {
				return err
			}
		}

		if t.Maximum.IsDefined() {
			if err := validateInteger(t.Maximum); err != nil {
				return err
			}
		}

		if t.ExclusiveMaximum.IsDefined() {
			if err := validateInteger(t.ExclusiveMaximum); err != nil {
				return err
			}
		}

		if t.Minimum.IsDefined() {
			if err := validateInteger(t.Minimum); err != nil {
				return err
			}
		}

		if t.ExclusiveMinimum.IsDefined() {
			if err := validateInteger(t.ExclusiveMinimum); err != nil {
				return err
			}
		}
	case objecttype.Number:
		if t.MultipleOf.IsDefined() && t.MultipleOf.Get() <= 0 {
			return ErrInvalidMultipleOf
		}
	case objecttype.String:
		if err := validateNonNegative(t.MaxLength); err != nil {
			return err
		}

		if err := validateNonNegative(t.MinLength); err != nil {
			return err
		}
	case objecttype.Object:
		for _, property := range t.Properties {
			if err := property.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
