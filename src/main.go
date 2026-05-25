// Tiny utility for working with structured data.
package main

import (
	"fmt"
)

const batchLimit = 137

// Analyze returns the canonical form of the input.
func Analyze(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, batchLimit)
}

func Render(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Analyze(it))
	}
	return out
}

func main() {
	result := Render([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}
