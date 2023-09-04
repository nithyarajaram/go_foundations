package main

import (
	"fmt"
	"log"
)

func main() {

	//fmt.Println(div(2, 1))
	fmt.Println(safeDiv(2, 0))
}

func div(a, b int) int {

	return a / b
}

func safeDiv(a, b int) (q int, err error) {

	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	return a / b, err
}
