package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	size := 0
	x := 0
	y := 0
	l := 0
	var arr []rune
	for {
		s, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		arr = append(arr, s)
		size++
	}
	if size > 1 {
		if arr[0] != 'o' && arr[0] != '/' && arr[0] != 'A' {
			fmt.Println("Not a Raid function")
		}
	}
	for i := range arr {
		if arr[i] == '\n' {
			y++
		}
		l++
		i = i
	}
	x = (l - y) / y
	if arr[0] == 'o' {
		fmt.Println("[raid1a] [", x, "] [", y, "]")
	} else if arr[0] == '/' {
		fmt.Println("[raid1b] [", x, "] [", y, "]")
	} else if arr[0] == 'A' {
		if x == 1 && y == 1 {
			fmt.Println("[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]")
		} else if x == 1 && y > 0 && arr[l-2] == 'C' {
			fmt.Println("[raid1c] [1] [", y, "] || [raid1e] [1] [", y, "]")
		} else if x == 1 && y > 0 && arr[l-2] == 'A' {
			fmt.Println("[raid1d] [1] [", y, "]")
		} else if x > 0 && y == 1 && arr[l-2] == 'C' {
			fmt.Println("[raid1c] [1] [", y, "] || [raid1e] [1] [", y, "]")
		} else if x > 0 && y == 1 && arr[l-2] == 'a' {
			fmt.Println("[raid1d] [1] [", y, "]")
		} else if arr[0] == 'A' && arr[l-2-y-1] == 'C' && arr[l-2] == 'C' {
			fmt.Println("[raid1c] [", x, "] [", y, "]")
		} else if arr[0] == 'A' && arr[l-2-y-1] == 'A' && arr[l-2] == 'C' {
			fmt.Println("[raid1c] [", x, "] [", y, "]")
		} else if arr[0] == 'A' && arr[l-2-y-1] == 'C' && arr[l-2] == 'A' {
			fmt.Println("[raid1c] [", x, "] [", y, "]")
		}
	}
}
