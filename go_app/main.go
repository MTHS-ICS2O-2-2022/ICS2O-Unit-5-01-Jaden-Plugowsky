// Created by: Jaden Plugowsky
// Created on: May 2023
//
// This program finds a random number, then it asks the user to guess it

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This function finds a random number, then it asks the user to guess it
		rand.Seed(time.Now().UnixNano())
    min := 1
    max := 6
	var randomNumber int
	var userGuess int

	// Input
	randomNumber = (rand.Intn(max - min) + min)
	fmt.Println("\nThis program finds a random number, then asking the user to guess the random number.")
	fmt.Println()
	fmt.Print("Enter your guess from 1-6: ")
	fmt.Scanln(&userGuess)

	// Process
	if userGuess == randomNumber {
		fmt.Println("You have guessed the correct number, which was", randomNumber, "!")
	}
	if userGuess != randomNumber {
		fmt.Println("You have guessed the number incorrectly, which was", randomNumber, ".")
	}

	// Output
	fmt.Println("\nDone.")
}

/*
	// Process
	volume = 4.0 / 3.0 * math.Pi * math.Pow(radius, 3)

	// Output
	fmt.Printf("\nThe volume of the Sphere is: %.2f cm³.", volume)

	fmt.Println("\n\nDone.")
*/