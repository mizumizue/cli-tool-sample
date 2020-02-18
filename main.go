package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func scan() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

var commands = map[int]string{
	1: "fetch all users",
	2: "find user by user_id",
	3: "create user",
	4: "update user",
	5: "delete user",
}

func main() {
	fmt.Printf("start user cli\n\n")

	for i := 1; i <= len(commands); i++ {
		fmt.Printf("%d: %s\n", i, commands[i])
	}

	fmt.Println("\nplease select command")

	var selected int
	for selected == 0 {
		text := scan()
		i, err := strconv.Atoi(text)
		if err != nil || commands[i] == "" {
			fmt.Println("please input list number")
			continue
		}
		selected = i
	}

	fmt.Printf("your input is %s\n", commands[selected])
}
