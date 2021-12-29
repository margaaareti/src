package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true // Возвращает true, если строка найдена в сегменте
		}
	}
	return false // или false, если строка не найдена
}

type Refrigerator []string // тип основан на сегменте строке, в котором хранятся названия продуктов из холодильника

func (r Refrigerator) Open() { // моделирует открытие двери
	fmt.Println("Opening Refrigerator")
}

func (r Refrigerator) Close() { // моделирует закрытие двери
	fmt.Println("Close Refrigerator")
}

func (r Refrigerator) FindFood(food string) error {
	defer r.Close()
	r.Open()
	if find(food, r) { //если содержит данный продукт
		fmt.Println("Found", food) //выводится сообщение о том, что продукт найден
	} else {
		return fmt.Errorf("%s not found", food) // в противном случае выводится сообщение об ошибке
	}
	return nil
}

func main() {

	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}

}
