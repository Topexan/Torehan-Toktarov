package main

import (
	"github.com/01-edu/z01"
	"os"
)

func strtoNum(s string) int {
	runes := []rune(s)
	num := 0
	for i := range runes {
		num *= 10
		if runes[i] == '1' {
			num = num + 1
		} else if runes[i] == '2' {
			num = num + 2
		} else if runes[i] == '3' {
			num = num + 3
		} else if runes[i] == '4' {
			num = num + 4
		} else if runes[i] == '5' {
			num = num + 5
		} else if runes[i] == '6' {
			num = num + 6
		} else if runes[i] == '7' {
			num = num + 7
		} else if runes[i] == '8' {
			num = num + 8
		} else if runes[i] == '9' {
			num = num + 9
		} else if runes[i] == '0' {
			num = num + 0
		}
	}
	return num
}

func Raid1c(x, y int) {
	if y >= 1 {
		if x > 0 {
			z01.PrintRune(65)
		}
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(66)
		}
		if x > 1 {
			z01.PrintRune(65)
		}
		if x > 0 {
			z01.PrintRune(10)
		}

	}

	if y > 1 {
		if x > 0 {
			for j := 1; j <= y-2; j++ {
				z01.PrintRune(66)
				for i := 1; i <= x-2; i++ {
					z01.PrintRune(32)
				}
				if x > 1 {
					z01.PrintRune(66)
				}
				z01.PrintRune(10)
			}
		}
	}

	if y > 1 {
		if x > 0 {
			z01.PrintRune(67)
		}
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(66)
		}
		if x > 1 {
			z01.PrintRune(67)
		}
		if x > 0 {
			z01.PrintRune(10)
		}
	}
}

func main() {
	arguments := os.Args
	l := len(arguments)
	if l == 3 {
		Raid1c(strtoNum(arguments[1]), strtoNum(arguments[2]))
	}
}
