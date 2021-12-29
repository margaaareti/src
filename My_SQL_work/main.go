package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func check(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

type User struct {
	Name string `json:"name"`
	Age  uint16
}

func main() {
	//подключаем БД, указывая название БД, логин:пароль@tcp(хост:порт)/название БД
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang_bd")
	check(err)
	defer db.Close() // закрываем БД

	//Добавляем элементы в столбцы таблицы
	//insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Alex', 25)") //Добавляем элементы в столбцы таблицы
	//check(err)
	//defer insert.Close() //закрываем INSERT

	//Выборка данных
	res, err := db.Query("SELECT `name`,`age` FROM `users` ")
	check(err)

	//Метод Next выдает true, если есть строки для обработки
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age) //проверяем, существуют ли значения
		check(err)

		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}
}
