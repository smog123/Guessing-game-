package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guess int
	secret := rand.Intn(100) + 1
	Attempt := 0
	maxAttempt := 5
	Guesses := []int{}

	fmt.Println("Welcome to guessing game")

	for Attempt < maxAttempt {
		//fmt.Println("Welcome to guessing game")
		fmt.Print("Enter a Number ")
		fmt.Scanln(&guess)

		Attempt++
		Guesses = append(Guesses, guess)
		if guess > secret {
			fmt.Println("Too high")
		} else if guess < secret {
			fmt.Println("Too low")
		} else {
			fmt.Println("you got it")
			fmt.Println("Your guesses", Guesses)
			fmt.Println("your attempts ", Attempt)
			return
		}
	}
	fmt.Println("out of attempts")
	fmt.Println("Guesses is ", Guesses)
	fmt.Println("the secret number is", secret)

}
