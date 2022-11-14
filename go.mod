module github.com/boundedinfinity/go-jsonschema

go 1.18

require (
	github.com/boundedinfinity/go-commoner v1.0.17
	github.com/boundedinfinity/go-marshaler v1.0.3
	github.com/boundedinfinity/mimetyper v1.0.10
	github.com/stretchr/testify v1.8.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/boundedinfinity/go-commoner => ../go-commoner

replace github.com/boundedinfinity/go-marshaler => ../go-marshaler
