package utils

import (
	"fmt"
	"strconv"
)

func ReadInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func StringToInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Erro na conversão:", err)
		return -1
	}
	return val
}
