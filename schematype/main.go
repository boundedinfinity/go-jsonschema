package schematype

//go:generate enumer -path=./main.go

type SchemaType string

const (
	String  SchemaType = "string"
	Number  SchemaType = "number"
	Integer SchemaType = "integer"
	Object  SchemaType = "object"
	Array   SchemaType = "array"
	Boolean SchemaType = "boolean"
	Null    SchemaType = "null"
)
