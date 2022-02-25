package jsonschema_test

import (
	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/optioner"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resolve", func() {
	var (
	// system *jsonschema.System
	)

	BeforeEach(func() {
		// system = jsonschema.New()

	})

	It("should resolve", func() {
		schema := &jsonschema.JsonSchmea{
			Id: optioner.Some("test-id"),
			
		}

		Expect(schema.Id.IsDefined()).To(BeTrue())
	})
})
