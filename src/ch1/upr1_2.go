// Echo1 выводит аргументы командной строки

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args[1:] {
		fmt.Println(i, s)
	}
}