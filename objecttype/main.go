package objecttype

//go:generate enumer -path=./main.go

type ObjectType string

const (
	String  ObjectType = "string"
	Number  ObjectType = "number"
	Integer ObjectType = "integer"
	Object  ObjectType = "object"
	Array   ObjectType = "array"
	Boolean ObjectType = "boolean"
	Null    ObjectType = "null"
)
