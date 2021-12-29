package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	Id                        uint16
	Title, Announce, FullText string
}

var posts = []Article{} //Срез со значением переданных с форм данных
var showPost = Article{}

func check(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error()) //Выводим ошибку на странице сайта
	}

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	check(err)
	defer db.Close()

	//Выьорка данных. Выбираем все поля из БД "articles"
	res, err := db.Query("SELECT * FROM `articles`")
	check(err)

	posts = []Article{} //обнуляем список каждый раз заходя на страницу
	for res.Next() {
		var post Article
		//обращаясь к объекту post помещаем в него данные
		err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.FullText)
		check(err)

		posts = append(posts, post)
	}

	//Указываем w, название блока для вывода на странице
	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error()) //Выводим ошибку на странице сайта
	}
	//Указываем w, название блока для вывода на странице
	t.ExecuteTemplate(w, "create", nil)
}

//получаем данные из формы и обрабатываем их
func save_article(w http.ResponseWriter, r *http.Request) {
	//указывается имя поля (name),из которого получаем данные
	title := r.FormValue("title")
	announce := r.FormValue("announce")
	full_text := r.FormValue("full_text")
	//Выводим ошибку при попытке отправления пустых строк через форму
	if title == "" || announce == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		//Подключаемся к базе данных
		db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
		check(err)

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`,`announce`, `full_text` )	VALUES ('%s','%s','%s')", title, announce, full_text))
		check(err)

		defer insert.Close()
		//Переадресовываем пользователя после отправки
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error()) //Выводим ошибку на странице сайта
	}

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	check(err)
	defer db.Close()

	//Выьорка данных. Выбираем все поля из БД "articles"
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	check(err)

	showPost = Article{} //обнуляем список каждый раз заходя на страницу
	for res.Next() {
		var post Article
		//обращаясь к объекту post помещаем в него данные
		err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.FullText)
		check(err)
		showPost = post
	}

	//Указываем w, название блока для вывода на странице
	t.ExecuteTemplate(w, "show", showPost)

}

func handleFunc() {

	//Отслеживание URL
	rtr := mux.NewRouter()

	rtr.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) //добавление стилей/скриптов
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/", rtr)                                                                         // обработка всех URL Адресов будет происходить через роутер mux
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) //добавление стилей/скриптов
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleFunc()

}
