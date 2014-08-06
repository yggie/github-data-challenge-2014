package models_test

import (
	"io/ioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/yggie/github-data-challenge-2014/models"
)

var _ = Describe("ParseEvents", func() {
	Context("with a valid json fixture", func() {
		var events []models.Event

		BeforeEach(func() {
			json, err := ioutil.ReadFile("../fixtures/github_events.json")
			if err != nil {
				panic(err)
			}
			events = models.ParseEvents(json)
		})

		AfterEach(func() {
			models.ClearEvents()
		})

		It("should parse exactly 15 Events out of 30", func() {
			Expect(events).To(HaveLen(15))
		})

		It("should parse the event ids", func() {
			Expect(events[0].Id()).ToNot(Equal(0))
		})

		It("should parse the event types", func() {
			Expect(events[0].Type()).ToNot(BeEmpty())
		})

		It("should parse the user information", func() {
			Expect(events[0].User().Login()).ToNot(BeEmpty())
		})

		It("should parse the repository information", func() {
			Expect(events[0].Repository().Name()).ToNot(BeEmpty())
		})

		It("should be persistable", func() {
			before := models.CountEvents()
			err := models.Persist(events[0])
			if err != nil {
				panic(err)
			}
			after := models.CountEvents()
			Expect(after - before).To(Equal(1))
		})
	})
})
