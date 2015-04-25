package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDerpBearAPIGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DerpBearAPIGo Suite")
}
