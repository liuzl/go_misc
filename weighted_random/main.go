package main

import (
	"fmt"
	"math/rand"
)

func WeightedRandomS1(weights []float32) int {
	if len(weights) == 0 {
		return 0
	}
	var choices []int
	for i, w := range weights {
		wi := int(w * 10)
		for j := 0; j < wi; j++ {
			choices = append(choices, i)
		}
	}
	return choices[rand.Int()%len(choices)]
}

func WeightedRandomS2(weights []float32) int {
	if len(weights) == 0 {
		return 0
	}
	var sum float32 = 0.0
	for _, w := range weights {
		sum += w
	}
	r := rand.Float32() * sum
	for i, w := range weights {
		r -= w
		if r < 0 {
			return i
		}
	}
	return len(weights) - 1
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(WeightedRandomS2([]float32{0.1, 0.3, 0.6}))
	}
}
