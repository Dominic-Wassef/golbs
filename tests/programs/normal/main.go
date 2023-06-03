package main

import "fmt"

/*
TESTING go_build_system (golbs)
*/

func main() {
	data := []string{
		"Chicago",
		"New York",
		"Dallas",
		"Miami",
		"Orlando",
		"Detroit",
		"Phoenix",
		"San Francisco",
		"Seattle",
		"Portland",
		"Buffalo",
		"Atlanta",
		"Houston",
		"Austin",
		"Tampa",
		"Jacksonville",
		"San Diego",
		"Las Vegas",
	}
	fmt.Println(filterCities(data, 'O'))
}

func filterCities(content []string, letter rune) (out []string) {
	const index = 0
	for _, n := range content {
		if rune(n[index]) == letter {
			out = append(out, n)
		}
	}
	return
}
