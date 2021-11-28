package objecttype

import (
	"encoding/json"
	"fmt"
)

// ObjectTypeOption contains a initialized or empty copy of the ObjectType type.
type ObjectTypeOption struct {
	v *ObjectType
}

// NewObjectTypeValue creates a new ObjectTypeOption type with an initialized value.
func NewObjectTypeValue(v ObjectType) ObjectTypeOption {
	return ObjectTypeOption{
		v: &v,
	}
}

// NewObjectTypeValue creates a new ObjectTypeOption type with the given value.
func NewObjectTypePtr(v *ObjectType) ObjectTypeOption {
	return ObjectTypeOption{
		v: v,
	}
}

// ObjectType2Ptr create a pointer from the given type.
func ObjectType2Ptr(v ObjectType) *ObjectType {
	return &v
}

// NewObjectTypeValue creates a new ObjectTypeOption type with an empty value.
func NewObjectTypeEmpty() ObjectTypeOption {
	return ObjectTypeOption{}
}

// NewObjectTypeEmptyIfZeroValue creates a new initialized ObjectTypeOption type if the input ObjectType doesn't equal the ObjectType's zero value, or an empty ObjectTypeOption otherwise.
func NewObjectTypeEmptyIfZeroValue(v ObjectType) ObjectTypeOption {
	var e ObjectType

	if v == e {
		return NewObjectTypeEmpty()
	}

	return NewObjectTypeValue(v)
}

//IsEmpty returns true if the contained ObjectType value is empty, false otherwise.
func (t ObjectTypeOption) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained ObjectType value is initialized, false otherwise.
func (t ObjectTypeOption) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained ObjectType value.
//NOTE: If the value is empty, this will return the ObjectType zero value.
func (t ObjectTypeOption) Get() ObjectType {
	var v ObjectType

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained ObjectType value if the contained value is initialized or the input value is the value is not initialized.
func (t ObjectTypeOption) GetOrElse(v ObjectType) ObjectType {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the ObjectTypeOption type into the JSON representation.
func (t ObjectTypeOption) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the ObjectTypeOption type.
func (t *ObjectTypeOption) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v ObjectType

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the ObjectTypeOption type into the YAML representation.
func (t ObjectTypeOption) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the ObjectTypeOption type.
func (t *ObjectTypeOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v ObjectType

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *ObjectTypeOption) String() string {
	return fmt.Sprintf("%v", t.Get())
}
