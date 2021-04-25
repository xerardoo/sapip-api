package models_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	m "github.com/xerardoo/sapip/models"
	_ "github.com/xerardoo/sapip/tester"

)

func TestFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Test Suite")
}

var _ = Describe("File", func() {

	file, err := os.Open("../men.png")
	stat, _ := file.Stat()
	It("cant be read", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(stat.Size()).Should(BeNumerically(">", 0))
		Expect(file.Name()).Should(Equal("../men.png"))
	})
	defer file.Close()

	err = m.UploadFileS3("photos", "men.png", file)
	It("cant be uploaded", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	path := m.GetFilePathS3("men.png")
	It("cant be uploaded", func() {
		Expect(path).Should(Equal("https://sapip.s3-us-west-2.amazonaws.com/men.png"))
	})
})
