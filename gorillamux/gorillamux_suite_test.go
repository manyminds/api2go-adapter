package gorillamux_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGorillamux(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gorillamux Suite")
}
