package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test10x(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "10x Suite")
}
