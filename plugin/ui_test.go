package plugin_test

import (
	"errors"
	"io"

	"github.com/sclevine/cflocal/plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("UI", func() {
	var (
		out, err, in *gbytes.Buffer
		ui           *plugin.UI
	)

	BeforeEach(func() {
		out = gbytes.NewBuffer()
		err = gbytes.NewBuffer()
		in = gbytes.NewBuffer()
		ui = &plugin.UI{Out: out, Err: err, In: in}
	})

	Describe("#Prompt", func() {
		It("should output the prompt and return the user's entry", func() {
			io.WriteString(in, "some answer\n")
			response := ui.Prompt("some question")
			Expect(out).To(gbytes.Say("some question"))
			Expect(response).To(Equal("some answer"))
		})

		Context("when the input cannot be read", func() {
			It("should output the prompt and return an empty string", func() {
				response := ui.Prompt("some question")
				Expect(out).To(gbytes.Say("some question"))
				Expect(response).To(BeEmpty())
			})
		})
	})

	Describe("#Output", func() {
		It("should output the provided format string", func() {
			ui.Output("%s format", "some")
			Expect(out).To(gbytes.Say("some format"))
		})
	})

	Describe("#Error", func() {
		It("should output the provided error as an error followed by FAILED", func() {
			ui.Error(errors.New("some error"))
			Expect(err).To(gbytes.Say("Error: some error"))
			Expect(out).To(gbytes.Say("FAILED"))
		})
	})
})