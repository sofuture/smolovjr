package main

import (
	"fmt"
	"os"
	"strconv"
)

func cli() error {
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
		return err
	}

	fmt.Print("Weeks to run program: ")
	fmt.Scanln(&weeksStr)
	weeks, err := strconv.Atoi(weeksStr)
	if err != nil {
		fmt.Errorf("unable to parse number out of: %v", weeksStr)
		return err
	}

	fmt.Print("Increment each week: ")
	fmt.Scanln(&incrementStr)
	increment, err := strconv.Atoi(incrementStr)
	if err != nil {
		fmt.Errorf("unable to parse number out of: %v", incrementStr)
		return err
	}

	SmolovJr.Print(os.Stdout, liftName, max, weeks, increment)

	return nil
}
