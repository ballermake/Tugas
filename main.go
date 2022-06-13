package main

import (
	"fmt"

	gage "github.com/umvgc/Parameter"
	gage2 "github.com/umvgc/Parameter/v2"
)

func main() {
	x := gage.CheckNum(3)
	y := gage2.CheckNum(1, 2, 3, 4, 5)

	fmt.Println("v1: ", x)
	fmt.Println("v2: ", y)
}
