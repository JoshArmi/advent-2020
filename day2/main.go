package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		txt := scanner.Text()
		if len(txt) > 0 {
			txtlines = append(txtlines, txt)
		}
	}

	file.Close()

	var validPasswordCount int

	for _, line := range txtlines {
		splits := strings.Split(line, ":")
		min, _ := strconv.Atoi(strings.Split(splits[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(splits[0], "-")[1], " ")[0])
		char := strings.Split(splits[0], " ")[1]
		count := strings.Count(splits[1], char)
		if min <= count && count <= max {
			validPasswordCount++
		}
	}

	fmt.Println(validPasswordCount)

	validPasswordCount = 0

	for _, line := range txtlines {
		splits := strings.Split(line, ":")
		min, _ := strconv.Atoi(strings.Split(splits[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(splits[0], "-")[1], " ")[0])
		char := strings.Split(splits[0], " ")[1]
		password := strings.Trim(splits[1], " ")
		first := string(password[min-1]) == char
		second := string(password[max-1]) == char
		if (first && !second) || (!first && second) {
			validPasswordCount++
		}
	}

	fmt.Println(validPasswordCount)

}
