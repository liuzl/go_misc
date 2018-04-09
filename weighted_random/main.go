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

func WeightedRandomS3(weights []float32) int {
	n := len(weights)
	if n == 0 {
		return 0
	}
	cdf := make([]float32, n)
	var sum float32 = 0.0
	for i, w := range weights {
		if i > 0 {
			cdf[i] = cdf[i-1] + w
		} else {
			cdf[i] = w
		}
		sum += w
	}
	r := rand.Float32() * sum
	var l, h int = 0, n - 1
	for l <= h {
		m := l + (h-l)/2
		if r <= cdf[m] {
			if m == 0 || (m > 0 && r > cdf[m-1]) {
				return m
			}
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(WeightedRandomS3([]float32{0.1, 0.3, 0.2}))
	}
}
