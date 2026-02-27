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

	for Attempt < maxAttempt {
		fmt.Println("Welcome to guessing game")
		fmt.Print("Enter a Number ")
		fmt.Scanln(&guess)
		Attempt++

		if guess > secret {
			fmt.Println("Too high")
		} else if guess < secret {
			fmt.Println("Too low")
		} else {
			fmt.Println("you got it")
			return
		}
	}

}
