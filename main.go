package main

import config "main/Config"

func main() {
	configFile := config.ReadConfigFile("./config.txt")
	config.ParseConfig(configFile)
}
