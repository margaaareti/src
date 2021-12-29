package main

import (
	"HeadFirst/datafile"
	"fmt"
	"log"
)

func main() {

	lines, err := datafile.GetString("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int) //объявляет карту, у которой ключами являются имена кандидатов, а знчениями- счетчики голосов
	for _, line := range lines {
		counts[line]++ // увеличивает счетчик голосов для текущего кандидта
	}

	for name, count := range counts {
		fmt.Printf("%v: %v\n", name, count)
	}

}
