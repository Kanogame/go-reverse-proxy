package config

import (
	"bufio"
	utils "main/Utils"
	"os"
	"strings"
)

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

func ParseConfig(config []string) (HttpConfiguration utils.Http) {
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

	//get config
	for i := 0; i < len(ParsedConfig); i++ {
		if strings.Contains(ParsedConfig[i], "port") {
			HttpConfiguration.Port = ConfigGetValue(ParsedConfig[i])
		} else if strings.Contains(ParsedConfig[i], "log") {
			HttpConfiguration.LogFolder = ConfigGetValue(ParsedConfig[i])
		} else if strings.Contains(ParsedConfig[i], "location") {

		}
	}
	return HttpConfiguration
}

func ConfigParseLocation(config []string, start int) {

}

func GetLocationPath(config string) string {
	var path string
	var isValue = false
	for i := 0; i < len(config); i++ {
		val := string(config[i])
		if val == "(" {
			isValue = !isValue
			continue
		} else if val == ")" {
			break
		}
		if isValue {
			path += string(val)
		}
	}
	return path
}

func ConfigGetValue(line string) string {
	var value string
	var isValue = false
	for i := 0; i < len(line); i++ {
		if isValue {
			value += string(line[i])
		}
		if line[i] == '"' {
			isValue = !isValue
		}
	}
	if len(value)-1 <= 0 {
		utils.PrintError("Error while parcing config: no value is specified")
		return ""
	}
	return value[:len(value)-1]
}
