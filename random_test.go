package main

import (
	"testing"
	"time"
)

func TestRNG(t *testing.T) {
	rng := &Generator{}

	// Basic functionality
	t.Run("generates different numbers", func(t *testing.T) {
		n1 := rng.Intn(10000)
		n2 := rng.Intn(10000)
		if n1 == n2 {
			t.Error("Generated same number twice")
		}
	})

	// Distribution test (chi-square)
	t.Run("distribution", func(t *testing.T) {
		buckets := make([]int, 10)
		for i := 0; i < 10000; i++ {
			n := rng.Intn(10000) % 10
			buckets[n]++
		}
		// Expect roughly 1000 per bucket
		for i, count := range buckets {
			if count < 900 || count > 1100 {
				t.Errorf("Bucket %d: %d items (expected ~1000)", i, count)
			}
		}
	})

	// Ensuring range
	t.Run("range check", func(t *testing.T) {
		maxValue := 10000
		for i := 0; i < 1000; i++ {
			n := rng.Intn(10000)
			if n < 0 || n >= maxValue {
				t.Errorf("Number %v outside expected range", n)
			}
		}
	})

	// Benchmark
	t.Run("performance", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < 10000; i++ {
			rng.Intn(10000)
		}
		duration := time.Since(start)
		if duration > time.Second {
			t.Error("RNG too slow")
		}
	})
}
