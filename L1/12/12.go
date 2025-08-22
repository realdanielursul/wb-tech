package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, w := range strs {
		set[w] = struct{}{}
	}

	for k := range set {
		fmt.Println(k)
	}
}
