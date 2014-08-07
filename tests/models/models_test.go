package models_test

import (
	"io/ioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/yggie/github-data-challenge-2014/mining"
	"github.com/yggie/github-data-challenge-2014/models"
	"github.com/yggie/github-data-challenge-2014/neo"
)

var _ = Describe("ParseEvents", func() {
	Context("with a valid json fixture", func() {
		var (
			results    *mining.EventsResult
			pushEvents []*models.PushEvent
		)

		BeforeEach(func() {
			json, err := ioutil.ReadFile("../fixtures/github_events.json")
			if err != nil {
				panic(err)
			}
			results = mining.ParseEvents(json)
			pushEvents = results.PushEvents
		})

		AfterEach(func() {
			neo.ClearEvents()
		})

		It("should parse exactly 15 Events out of 30", func() {
			Expect(pushEvents).To(HaveLen(15))
		})

		It("should parse the event ids", func() {
			Expect(pushEvents[0].Id).ToNot(Equal(0))
		})

		It("should correctly parse the event types", func() {
			Expect(pushEvents[0].EventType).To(Equal("PushEvent"))
		})

		It("should parse the user information", func() {
			Expect(pushEvents[0].User.Login).ToNot(BeEmpty())
		})

		It("should parse the repository information", func() {
			Expect(pushEvents[0].Repository.Name).ToNot(BeEmpty())
		})

		It("should be persistable", func() {
			before := neo.CountEvents()
			err := neo.PersistPushEvent(pushEvents[0])
			if err != nil {
				panic(err)
			}
			after := neo.CountEvents()
			Expect(after - before).To(Equal(1))
		})
	})
})
