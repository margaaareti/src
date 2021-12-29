package main

import "fmt"

func main() {
	month := []string{1: "Январь", 2: "Февраль", 3: "Март", 4: "Апрель", 5: "Май", 6: "Июнь", 7: "Июль", 8: "Август"}

	Q2 := month[1:5]
	winter := month[1:3]

	for _, q := range Q2 {
		for _, z := range winter {
			if q == z {
				fmt.Printf("%q находится в обеих срезах\n", q)
			}
		}
	}
}
