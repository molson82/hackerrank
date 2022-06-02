package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ransom Note HackerRank Tests", func() {
	Describe("true mock test", func() {
		It("should always pass", func() {
			Expect(true).To(Equal(true))
		})
	})
	Describe("actual tests", func() {
		It("should return yes", func() {
			mag := []string{"give", "me", "one", "grand", "today", "night"}
			note := []string{"give", "one", "grand", "today"}
			Expect(checkMagazine(mag, note)).To(Equal("Yes"))
		})
		It("should return no", func() {
			mag := []string{"two", "times", "three", "is", "not", "four"}
			note := []string{"two", "times", "two", "is", "four"}
			Expect(checkMagazine(mag, note)).To(Equal("No"))
		})
		It("should return no", func() {
			mag := []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}
			note := []string{"ive", "got", "some", "coconuts"}
			Expect(checkMagazine(mag, note)).To(Equal("No"))
		})
	})
})
