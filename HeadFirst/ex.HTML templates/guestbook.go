//стр.481

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

//Функции обработки запроса и проверки ошибок

//код сообщения об ошибках перемещается в эту функцию
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

//функциям-обработчикам передается значение ResponseWriter......а также указатель на значение Request
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signature.txt")     // добавляем вызов GetStrings
	html, err := template.ParseFiles("view.html") // содержимое view.html используется для создания нового значения Template
	check(err)                                    // сообщает об ошибках
	guestbook := Guestbook{
		SignatureCount: len(signatures), // в данное поле сохраняется длина сегмента записей
		Signatures:     signatures,      // а тут сохраняется сам сегмент записей
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func getStrings(fileName string) []string {
	var lines []string             // Каждая строка файла присоединяется к сегменту в виде отдельного элемента
	file, err := os.Open(fileName) // Открывает файл
	if os.IsNotExist(err) {        // Если возвращается ошибка,указыв. на то, что файл не существует ...
		return nil // вернуть nil вместо сегмента строк
	}
	check(err)                        // Для любой другой ошибки сообщить о ней и завершить работу
	defer file.Close()                // После выхода из функции позаботится о том, чтобы файл был закрыт
	scanner := bufio.NewScanner(file) //Создается сканер для содержимого файла
	for scanner.Scan() {              //Для каждой строки файла...
		lines = append(lines, scanner.Text()) // ...ее текст присоединяется к сегменту.
	}
	check(scanner.Err()) // Сообщает о любых ошибках, обнаруженных во время сканирования
	return lines
}

func newHandler(writer http.ResponseWriter, request *http.Request) { // Добавляем другую функцию обработчик с такими же параметрами, как у viewHandler
	html, err := template.ParseFiles("new.html") // загружаем new.html, как текст шаблона
	check(err)
	_ = html.Execute(writer, nil) // записываем шаблон в ответ
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")                           // получаем значение поля "signature" формы
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE                    // параметры открытия файла
	file, err := os.OpenFile("signature.txt", options, os.FileMode(0600)) // открывает файл
	check(err)
	_, err = fmt.Fprintln(file, signature) // записывает текст в новой строке файла
	//_, err = writer.Write([]byte(signature)) // записывает значение поля в ответ
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound) //Необходимо передать Redirect значение ResponseWriter,исходный запрос (request),путь для перенапрвления и код ответа
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)          // Функция viewHandler вызывается для любых запросов с путем «/guestbook»
	http.HandleFunc("/guestbook/new", newHandler)       // Запросы на получение формы HTML будут обрабатываться функцией newHandler
	http.HandleFunc("/guestbook/create", createHandler) // Запросы на отправку формы будут обрабатываться функцией createHandler

	err := http.ListenAndServe("localhost:8080", nil) // сервер запускается для прослушивания с портом 8080
	log.Fatal(err)                                    // ошибка никогда не равна nil, поэтому мы никогда не вызываем для нее check

}
