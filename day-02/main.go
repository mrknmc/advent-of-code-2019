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

func CalcResult(numbers []int) int {
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
	return numbers[0]
}

func CheckResults(numbers []int) (int, int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			newNumbers := make([]int, len(numbers))
			copy(newNumbers, numbers)
			newNumbers[1] = i
			newNumbers[2] = j
			result := CalcResult(newNumbers)
			if result == 19690720 {
				return i, j
			}
		}
	}
	return 0, 0
}

func main() {
	numbers := ParseNumbers("data.txt")
	// fmt.Println(numbers)

	numbers2 := make([]int, len(numbers))
	copy(numbers2, numbers)

	result := CalcResult(numbers2)
	fmt.Println(result)

	noun, verb := CheckResults(numbers)
	fmt.Println(noun, verb)
}
