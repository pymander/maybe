package main

import (
	"crypto/rand"
	"encoding/binary"
)

type Generator struct {
	buf [8]byte
}

// Generate a random number in the range [0,n). Uses 63 bits to avoid negative numbers.
func (g *Generator) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid range")
	}

	// This helps us avoid modulo bias (https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#Modulo_bias)
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := g.Int63()
	for v > max {
		v = g.Int63()
	}

	// Keep our number within range using modulo.
	return v % n
}

func (g *Generator) Int63() int64 {
	// crypto/rand does all of our /dev/urandom stuff for us.
	rand.Read(g.buf[:])
	return int64(binary.LittleEndian.Uint64(g.buf[:]) & (1<<63 - 1))
}

func (g *Generator) Float64() float64 {
	return float64(g.Int63()) / (1 << 63)
}

// Generate a random integer in the range [0,n).
func (g *Generator) Intn(n int) int {
	if n <= 0 {
		panic("invalid range")
	}
	return int(g.Int63n(int64(n)))
}
