package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

var version = "1.0.0"

func main() {
	var chance int
	var verbose bool

	flag.IntVar(&chance, "chance", 50, "Percent chance from 1 to 99")
	flag.BoolVar(&verbose, "verbose", false, "Be noisy")
	flag.Bool("version", false, "Print the current version")
	flag.Parse()

	// This is the most complicated way to look up this flag that I could
	// think of.
	if flag.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Printf("This is \"maybe\" version %s.\n", version)
		os.Exit(0)
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
	num := rand.Intn(100)

	if num < chance {
		if verbose {
			fmt.Printf("The %d percent chance succeeded with a %d!\n", chance, num)
		}
		os.Exit(0)
	} else {
		if verbose {
			fmt.Printf("The %d percent chance failed with a %d!\n", chance, num)
		}
		os.Exit(1)
	}
}
