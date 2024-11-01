package main

import (
	"fmt"
	"os"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	if contents, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(grade(100))
	fmt.Println(grade(30))
	fmt.Println(grade(59))
	fmt.Println(grade(60))
	fmt.Println(grade(82))
	fmt.Println(grade(99))
	fmt.Println(grade(100))
}
