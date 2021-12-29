package prose

import "testing"

type testData struct {
	list []string // Сегмент, который передается JoinWithCommas
	want string   //Строка, которую должна вернуть функция JoinWithCommas для этого сегмента.

}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{ // создание сегмента значений testData
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange and pear"},
	}

	for _, test := range tests { //обарбатывает каждое значение testData в сегменте
		got := JoinWithCommas(test.list) //cегмент передается JoinWithCommas
		if got != test.want {            //если полученное возвращаемое значение не равно ожидаемому
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}

// Имя функции должно начинаться с Test. Имя после Test может быть любым.

/* Эти три функции заменит одна (стр. 452)

func TestTwoElements(t *testing.T) { // Функции передается указатель на значение testing.T
	list := []string{"apple", "orange"}
	want := "apple and orange"  // ожидаемое возвращаемое значение
	got := JoinWithCommas(list) // полученное возвращаемое значение
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple,orange and pear" // ожидаемое возвращаемое значение
	got := JoinWithCommas(list)     // полученное возвращаемое значение
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestOneElements(t *testing.T) {
	list := []string{"apple"}
	want := "apple"             // ожидаемое возвращаемое значение
	got := JoinWithCommas(list) // полученное возвращаемое значение
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
*/
