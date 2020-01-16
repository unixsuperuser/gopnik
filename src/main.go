package main

import (
	"./pixelcombiner"
)

func main() {
	test1 := pixelcombiner.LoadImage("../test.png")
	test2 := pixelcombiner.LoadImage("../b.png")
	pixelcombiner.Combine(test1, test2)
	// fmt.Println(combined)
}