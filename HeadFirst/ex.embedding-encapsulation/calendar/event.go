package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date  //значение date встраивается в виде анонимного поля
}

func (e *Event) Title() string { //get метод
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 { //если после содержит более 30 символов, возвращается ошибка
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
