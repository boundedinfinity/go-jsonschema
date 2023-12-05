package idiomatic

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type JsonSchema interface {
	marshaler.TypeNamer
	Common() *JsonSchemaCommon
	Validate() error
	Copy() JsonSchema
}

type jsonSchemaDescriminator struct {
	Type string `json:"type" yaml:"type"`
	Ref  string `json:"$ref" yaml:"$ref"`
}

type JsonSchemaCommon struct {
	Id          string   `json:"$id" yaml:"$id"`
	Type        string   `json:"type" yaml:"type"`
	Ref         string   `json:"$ref" yaml:"$ref"`
	Schema      string   `json:"$schema" yaml:"$schema"`
	Title       string   `json:"title" yaml:"title"`
	Description string   `json:"description" yaml:"description"`
	ReadOnly    bool     `json:"readOnly" yaml:"readOnly"`
	WriteOnly   bool     `json:"writeOnly" yaml:"writeOnly"`
	Deprecated  bool     `json:"deprecated" yaml:"deprecated"`
	Examples    []string `json:"examples" yaml:"examples"`
	Comment     string   `json:"$comment" yaml:"$comment"`
}

func (t JsonSchemaCommon) Copy() JsonSchemaCommon {
	return JsonSchemaCommon{
		Id:          t.Id,
		Type:        t.Title,
		Ref:         t.Ref,
		Schema:      t.Schema,
		Title:       t.Title,
		Description: t.Description,
		ReadOnly:    t.ReadOnly,
		WriteOnly:   t.WriteOnly,
		Deprecated:  t.Deprecated,
		Examples:    t.Examples,
		Comment:     t.Comment,
	}
}

func (t *JsonSchemaCommon) Merge(other JsonSchemaCommon) {
	t.Id = mergeSimple(other.Id, t.Id)
	t.Type = mergeSimple(other.Type, t.Type)
	t.Ref = mergeSimple(other.Ref, t.Ref)
	t.Schema = mergeSimple(other.Schema, t.Schema)
	t.Title = mergeSimple(other.Title, t.Title)
	t.Description = mergeDescription(t.Description, other.Description)
	t.ReadOnly = mergeSimple(other.ReadOnly, t.ReadOnly)
	t.WriteOnly = mergeSimple(other.WriteOnly, t.WriteOnly)
	t.Deprecated = mergeSimple(other.Deprecated, t.Deprecated)
	t.Examples = append(other.Examples, t.Examples...)
	t.Comment = mergeSimple(other.Comment, t.Comment)
}

func (t JsonSchemaCommon) Validate() error {
	return nil
}

func mergeSimple[T comparable](list ...T) T {
	v, _ := slicer.FirstNotZero(list...)
	return v
}

func mergeDescription(a, b string) string {
	var d string

	if !stringer.IsEmpty(a) && !stringer.IsEmpty(b) {
		d = fmt.Sprintf("%v\n%v", b, a)
	}

	if !stringer.IsEmpty(b) {
		d = b
	}

	return d
}
