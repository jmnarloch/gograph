package digraph_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDigraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Digraph Suite")
}
