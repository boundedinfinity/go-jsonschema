module github.com/boundedinfinity/jsonschema

go 1.18

require (
	github.com/boundedinfinity/mimetyper v0.0.0-20211121204855-2eb5ba3c3d07
	github.com/boundedinfinity/optioner v0.0.0-20211128195323-aef6932f5324
	github.com/onsi/ginkgo/v2 v2.1.3
	github.com/onsi/gomega v1.18.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.6 // indirect
)

replace github.com/boundedinfinity/optioner => ../optioner
