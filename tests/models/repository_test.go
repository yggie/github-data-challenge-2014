package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/yggie/github-data-challenge-2014/models"
)

var _ = Describe("Repository", func() {
	var repository models.Repository

	BeforeEach(func() {
		repository = models.Repository{
			Id:   123456,
			Name: "my-repository",
			Url:  "http://example.com",
			Languages: models.Languages{
				models.CPP:   50,
				models.JAVA:  25,
				models.SHELL: 25,
			},
		}
	})

	Context(".LanguageDistribution()", func() {
		var distribution models.LanguageDistribution

		BeforeEach(func() {
			distribution = repository.LanguageDistribution()
		})

		It("should compute the correct value for the language distribution", func() {
			Expect(distribution[models.CPP]).To(BeNumerically("==", 50.0))
			Expect(distribution[models.JAVA]).To(BeNumerically("==", 25.0))
			Expect(distribution[models.SHELL]).To(BeNumerically("==", 25.0))
		})
	})
})
