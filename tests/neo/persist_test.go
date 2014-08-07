package neo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/yggie/github-data-challenge-2014/models"
	"github.com/yggie/github-data-challenge-2014/neo"
)

var _ = Describe("Persist", func() {
	neo.Clear(neo.ALL)
	var err error

	AfterEach(func() {
		neo.Clear(neo.ALL)
	})

	Context("PersistPushEvent", func() {
		var pushEvent models.PushEvent

		BeforeEach(func() {
			pushEvent = models.PushEvent{
				Size:         3,
				DistinctSize: 2,
				PushId:       123123123,
				Commits:      []*models.Commit{},
				Event: &models.Event{
					Id:        1,
					EventType: "PushEvent",
					CreatedAt: "Today",
					User: &models.User{
						Id:         1,
						Login:      "samuex",
						GravatarId: "abcdef",
						AvatarUrl:  "http://example.com",
					},
					Repository: &models.Repository{
						Id:   1,
						Name: "my-repo",
						Url:  "http://github.com/randomguy/my-repo",
					},
				},
			}

			err = neo.PersistPushEvent(&pushEvent)
		})

		It("should not have any errors", func() {
			Expect(err).To(BeNil())
		})

		It("should persist the event object", func() {
			Expect(neo.Count(neo.EVENTS)).To(Equal(1))
		})

		It("should persist the repository object", func() {
			Expect(neo.Count(neo.REPOSITORIES)).To(Equal(1))
		})
	})
})
