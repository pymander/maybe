package main

import (
	"math"
	"time"
)

// Calculate the phase of the moon according to https://en.wikipedia.org/wiki/Lunar_phase#Calculating_phase
const meanSynodicMonth = 29.53059

// moonPhase calculates the approximate phase of the moon at time t
func moonPhase(t time.Time) float64 {
	// The known New Moon needs to be recent!
	knownNewMoon := time.Date(2025, 1, 28, 12, 35, 0, 0, time.UTC)

	days := t.Sub(knownNewMoon).Hours() / 24
	moonAge := math.Mod(days, meanSynodicMonth)

	return getMoonMultiplier(moonAge)
}

// Transform the Moon's age into a weighted multiplier.
func getMoonMultiplier(age float64) float64 {
    // Convert age to radians 
    radians := age * 2 * math.Pi
    
    // Cosine wave oscillates between -1 and 1
    // - Full moon (age ~14.765) should be lowest (0.8)
    // - New moon (age 0 or meanSynodicMonth) should be highest (1.2)
    base := math.Cos(radians)
    
    // Transform -1...1 into 0.8...1.2
    multiplier := 1.0 + (base * 0.2)
    
    return multiplier
}
