package main

import (
	"fmt"
	"math/rand"
)

func NumberGenerator() int {
	random := rand.Intn(100)
	return random
}

func main() {
	var num int
	fmt.Printf("Choose your number between 0 to 100")
	fmt.Scan(&num)

	if NumberGenerator() == num {
		fmt.Printf("Congratulations")
	} else {
		fmt.Printf("Try Again")
	}

}
