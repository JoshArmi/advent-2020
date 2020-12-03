package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, line := range txtlines {
		for _, line2 := range txtlines {
			first, _ := strconv.Atoi(line)
			second, _ := strconv.Atoi(line2)
			if first+second == 2020 {
				fmt.Println(strconv.Itoa(first * second))
			}
		}
	}

	for _, line := range txtlines {
		for _, line2 := range txtlines {
			for _, line3 := range txtlines {
				first, _ := strconv.Atoi(line)
				second, _ := strconv.Atoi(line2)
				third, _ := strconv.Atoi(line3)
				if first+second+third == 2020 {
					fmt.Println(strconv.Itoa(first * second * third))
				}
			}
		}
	}
}
