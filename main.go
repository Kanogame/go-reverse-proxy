package main

import (
	"fmt"
	config "main/Config"
	httpserver "main/HttpServer"
	utils "main/Utils"
)

func main() {
	configFile := config.ReadConfigFile("./config.txt")
	configStruct, UndefinedServers := config.ParseConfig(configFile)
	fmt.Println("Config read & parsed successfully")
	static, proxy, load := config.DefineServers(UndefinedServers)
	var location = utils.Locations{Static: &static, Proxy: &proxy, Load: &load}
	httpserver.StartHttpServer(configStruct.Port, &location, &configStruct)
}
