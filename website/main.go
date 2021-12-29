package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16 //возраст не может быть отрицательным
	Money                 int
	Avg_grades, happiness float64
	Hobbies               []string
}

//Устанавливаем методы для объекта User

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) { // любые изменения вносятся по ссылке на объект
	u.Name = newName
}

//обращаясь к w типа http.ResponseWriter мы ВЫВОДИМ содержимое сайта на экран
//r *http.Request остлеживает работу метода POST (передачи введеных на сайте данных к серверу)
func home_page(w http.ResponseWriter, r *http.Request) {
	dmitro := User{"Dmitro", 25, -50, 4.2, 0.8, []string{"Football", "Scate", "Dance"}}
	tmpl, _ := template.ParseFiles("templates/home_page.html") // подгружает html шаблон путь/название
	tmpl.Execute(w, dmitro)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	dmitro := User{"Dmitro", 25, -50, 4.2, 0.8, []string{"Footbal", "MMA", "Sleeping"}}
	dmitro.setNewName("Beer")
	// указываем, что будет выводить w на странице при обращении к ней
	fmt.Fprintf(w, dmitro.getAllInfo())
}

func handleRequest() {
	http.HandleFunc("/", home_page) //отслеживает переход по указанному URL-адресу ('/'), вызывая метод (home-page), указанный для данного адреса
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil) //указываем порт для прослушивания (8080),
}

func main() {

	//Dmitro := User{name: "Dmitro", age: 25, money: -50, avg_grades: 4.3, happiness: 0.8}
	handleRequest()

}
