package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func CalcFuel(val float64) float64 {
	return math.Floor(val/3) - 2
}

func IterFuel(fuel float64) float64 {
	finalFuel := fuel
	stepFuel := fuel
	for {
		moreFuel := CalcFuel(stepFuel)
		if moreFuel <= 0 {
			break
		}
		finalFuel += moreFuel
		stepFuel = moreFuel
	}
	return finalFuel
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	fuel := 0.0
	fs := make([]float64, 0)
	for fscanner.Scan() {
		i, _ := strconv.ParseFloat(fscanner.Text(), 64)
		f := CalcFuel(i)
		fs = append(fs, f)
		fuel += f
	}
	fmt.Printf("Fuel: %.2f\n", fuel)

	finalFuel := 0.0
	for _, f := range fs {
		extraFuel := IterFuel(f)
		finalFuel += extraFuel
	}
	fmt.Printf("Final fuel: %.2f\n", finalFuel)
}
