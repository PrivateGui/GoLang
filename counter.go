package main

import (
	"bufio"
	"fmt"
	"os"
)

func errors(err error) bool {
	if err != nil {
		fmt.Println("Enter a valid number.")
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("---Simple Counter // Made by ZoneRCM---")
	fmt.Println()
	var min int
	var max int
	fmt.Print("Enter your first number: ")
	_, err := fmt.Scan(&min)
	if errors(err) == true {
		return
	}

	minorg := min

	fmt.Println()
	fmt.Print("Enter your second number: ")
	_, err2 := fmt.Scan(&max)
	if errors(err2) == true {
		return
	}
	maxorg := max
	fmt.Println()

	for ; min <= max; min++ {
		fmt.Println(min)
	}
	fmt.Println()
	fmt.Printf("Counted from %d to %d.", minorg, maxorg)
	fmt.Println()
	fmt.Println("\nPress Enter to exit...")

	reader := bufio.NewReader(os.Stdin)
	// This one consumes the leftover '\n' from the second fmt.Scan
	reader.ReadString('\n')
	// This one actually waits for the user to press Enter to exit
	reader.ReadString('\n')
}
