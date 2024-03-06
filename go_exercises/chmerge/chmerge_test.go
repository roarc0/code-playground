package chmerge

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/require"
)

func consume(c <-chan string) []string {
	var msgs []string
	for m := range c {
		msgs = append(msgs, m)
	}
	return msgs
}

var _ = Describe("Fixed Tests", func() {
	It("should merge 0 channels", func() {
		c := Merge[string]()

		actual := consume(c)
		Expect(actual).To(BeEmpty())
	})

	It("should merge 1 channel", func() {
		a := make(chan string, 3)
		a <- "foo"
		a <- "bar"
		a <- "baz"
		close(a)

		merged := Merge(a)

		actual := consume(merged)
		Expect(actual).To(ConsistOf([]string{"foo", "bar", "baz"}))
	})

	It("should merge 2 channels", func() {
		a := make(chan string, 3)
		a <- "a-foo"
		a <- "a-bar"
		a <- "a-baz"
		close(a)

		b := make(chan string, 3)
		b <- "b-foo"
		b <- "b-bar"
		b <- "b-baz"
		close(b)

		merged := Merge(a, b)

		actual := consume(merged)
		Expect(actual).To(ConsistOf([]string{
			"a-foo", "a-bar", "a-baz",
			"b-foo", "b-bar", "b-baz",
		}))
	})

	It("should merge 3 channels", func() {
		a := make(chan string, 3)
		a <- "a-foo"
		a <- "a-bar"
		a <- "a-baz"
		close(a)

		b := make(chan string, 3)
		b <- "b-foo"
		b <- "b-bar"
		b <- "b-baz"
		close(b)

		c := make(chan string, 3)
		c <- "c-foo"
		c <- "c-bar"
		c <- "c-baz"
		close(c)

		merged := Merge(a, b, c)

		actual := consume(merged)
		Expect(actual).To(ConsistOf([]string{
			"a-foo", "a-bar", "a-baz",
			"b-foo", "b-bar", "b-baz",
			"c-foo", "c-bar", "c-baz",
		}))
	})

	It("should consume from first channel while second one stays empty", func() {
		a := make(chan string)
		b := make(chan string)

		go func() {
			a <- "first"
			a <- "second"
			a <- "third"
			close(a)
			close(b)
		}()

		c := Merge(a, b)

		Expect(consume(c)).To(ConsistOf([]string{"first", "second", "third"}))
	})

	It("should consume from second channel while first one stays empty", func() {
		a := make(chan string)
		b := make(chan string)

		go func() {
			b <- "first"
			b <- "second"
			b <- "third"
			close(a)
			close(b)
		}()

		c := Merge(a, b)

		Expect(consume(c)).To(ConsistOf([]string{"first", "second", "third"}))
	})
})

func TestMerge(t *testing.T) {
	a := make(chan string, 3)
	a <- "a-foo"
	a <- "a-bar"
	a <- "a-baz"
	close(a)

	b := make(chan string, 3)
	b <- "b-foo"
	b <- "b-bar"
	b <- "b-baz"
	close(b)

	c := make(chan string, 3)
	c <- "c-foo"
	c <- "c-bar"
	c <- "c-baz"
	close(c)

	merged := Merge(a, b, c)

	actual := consume(merged)
	require.ElementsMatch(t, actual,
		[]string{
			"a-foo", "a-bar", "a-baz",
			"b-foo", "b-bar", "b-baz",
			"c-foo", "c-bar", "c-baz",
		})
}
