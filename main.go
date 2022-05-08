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
	fmt.Println("Columns: ", columns)
	fmt.Println("Rows: ", rows)
}
