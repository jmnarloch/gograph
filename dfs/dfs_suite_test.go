package dfs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDfs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dfs Suite")
}
