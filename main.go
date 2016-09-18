package main

import (
	"fmt"
	"os"
	"strconv"
)

var SmolovJr = Program{
	Days: []Day{
		Day{
			Name:       "Monday",
			Sets:       6,
			Reps:       6,
			Percentage: 0.7,
		},
		Day{
			Name:       "Wednesday",
			Sets:       7,
			Reps:       5,
			Percentage: 0.75,
		},
		Day{
			Name:       "Friday",
			Sets:       8,
			Reps:       4,
			Percentage: 0.8,
		},
		Day{
			Name:       "Saturday",
			Sets:       10,
			Reps:       3,
			Percentage: 0.85,
		},
	},
}

func main() {
	var liftName string
	var weeksStr string
	var incrementStr string
	var maxStr string

	fmt.Print("Lift name: ")
	fmt.Scanln(&liftName)

	fmt.Print("Current 1RM: ")
	fmt.Scanln(&maxStr)
	max, err := strconv.Atoi(maxStr)
	if err != nil {
		fmt.Errorf("unable to parse number out of: %v", maxStr)
		os.Exit(1)
	}

	fmt.Print("Weeks to run program: ")
	fmt.Scanln(&weeksStr)
	weeks, err := strconv.Atoi(weeksStr)
	if err != nil {
		fmt.Errorf("unable to parse number out of: %v", weeksStr)
		os.Exit(1)
	}

	fmt.Print("Increment each week: ")
	fmt.Scanln(&incrementStr)
	increment, err := strconv.Atoi(incrementStr)
	if err != nil {
		fmt.Errorf("unable to parse number out of: %v", incrementStr)
		os.Exit(1)
	}

	SmolovJr.Print(liftName, max, weeks, increment)
}
