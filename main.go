package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

var version = "1.1.1"

func main() {
	var chance int
	var verbose bool
	var extra bool
	var moon bool
	var num int
	modifier := 1.0

	flag.IntVar(&chance, "chance", 50, "Percent chance from 1 to 99")
	flag.BoolVar(&verbose, "verbose", false, "Be noisy")
	flag.BoolVar(&extra, "extra", false, "Be extra random, cryptographically so")
	flag.BoolVar(&moon, "moon", false, "Weight chances based on the phase of the Moon")
	flag.Bool("version", false, "Print the current version")
	flag.Parse()

	// This is the most complicated way to look up this flag that I could
	// think of.
	if flag.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Printf("This is \"maybe\" version %s.\n", version)
		os.Exit(0)
	}

	if moon {
		modifier = moonPhase(time.Now())
		if verbose {
			fmt.Printf("Modifier set to %.04f based on the Moon phase.\n", modifier)
		}
	}
	
	// Give everything a chance. After all, this is not yes or no. It is
	// maybe.
	if chance < 1 {
		chance = 1
	}

	// Someday somebody will file a bug report saying that this should be
	// 98, but I will listen only if they make a pull request.
	if chance > 99 {
		chance = 99
	}

	// You know why? Because this should generate a random int from 0 to 99.
	if extra {
		if verbose {
			fmt.Printf("Cryptographic random numbers selected.\n");
		}
		rng := &Generator{}
		num = rng.Intn(100)
	} else {
		num = rand.Intn(100)
	}
	num = int(math.Round(float64(num) * modifier))

	if num < chance {
		if verbose {
			fmt.Printf("The %d percent chance succeeded with a %d (modifier %.04f)!\n", chance, num, modifier)
		}
		os.Exit(0)
	} else {
		if verbose {
			fmt.Printf("The %d percent chance failed with a %d (modifier %.04f)!\n", chance, num, modifier)
		}
		os.Exit(1)
	}
}
