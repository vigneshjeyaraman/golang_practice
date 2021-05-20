package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number you want to check")
	num1, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num1 = strings.TrimSpace(num1)
	if num1[0] == num1[len(num1)-1] {
		fmt.Println("Number is even ended")
	}
	num2 := 1121
	s := fmt.Sprintf("%d", num2)
	if s[0] == s[len(num1)-1] {
		fmt.Println("Number is even ended")
	}

}
