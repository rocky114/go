package algorithm

import "fmt"

//var version string

const (
	coordinate1 = iota + 1
	coordinate2
	coordinate3
	coordinate4
)

var (
	currentRow = 1
	currentCol = 1

	initialRow = 1
	initialCol = 1

	currentCoordinate = coordinate1

	loop = 0
)

func Matrix(n int) [][]int {
	ret := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		item := make([]int, 0, n)
		ret = append(ret, item)
	}

	for i := 1; i <= n*n; i++ {
		if i == 1 {
			fmt.Println("current coordinate: ", i, currentRow, currentCol)
		} else {
			coordinate(n)
			fmt.Println("current coordinate: ", i, currentRow+loop, currentCol+loop)
		}
		//ret[currentRow-1][currentCol-1] = i
	}

	return ret
}

func coordinate(n int) {
	if currentCoordinate == coordinate1 {
		if currentCol < n {
			currentCol++
		} else {
			// currentCol == n
			currentRow++
			currentCoordinate = coordinate2
		}
	} else if currentCoordinate == coordinate2 {
		if currentRow < n {
			currentRow++
		} else {
			// currentRow == n
			currentCol--
			currentCoordinate = coordinate3
		}
	} else if currentCoordinate == coordinate3 {
		if currentCol > initialCol {
			currentCol--
		} else {
			// currentCol == 1
			currentRow--
			currentCoordinate = coordinate4
		}
	} else {
		// currentCoordinate == coordinate4
		if currentRow == initialRow+1 && currentCol == initialCol {
			currentCoordinate = coordinate1
			currentRow--
			loop++

		} else {
			currentRow--
		}
	}
}
