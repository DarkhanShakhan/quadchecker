package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\t')
	length := len(name)
	columns := 0
	for _, v := range name {
		if v == '\n' {
			break
		}
		columns++
	}
	rows := length / columns
	fmt.Println(name == quadA(rows, columns))
	fmt.Println(name == quadB(rows, columns))
}

func quadA(h, v int) string {
	result := ""
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if i == 1 && j == 1 || i == 1 && j == v || i == h && j == v || i == h && j == 1 {
				result += "o"
			} else if i == 1 || i == h {
				result += "-"
			} else if j == 1 || j == v {
				result += "|"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadB(h, v int) string {
	result := ""
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if i == 1 && j == 1 {
				result += "/"
			} else if i == 1 && j == v || i == h && j == 1 {
				result += "\\"
			} else if i == h && j == v {
				result += "/"
			} else if i == 1 || i == h || j == 1 || j == v {
				result += "*"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}
