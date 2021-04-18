package main

import "fmt"

func printStringStringMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf(key + value + "\n")
	}
}

func main() {
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}
	printStringStringMap(colors1)

	var colors2 map[string]string
	printStringStringMap(colors2)

	// same as colors2
	colors3 := make(map[string]string)
	printStringStringMap(colors3)

	colors3["white"] = "#ffffff"
	printStringStringMap(colors3)

	colors3["white"] = "White"
	printStringStringMap(colors3)

	delete(colors3, "white")
	printStringStringMap(colors3)

	colors4 := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}
	printStringStringMap(colors4)
}
