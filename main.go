package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var data []cinema

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
		var temp cinema
		temp.city = lines[i]
		i++
		temp.cinema = lines[i]
		i++
		a, err := strconv.ParseFloat(lines[i], 64)
		if err != nil {
			temp.score = a
			i++
		}
		temp.film = lines[i]
		i++
		b, err := strconv.Atoi(lines[i])
		if err == nil {
			temp.number = b
			i++
		}
		for i := 0; i < len(lines); i++ {
			temp.sans = append(temp.sans, lines[i])
			i++
			if lines[i] == "******" {
				break
			}

		}
		data = append(data, temp)
	}
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i].city)
	}

}
