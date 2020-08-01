package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func criminals_run() {
	var groupA, groupB []bool
	fillWithProbability(&groupA, 10000000, 5)
	fillWithProbability(&groupB, 10000000, 10)

	x, y := getSampleSizeAndProbability(groupA, 1e-20)

	fmt.Printf("%d %f%%\n", x, y * 100)

	x, y = getSampleSizeAndProbability(groupB, 1e-20)

	fmt.Printf("%d %f%%\n", x, y * 100)
}

func fillWithProbability(list *[]bool, number int, probability int) {
	for i := 0; i < number; i++ {
		value := rand.Int() % 100
		if value < probability {
			*list = append(*list, true)
		} else {
			*list = append(*list, false)
		}
	}
}

func countInnocentsAndCriminals(list []bool) (innocents int, criminals int) {
	for _, x := range list {
		if x {
			criminals += 1
		} else {
			innocents += 1
		}
	}
	return
}

func getSampleSizeAndProbability(list []bool, precision float64) (int, float64) {
	var measurements []float64

	sampleSize := 1000
	standardDeviation := math.MaxFloat64

	for standardDeviation > precision {
		sampleSize *= 10

		innocents, criminals := countInnocentsAndCriminals(list[:sampleSize])
		measurements = append(measurements, float64(criminals) / float64(innocents))
		standardDeviation = calculateStandardDeviation(measurements)
	}

	if measurements == nil {
		return 0, 0
	}

	return sampleSize, measurements[len(measurements)-1]
}

func calculateStandardDeviation(list []float64) float64 {
	if len(list) == 1 {
		return math.MaxFloat64
	}

	var average = calculateAverage(list)

	var magicSum float64
	for _, x := range list {
		magicSum += math.Pow(x - average, 2)
	}

	var variance = float64(1 / (len(list) - 1)) * magicSum

	return math.Sqrt(variance)
}

func calculateAverage(list []float64) float64 {
	var sum float64
	for _, x := range list {
		sum += x
	}
	return sum / float64(len(list))
}

func printSlice(slice []bool) {
	for _, x := range slice {
		print(strconv.FormatBool(x) + " ")
	}
	println()
}
