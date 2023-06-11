package main

import config "main/Config"

func main() {
	configFile := config.ReadConfigFile("./config.txt")
	_, UndefinedServers := config.ParseConfig(configFile)
	config.DefineServers(UndefinedServers)
}
