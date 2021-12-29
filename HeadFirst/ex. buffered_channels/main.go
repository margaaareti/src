package main

import (
	"fmt"
	"time"
)

func sendLetters(channel chan string) { //получает канал в параметре
	time.Sleep(1 * time.Second)
	channel <- "a"
	time.Sleep(1 * time.Second)
	channel <- "b"
	time.Sleep(1 * time.Second)
	channel <- "c"
	time.Sleep(1 * time.Second)
	channel <- "d"

}

func main() {
	fmt.Println(time.Now())         // Выводит время запуска программы
	channel := make(chan string, 3) //создает небуферизованный канал, как это делалось ранее
	go sendLetters(channel)         // запускает sendLetters в новой горутине
	time.Sleep(5 * time.Second)     // приостанавливает горутину main на 5 секунд
	//Получает и выводит четыре значения с текущим временем
	fmt.Println(<-channel, time.Now().Second())
	fmt.Println(<-channel, time.Now().Second())
	fmt.Println(<-channel, time.Now().Second())
	fmt.Println(<-channel, time.Now().Second())
	fmt.Println(time.Now())
}
