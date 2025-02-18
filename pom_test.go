package main

import (
    "testing"
    "time"
)

// Pulled most of these Moon dates from https://www.fullmoonology.com/moon-phases-2026/
func TestMoonPhase(t *testing.T) {
    tests := []struct {
        name     string
        date     time.Time
        expected float64
    }{
        {
            "Known New Moon 1",
            time.Date(2025, 2, 28, 0, 44, 0, 0, time.UTC),
            1.2,
        },
        {
            "Known New Moon 2",
            time.Date(2026, 1, 18, 19, 52, 0, 0, time.UTC),
            1.2,
        },
        {
            "Known New Moon 3",
            time.Date(2026, 2, 17, 12, 1, 0, 0, time.UTC),
            1.2,
        },
        {
            "Known New Moon 4",
            time.Date(2028, 8, 20, 10, 44, 0, 0, time.UTC),
            1.2,
        },
        {
            "Known Full Moon 1",
            time.Date(2025, 5, 12, 16, 55, 0, 0, time.UTC),
            0.8,
        },
        {
            "Known Full Moon 2",
            time.Date(2026, 11, 24, 14, 53, 0, 0, time.UTC),
            0.8,
        },
        {
            "Known Full Moon 3",
            time.Date(2028, 8, 5, 8, 11, 0, 0, time.UTC),
            0.8,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := moonPhase(tt.date)
            if diff := got - tt.expected; diff > 0.05 || diff < -0.05 {
                t.Errorf("moonPhase() = %v, want %v", got, tt.expected)
            }
        })
    }
}
