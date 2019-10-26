package main

import (
	//	"github.com/01-edu/z01"
	"fmt"
	"os"
)

func PrintNbr(n rune) int {
	if n == '.' {
		return 0
	} else if n == '1' {
		return 1
	} else if n == '2' {
		return 2
	} else if n == '3' {
		return 3
	} else if n == '4' {
		return 4
	} else if n == '5' {
		return 5
	} else if n == '6' {
		return 6
	} else if n == '7' {
		return 7
	} else if n == '8' {
		return 8
	} else if n == '9' {
		return 9
	} else {
		return 0
	}
}

func FillTables(str []string) [9][9]int {
	var table [9][9]int
	for i := 0; i < 9; i++ {
		runes := []rune(str[i+1])
		for j := 0; j < 9; j++ {
			table[i][j] = PrintNbr(rune(runes[j]))
		}
	}
	return table
}

func isOk(table *[9][9]int, row, col, num int) bool {
	//isRow := false
	//isColumn := false
	//isBox := false
	rBox := row - row%3
	cBox := col - col%3
	for i := 0; i < 9; i++ {
		if table[row][i] == num {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if table[j][col] == num {
			return false
		}
	}
	for k := rBox; k < rBox+3; k++ {
		for l := cBox; l < cBox+3; l++ {
			if table[k][l] == num {
				return false
			}
		}
	}

	return true
	//if !isRow && !isColumn && !isBox {
	//	return true
	//} else {
	//	return false
	//}

}

/*func hasZeros(table [9][9]int) bool{
	r := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				r++
			}
		}
	}
	if r == 0 {
		return false
	} else {
		return true
	}
}*/

func solSud(table *[9][9]int) bool {
	ii := 0
	jj := 0
	isThereEmptyCell := false

	for i := 0; i < 9 && !isThereEmptyCell; i++ {
		for j := 0; j < 9 && !isThereEmptyCell; j++ {
			if table[i][j] == 0 {
				isThereEmptyCell = true
				ii = i
				jj = j
			}
		}
	}

	if !isThereEmptyCell {
		return true
	}

	for x := 1; x < 10; x++ {
		if isOk(table, ii, jj, x) {
			table[ii][jj] = x

			if solSud(table) {
				return true
			}
			table[ii][jj] = 0
		}
	}
	return false
}

func main() {
	arguments := os.Args
	table := FillTables(arguments)
	if solSud(&table) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Print(table[i][j], " ")
			}
			fmt.Println()
		}
	}
}
