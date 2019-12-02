package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseNumbers(filename string) []int {
	data, _ := ioutil.ReadFile(filename)

	var ret = []int{}

	numbers := strings.Split(string(data), ",")
	for _, s := range numbers {
		n, _ := strconv.Atoi(s)
		ret = append(ret, n)
	}
	return ret
}

func main() {
	numbers := ParseNumbers("data.txt")
	fmt.Println(numbers)
	for i := 0; i <= len(numbers); i += 4 {
		if numbers[i] == 99 {
			break
		}
		if numbers[i] == 1 {
			// addition
			leftPos := numbers[i+1]
			rightPos := numbers[i+2]
			resPos := numbers[i+3]

			numbers[resPos] = numbers[leftPos] + numbers[rightPos]
		}
		if numbers[i] == 2 {
			// multiplication
			leftPos := numbers[i+1]
			rightPos := numbers[i+2]
			resPos := numbers[i+3]

			numbers[resPos] = numbers[leftPos] * numbers[rightPos]
		}
	}
	fmt.Println(numbers)
}
