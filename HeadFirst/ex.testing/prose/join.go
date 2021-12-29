package prose

import (
	"strings"
)

func JoinWithCommas(phrases []string) string { // получаем сегмет строк для объединения
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ") // разделяет все строки, кроме последней, запятыми
		result += " and "                                      // Вставляет слово "and" перед последней строкой
		result += phrases[len(phrases)-1]                      // добавляет последнюю строку
		return result
	}

}
