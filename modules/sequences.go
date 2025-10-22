package modules

import (
	"fmt"

	"playground/collections"

	"github.com/spf13/cobra"
)

func numbers() collections.Enumerable[int] {
	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

func NewSequencesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sequences",
		Short: "Run sequences demonstration",
		Run:   runSequences,
	}

	return cmd
}

func runSequences(_ *cobra.Command, _ []string) {
	sequence := collections.Map(
		numbers().Take(10),
		func(n int) int {
			return n * n
		},
	).Filter(func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(collections.Reduce(
		sequence,
		0,
		func(acc, n int) int {
			return acc + n
		},
	))

	fmt.Println()

	nestedSequences := func() collections.Enumerable[collections.Enumerable[int]] {
		return func(yield func(collections.Enumerable[int]) bool) {
			if !yield(numbers().Take(10)) {
				return
			}

			if !yield(numbers().Take(10)) {
				return
			}
		}
	}()

	collections.FlatMap(nestedSequences, func(i int) int {
		return i * i
	}).ForEach(func(n int) {
		fmt.Println(n)
	})
}
