package main

import (
	"fmt"
	"strconv"
)

func binToDecimal(binStr string) (int64, error) {
	result, err := strconv.ParseInt(binStr, 2, 64)
	return result, err
}

func main() {
	res, err := binToDecimal("1010")
	fmt.Println(res, err)
}