package main

import (
	"bufio"
	"fmt"
	utils "main/Utils"
	"os"
	"text/scanner"
)

func main() {
	config := ReadConfigFile("./config")
}

func ReadConfigFile(path string) string[] {
	file, err := os.Open(path)
	if utils.HandleUserError(err) {
		utils.PrintError("No config specified, no config in default directory")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}