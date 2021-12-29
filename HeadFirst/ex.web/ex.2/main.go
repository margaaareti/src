package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message)) //строка преобразуется в сегмент байтов, как и прежде, записывается в ответ
	if err != nil {
		log.Fatal(err)
	}

}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello,web!") // строка записывается в ответ

}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut,web!")

}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste,web!")
}

//вызывает ответ в соответствии с URL . . .
func main() {

	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
