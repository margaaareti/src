// стр. 430
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) { // передаем канал responSize для передачи размеров страниц
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal()
	}

	channel <- Page{URL: url, Size: len(body)} // отправляет структуру Page с текущим URL- адресом и размером страницы

}

func main() {

	pages := make(chan Page) // page- тип передаваемого канала
	urls := []string{"https://www.tegos.ru/",
		"https://golang.org/", "https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages                             //получеам page
		fmt.Printf("%s: %d\n", page.URL, page.Size) // URL- адрес выводится вместе с размером
	}

}
