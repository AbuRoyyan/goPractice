package main

import (
	"fmt"
	"strconv"
)

func hexToDecimal(hexStr string) (int64, error) {
	result, err := strconv.ParseInt(hexStr, 16, 64)
	return result, err
}

func main() {
	res, err := hexToDecimal("FF")
	fmt.Println(res, err)
}