package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type cnm struct {
	city   string
	cinema string
	score  float64
	film   string
	number int
	sans   []string
}

// function for clearning commandline
func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	//variables
	var data []cnm
	var city []string
	var cinema []string
	var film []string
	var input1 int
	var input2 int
	var input3 int
	m1 := make(map[string]float64)
	m2 := make(map[string]int)

	//reading from data file
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

	//setting
	for i := 0; i < len(lines); i++ {
		var temp cnm
		temp.city = lines[i]
		i++
		temp.cinema = lines[i]
		i++
		a, err := strconv.ParseFloat(lines[i], 64)
		if err == nil {
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
		for j := 0; j < len(lines); j++ {
			temp.sans = append(temp.sans, lines[i])
			i++
			if lines[i] == "******" {
				break
			}
		}
		data = append(data, temp)
	}
	//deleting repeated cities
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

	//sort city
	for i := 0; i < len(city)-1; i++ {
		for j := 0; j < len(city)-i-1; j++ {
			if city[j] > city[j+1] {
				temp := city[j]
				city[j] = city[j+1]
				city[j+1] = temp
			}
		}
	}

	//print city
	cls()
	for i := 0; i < len(city); i++ {
		fmt.Println(i+1, city[i])
	}

	//input city
	for i := 0; i < 2; {
		fmt.Println("input city number")
		fmt.Scan(&input1)
		if input1 <= len(city) && input1 > 0 {
			break
		}
	}

	//sort cinema
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

	//print cinema
	cls()
	fmt.Println("city :", city[input1-1])
	for i := 0; i < len(cinema); i++ {
		fmt.Println(i+1, "cinema :", cinema[i], "score :", m1[cinema[i]])
	}

	//input cinema
	for i := 0; i < 2; {
		fmt.Println("input cinema number")
		fmt.Scan(&input2)
		if input2 <= len(cinema) && input2 > 0 {
			break
		}
	}

	//sort film
	for i := 0; i < len(data); i++ {
		if cinema[input2-1] == data[i].cinema {
			m2[data[i].film] = data[i].number
		}
	}
	for k := range m2 {
		film = append(film, k)
	}
	for i := 0; i < len(film)-1; i++ {
		for j := 0; j < len(film)-i-1; j++ {
			if m1[film[j]] < m1[film[j+1]] {
				temp := film[j]
				film[j] = film[j+1]
				film[j+1] = temp
			}
		}
	}

	//print film
	cls()
	fmt.Println("city :", city[input1-1])
	fmt.Println("cinema :", cinema[input2-1])
	for i := 0; i < len(film); i++ {
		fmt.Println(i+1, "film :", film[i], "number :", m2[film[i]])
	}

	//input film
	for i := 0; i < 2; {
		fmt.Println("input film number")
		fmt.Scan(&input3)
		if input3 <= len(film) && input3 > 0 {
			break
		}
	}

	for i := 0; i < len(data); i++ {
		if city[input1-1] == data[i].city {
			if cinema[input2-1] == data[i].cinema {
				if film[input3-1] == data[i].film {
					for j := 0; j < len(data[i].sans); j++ {
						fmt.Println(j+1, data[i].sans[j])
					}
				}
			}
		}
	}

}
