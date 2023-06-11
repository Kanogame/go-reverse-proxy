package main

import config "main/Config"

func main() {
	configFile := config.ReadConfigFile("./config.txt")
	HttpArgs, UndefinedServers := config.ParseConfig(configFile)

}
