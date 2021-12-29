package main

import (
	"HeadFirst/ex.struct/example2/geo"
	"fmt"
	"log"
)

func main() {

	location := geo.Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLongtitude(-122.8)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Coordinates.Longtitude())

}
