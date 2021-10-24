package jsonschema_test

import (
	"testing"

	"github.com/boundedinfinity/jsonschema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJsonSchema(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JsonSchema Suite")
}

var _ = Describe("JsonSchema", func() {
	var (
		system *jsonschema.System
	)

	BeforeEach(func() {
		system = jsonschema.New()
	})

	Describe("Loading schemas", func() {
		It("should load yaml", func() {
			// err := system.LoadSchema("file://test/address.schema.yaml")
			err := system.LoadSchema("file://test")

			Expect(err).To(BeNil())
		})
	})
})
