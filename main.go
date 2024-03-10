package main

import (
	"fmt"
	"os"
)

type cinema struct {
	city   string
	cinema string
	score  float64
	film   string
	number int
	sans   []string
}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()

}
