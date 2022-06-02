package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFunc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Basic Tests")
}
