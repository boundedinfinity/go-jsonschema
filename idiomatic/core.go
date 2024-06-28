package idiomatic

import "github.com/boundedinfinity/go-commoner/idiomatic/mapper"

// https://www.learnjsonschema.com/2020-12/

type Uri string

type JsonSchemaCore struct {
	Schema        string                `json:"$schema,omitempty" yaml:"$schema,omitempty"`
	Id            string                `json:"$id,omitempty" yaml:"$id,omitempty"`
	Ref           string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Comment       string                `json:"$comment,omitempty" yaml:"$comment,omitempty"`
	Defs          map[string]JsonSchema `json:"$defs,omitempty" yaml:"$defs,omitempty"`
	Anchor        string                `json:"$anchor,omitempty" yaml:"$anchor,omitempty"`
	DynamicAnchor string                `json:"$dynamicAnchor,omitempty" yaml:"$dynamicAnchor,omitempty"`
	DynamicRef    string                `json:"$dynamicRef,omitempty" yaml:"$dynamicRef,omitempty"`
	Vocabulary    map[Uri]bool          `json:"$vocabulary,omitempty" yaml:"$vocabulary,omitempty"`
	Title         string                `json:"title,omitempty" yaml:"title,omitempty"`
	Description   string                `json:"description,omitempty" yaml:"description,omitempty"`
	Default       map[string]any        `json:"default,omitempty" yaml:"default,omitempty"`
	Deprecated    bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Examples      []map[string]any      `json:"examples,omitempty" yaml:"examples,omitempty"`
	ReadOnly      bool                  `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly     bool                  `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	AnyOf         []JsonSchema          `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	OneOf         []JsonSchema          `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AllOf         []JsonSchema          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	Then          JsonSchema            `json:"then,omitempty" yaml:"then,omitempty"`
	If            JsonSchema            `json:"if,omitempty" yaml:"if,omitempty"`
	Else          JsonSchema            `json:"else,omitempty" yaml:"else,omitempty"`
	Not           JsonSchema            `json:"not,omitempty" yaml:"not,omitempty"`
	Enum          []any                 `json:"enum,omitempty" yaml:"enum,omitempty"`
	Const         []any                 `json:"const,omitempty" yaml:"const,omitempty"`
}

func (t JsonSchemaCore) Copy() JsonSchemaCore {
	examples := []map[string]any{}
	for _, example := range t.Examples {
		examples = append(examples, mapper.Copy(example))
	}

	return JsonSchemaCore{
		Schema:        t.Schema,
		Id:            t.Id,
		Ref:           t.Ref,
		Comment:       t.Comment,
		Defs:          copySchemaMap(t.Defs),
		Anchor:        t.Anchor,
		DynamicAnchor: t.DynamicAnchor,
		DynamicRef:    t.DynamicRef,
		Vocabulary:    mapper.Copy(t.Vocabulary),
		Title:         t.Title,
		Description:   t.Description,
		Default:       mapper.Copy(t.Default),
		Examples:      examples,
		ReadOnly:      t.ReadOnly,
		WriteOnly:     t.WriteOnly,
		AnyOf:         copySchemaArray(t.AnyOf),
		OneOf:         copySchemaArray(t.OneOf),
		AllOf:         copySchemaArray(t.AllOf),
		Then:          t.Then.Copy(),
		If:            t.If.Copy(),
		Else:          t.Else.Copy(),
		Not:           t.Not.Copy(),
		Enum:          t.Enum[:],
		Const:         t.Const[:],
	}
}

func (t JsonSchemaCore) Validate() error {
	return nil
}
