package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	// get city name
	cityReader := bufio.NewReader(os.Stdin)
	println("What's the name of the city you grew up in?")
	cityName, _ := cityReader.ReadString('\n')
	cityName = strings.TrimSpace(cityName)

	// get pet name
	petReader := bufio.NewReader(os.Stdin)
	println("What's your pet's name?")
	petName, _ := petReader.ReadString('\n')
	petName = strings.TrimSpace(petName)

	// generate band name
	println("You name could be '", cityName, petName, "'!")
}
