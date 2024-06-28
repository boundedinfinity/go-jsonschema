package idiomatic

type JsonSchema interface {
	GetId() string
	GetType() string
	Validate() error
	Copy() JsonSchema
}
