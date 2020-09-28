package projecteuler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	var size []int
	var array2d [][]int

	err := loadData(&size, &array2d)

	if err == nil {
		//for _, row := range array2d {
		//	fmt.Println(row)
		//}

		highestHorizontal := findHighestHorizontal(&array2d)
		fmt.Println(highestHorizontal)

		highestVertical := findHighestVertical(&array2d)
		fmt.Println(highestVertical)

		highestDiagonal := findHighestDiagonal(&array2d)
		fmt.Println(highestDiagonal)
	} else {
		fmt.Println(err.Error())
	}
}

func loadData(size *[]int, outputArray *[][]int) error {
	reader := bufio.NewReader(os.Stdin)
	sizeString, ok := reader.ReadString('\n')
	if ok == nil {
		sizeSlice := strings.Fields(sizeString)
		x, _ := strconv.Atoi(sizeSlice[0])
		y, _ := strconv.Atoi(sizeSlice[1])

		*size = []int{x, y}
		*outputArray = make([][]int, y)

		for i := 0; i < y; i++ {
			rowString, _ := reader.ReadString('\n')
			rowSlice := strings.Fields(rowString)

			for j := 0; j < x; j++ {
				num, _ := strconv.Atoi(rowSlice[j])
				(*outputArray)[i] = append((*outputArray)[i], num)
			}
		}

		return nil
	} else {
		return errors.New("error loading")
	}
}

func product(array []int) int {
	result := 1
	for _, s := range array {
		result *= s
	}
	return result
}

func findHighestHorizontal(array2d *[][]int) int {
	highestProduct := -1

	for _, row := range *array2d {
		for i := 0; i < len(row) - 3; i++ {
			tmp := product(row[i:i+3])

			if tmp > highestProduct {highestProduct = tmp}
		}
	}

	return highestProduct
}

func findHighestVertical(array2d *[][]int) int {
	highestProduct := -1

	for i := 0; i < len(*array2d) - 3; i++ {
		for j := 0; j < len((*array2d)[0]); j++ {
			tmp := product([]int {
				(*array2d)[i][j],
				(*array2d)[i + 1][j],
				(*array2d)[i + 2][j],
				(*array2d)[i + 3][j]})

			if tmp > highestProduct {highestProduct = tmp}
		}
	}

	return highestProduct
}

func findHighestDiagonal(array2d *[][]int) int {
	highestProduct := -1

	for i := 0; i < len(*array2d) - 3; i++ {
		for j := 0; j < len((*array2d)[0]) - 3; j++ {
			tmp := product([]int {
				(*array2d)[i][j],
				(*array2d)[i + 1][j + 1],
				(*array2d)[i + 2][j + 2],
				(*array2d)[i + 3][j + 3]})

			if tmp > highestProduct {highestProduct = tmp}
		}
	}

	for i := 3; i < len(*array2d); i++ {
		for j := 3; j < len((*array2d)[0]); j++ {
			tmp := product([]int {
				(*array2d)[i][j],
				(*array2d)[i - 1][j - 1],
				(*array2d)[i - 2][j - 2],
				(*array2d)[i - 3][j - 3]})

			if tmp > highestProduct {highestProduct = tmp}
		}
	}

	return highestProduct
}