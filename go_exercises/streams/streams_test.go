package streams

import (
	"context"
	"fmt"
	"testing"
)

func TestStreams(t *testing.T) {
	data := []int{0, 2, 4, 6, 8}

	ctx := context.Background()
	result := Collect(ctx, Transform(ctx, Dividend(100), Filter(ctx, NonZero, Stream(ctx, data))))
	fmt.Printf("%v\n", result)
}

func NonZero(n int) bool {
	return n != 0
}

func Dividend(divisor int) func(int) int {
	return func(n int) int {
		return divisor / n
	}
}
