package main

import config "main/Config"

func main() {
	config := config.ReadConfigFile("./config.txt")
	config.ParseConfig(config)
}
