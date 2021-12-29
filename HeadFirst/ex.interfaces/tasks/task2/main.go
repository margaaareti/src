package main

import "HeadFirst/ex.interfaces/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder) // Утверждение типа используется для перехода по значение TapeRecorder
	if ok {
		recorder.Record() // метод Record вызывается только в том случае, если исходное значение имело тип TapeRecorder
	}

}

func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})

}
