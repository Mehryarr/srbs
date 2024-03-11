package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	var data []a
	var city []string
	var cinema []string
	var input1 int
	m1 := make(map[string]float64)

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
	//fmt.Println(data)
	sw := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(city); j++ {
			if data[i].city == city[j] {
				sw++
				break
			}
		}
		if sw == 0 {
			city = append(city, data[i].city)
		}
		sw = 0
	}

	for i := 0; i < len(city)-1; i++ {
		for j := 0; j < len(city)-i-1; j++ {
			if city[j] > city[j+1] {
				temp := city[j]
				city[j] = city[j+1]
				city[j+1] = temp
			}
		}
	}

	cls()
	for i := 0; i < len(city); i++ {
		fmt.Println(i+1, city[i])
	}

	for i := 0; i < 2; {
		fmt.Println("input city number")
		fmt.Scan(&input1)
		if input1 <= len(city) && input1 > 0 {
			break
		}
	}

	for i := 0; i < len(data); i++ {
		if city[input1-1] == data[i].city {
			m1[data[i].cinema] = data[i].score
		}
	}
	for k := range m1 {
		cinema = append(cinema, k)
	}
	for i := 0; i < len(cinema)-1; i++ {
		for j := 0; j < len(cinema)-i-1; j++ {
			if m1[cinema[j]] < m1[cinema[j+1]] {
				temp := cinema[j]
				cinema[j] = cinema[j+1]
				cinema[j+1] = temp
			}
		}
	}

}
