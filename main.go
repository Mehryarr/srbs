package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type a struct {
	city   string
	cinema string
	score  float64
	film   string
	number int
	sans   []string
}

func main() {

	var data []a

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	file.Close()

	for i := 0; i < len(lines); i++ {
		var temp a
		temp.city = lines[i]
		i++
		temp.cinema = lines[i]
		i++
		b, err := strconv.ParseFloat(lines[i], 64)
		if err == nil {
			temp.score = b
			i++
		}
		temp.film = lines[i]
		i++
		c, err := strconv.Atoi(lines[i])
		if err == nil {
			temp.number = c
			i++
		}
		for j := 0; j < len(lines); j++ {
			temp.sans = append(temp.sans, lines[i])
			i++
			if lines[i] == "******" {
				break
			}
		}
		data = append(data, temp)
	}
	fmt.Println(data)
}
