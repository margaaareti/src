package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number") // другие числа генерироваться не должны, но если все же произойдет- это повод для паники
	}
}
func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
