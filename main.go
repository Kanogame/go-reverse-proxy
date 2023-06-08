package main

import (
	"bufio"
	utils "main/Utils"
	"os"
	"strings"
)

func main() {
	config := ReadConfigFile("./config.txt")
	ParseConfig(config)
}

func ReadConfigFile(path string) []string {
	file, err := os.Open(path)
	if !utils.HandleUserError(err) {
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

func ParseConfig(config []string) {
	var ParsedConfig []string
	//get work area
	for i := 0; i < len(config); i++ {
		if strings.Contains(config[i], "http") {
			var end int
			for j := 0; j < len(config); j++ {
				if strings.Contains(config[(len(config)-1)-j], "}") {
					end = len(config) - 1 - j
					break
				}
			}
			ParsedConfig = config[i+1 : end-1]
			break
		}
	}

}
