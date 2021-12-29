package main

//стр 462
import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello , web") //строка преобразуется в сегмент байтов
	_, err := writer.Write(message)  // строка добавляется в ответ
	if err != nil {
		log.Fatal()
	}
}

func main() {

	http.HandleFunc("/hello", viewHandler)
	// если был получен запрос для URL- адреса, завершающегося суффиксом </hello>,
	// для генерирования ответа вызывается функция viewHandler
	err := http.ListenAndServe("localhost:8080", nil) // Обрабатывает запрос браузера и отвечает на него
	log.Fatal(err)
}

// http://localhost:8080/hello вводим в браузере для запуска
