package neo_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNeo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Neo Test Suite")
}
