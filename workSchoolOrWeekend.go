// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program checks if the user should go to work, school or if it's the weekend

package main

import "fmt"

func main() {
	// This function checks if the user should go to work, school or if it's the weekend
	var userDay string
	var userAge int

	// input
	fmt.Println("This program checks if the user should go to work, school or if it's the weekend")
	fmt.Println()
	fmt.Print("Enter the day: ")
	fmt.Scanln(&userDay)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&userAge)
	fmt.Println()

	// process
	if !(userDay == "saturday" || userDay == "sunday") && (userAge >= 18 ) {
		// output
		fmt.Println("Time for work")
	} else if !(userDay == "saturday" || userDay == "sunday") && (userAge < 18 ) {
		// output
		fmt.Println("Time for school")
	} else {
		// output
		fmt.Println("Time to relax")
	}
}
