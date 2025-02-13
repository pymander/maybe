package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var chance int
	var verbose bool

	flag.IntVar(&chance, "chance", 50, "Percent chance from 1 to 99")
	flag.BoolVar(&verbose, "verbose", false, "Be noisy")
	flag.Parse()

	if chance < 1 {
		chance = 1
	}

	if chance > 99 {
		chance = 99
	}

	num := rand.Intn(100) // random int 0-99

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
