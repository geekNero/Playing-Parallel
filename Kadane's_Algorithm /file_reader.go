package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile() []string {
	data, err := os.ReadFile("vlarge.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}
	list := strings.Split(string(data), "\n")
	return list
}
