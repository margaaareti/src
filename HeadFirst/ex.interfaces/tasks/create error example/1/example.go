// страница 374

package main

import "fmt"

type error interface {
	Error() string
}

type ComedyError string // определяем тип с базовым типом string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {

	var err error
	err = ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)

}
