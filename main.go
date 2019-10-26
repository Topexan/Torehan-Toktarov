package main

import (
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
	isRow := false
	isColumn := false
	isBox := false
	rBox := row - row%3
	cBox := col - col%3
	for i := 0; i < 9; i++ {
		if table[row][i] == num {
			isRow = true
		}
	}
	for j := 0; j < 9; j++ {
		if table[j][col] == num {
			isColumn = true
		}
	}
	for k := rBox; k < rBox+3; k++ {
		for l := cBox; l < cBox+3; l++ {
			if table[k][l] == num {
				isBox = true
			}
		}
	}
	if !isRow && !isColumn && !isBox {
		return true
	} else {
		return false
	}

}

func hasZeros(table *[9][9]int) bool {
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
}

func solSud(table *[9][9]int) bool {
	if !hasZeros(table) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isOk(table, i, j, num) {
						table[i][j] = num
						if solSud(table) {
							return true
						} else {
							table[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
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
