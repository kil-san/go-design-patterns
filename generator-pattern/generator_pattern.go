package generator

import (
	"fmt"
)

func iterate(words []string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, word := range words {
			out <- word
		}
	}()

	return out
}

// GeneratorDemo test generator pattern
func GeneratorDemo() {
	fmt.Println("Generator Demo.........")
	words := []string{
		"apple",
		"ball",
		"cat",
		"dog",
		"egg",
	}

	for word := range iterate(words) {
		fmt.Println(word)
	}
}
