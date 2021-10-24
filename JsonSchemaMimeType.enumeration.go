//************************************************************************************
//*                                                                                  *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package jsonschema

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JsonSchemaMimeType string

const (
	JsonSchemaMimeType_Xxx JsonSchemaMimeType = "xxx"
)

var (
	JsonSchemaMimeTypeList = []JsonSchemaMimeType{
		JsonSchemaMimeType_Xxx,
	}
)

func IsJsonSchemaMimeType(v string) bool {
	var f bool

	for _, e := range JsonSchemaMimeTypeList {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func JsonSchemaMimeTypeParse(v string) (JsonSchemaMimeType, error) {
	var o JsonSchemaMimeType
	var f bool
	n := strings.ToLower(v)

	for _, e := range JsonSchemaMimeTypeList {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, ErrJsonSchemaMimeTypeNotFound(v)
	}

	return o, nil
}

func JsonSchemaMimeTypeListToStrings(vs []JsonSchemaMimeType) []string {
	var ss []string

	for _, v := range vs {
		ss = append(ss, v.String())
	}

	return ss
}

func ErrJsonSchemaMimeTypeNotFound(v string) error {
	var ss []string

	for _, e := range JsonSchemaMimeTypeList {
		ss = append(ss, string(e))
	}

	return fmt.Errorf(
		"invalid enumeration type '%v', must be one of %v",
		v, strings.Join(ss, ","),
	)
}

func (t JsonSchemaMimeType) String() string {
	return string(t)
}

func (t JsonSchemaMimeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *JsonSchemaMimeType) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := JsonSchemaMimeTypeParse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}