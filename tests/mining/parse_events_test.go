package mining_test

import (
	"io/ioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/yggie/github-data-challenge-2014/mining"
	"github.com/yggie/github-data-challenge-2014/models"
)

var _ = Describe("ParseEvents", func() {
	Context("with a valid json fixture", func() {
		var results *mining.EventsResult

		BeforeEach(func() {
			json, err := ioutil.ReadFile("../fixtures/github_events.json")
			if err != nil {
				panic(err)
			}
			results = mining.ParseEvents(json)
		})

		Context("when extracting PushEvents", func() {
			var pushEvent *models.PushEvent

			BeforeEach(func() {
				pushEvent = results.PushEvents[0]
			})

			It("should parse exactly 15 Events out of 30", func() {
				Expect(results.PushEvents).To(HaveLen(15))
			})

			It("should parse the event ids", func() {
				Expect(pushEvent.Id).ToNot(Equal(0))
			})

			It("should correctly parse the event types", func() {
				Expect(pushEvent.EventType).To(Equal("PushEvent"))
			})

			It("should parse the user information", func() {
				Expect(pushEvent.User).ToNot(BeNil())
			})

			It("should parse the repository information", func() {
				Expect(pushEvent.Repository).ToNot(BeNil())
			})
		})
	})
})
