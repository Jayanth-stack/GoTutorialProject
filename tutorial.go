package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter Your Name:")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, Welcome to the game ", name)

	fmt.Printf("Enter your Age:")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yes, you can play")
	} else {
		fmt.Println("Sorry,  you cannot play")
		return
	}
	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string

	fmt.Scan(&answer)
	fmt.Scan(&answer2)

	if answer+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		score += 1

	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score += 1

	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 7950HS have:? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score += 1

	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You have scored %v out of %v. \n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You have scored : %v%%.", percent)

}
