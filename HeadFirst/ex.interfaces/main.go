package main

import "HeadFirst/ex.interfaces/gadget"

type Player interface {
	Play(string) //должен содержать метод Play, с параметром string
	Stop()       // также необходим метод Stop
}

func playList(device Player, songs []string) {
	for _, song := range songs { // перебирает все песни в цикле
		device.Play(song) // воспроизведение текущей песни
	}
	device.Stop() //плеер останавливается после завершения
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"} // Создается сегмент с названиями песен
	var player Player = gadget.TapePlayer{}                   //Обнов. переменную, для хран. люб. значения,поддерживающего Player
	playList(player, mixtape)                                 // TapePlayer передается в playList
	player = gadget.TapeRecorder{}
	playList(player, mixtape)

}
