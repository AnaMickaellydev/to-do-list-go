package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func InputInt(prompt string) int {
	for {
		input := InputString(prompt)
		number, err := strconv.Atoi(input)
		if err == nil {
			return number
		}
		fmt.Println("Por favor, digite um número válido.")
	}
}
