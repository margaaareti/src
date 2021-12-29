package main

import (
	"fmt"
	"log"
)

type error interface {
	Error() string
}

type OverheatError float64 // определяет тип с базовым типом float64

func (o OverheatError) Error() string { // поддерживает интерфейс Error
	return fmt.Sprintf("Overheating be %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess) // возвращает OverheatError с превышением безопасной температуры
	}
	return nil
}

func main() {

	var err error = checkTemperature(121.22, 100.0)
	if err != nil {
		log.Fatal(err)
	}

}
