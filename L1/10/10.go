package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.4, -6}
	groups := make(map[int][]float64)

	for _, temp := range temps {
		key := int(math.Floor(temp/10.0)) * 10
		groups[key] = append(groups[key], temp)
	}

	for group := range groups {
		fmt.Println(group, groups[group])
	}
}
