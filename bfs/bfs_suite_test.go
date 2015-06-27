package bfs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBfs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bfs Suite")
}
