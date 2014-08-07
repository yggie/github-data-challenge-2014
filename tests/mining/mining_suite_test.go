package mining_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMining(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mining Test Suite")
}
