package tests

import (
	"github.com/rickschubert/html-link-parser/htmllinkextractor"
	"io/ioutil"
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

func readExampleHtmlFromDisk(path string) string {
	exampleOneHtml, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(exampleOneHtml)

}

func runStandardTestCase(pathToExampleHtmlFile string, expectedParsedLinks []htmllinkextractor.Link) {
	exampleHtmlStringified := readExampleHtmlFromDisk(pathToExampleHtmlFile)
	Expect(htmllinkextractor.ExtractLinks(exampleHtmlStringified)).To(Equal(expectedParsedLinks))
}

var _ = Describe("htmllinkextractor Integration Test Suite", func() {
	It("Extracts all tags from html example 1", func() {
		runStandardTestCase("./ex1.html", []htmllinkextractor.Link{
			{
				Href: "/other-page",
				Text: "A link to another page",
			},
		})
	})

	It("Extracts all tags from html example 2", func() {
		runStandardTestCase("./ex2.html", []htmllinkextractor.Link{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			},
		})
	})

	It("Extracts all tags from html example 3", func() {
		runStandardTestCase("./ex3.html", []htmllinkextractor.Link{
			{
				Href: "#",
				Text: "Login",
			},
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
		})
	})

	It("Extracts all tags from html example 4", func() {
		runStandardTestCase("./ex4.html", []htmllinkextractor.Link{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		})
	})
})
