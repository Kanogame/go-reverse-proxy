package config

import utils "main/Utils"

func DefineServers(servers []utils.UndefinedLocation) (static []utils.StaticLocations, proxy []utils.ProxyLocations, load []utils.LoadLocations) {
	for _, server := range servers {
		if server.Utype == "static" {
			static = append(static, utils.StaticLocations{WebPath: server.WebPath, FilePath: server.Path})
		} else if server.Utype == "proxy" {
			proxy = append(proxy, utils.ProxyLocations{WebPath: server.WebPath, EndPoint: server.Path})
		} else if server.Utype == "proxy_load" {
			endPoints := separateEndPoints(server.WebPath)
			ServerEndPoints := make(map[string]int)
			for _, endPoint := range endPoints {
				ServerEndPoints[endPoint] = 0
			}
			load = append(load, utils.LoadLocations{WebPath: server.WebPath, EndPoints: ServerEndPoints})
		}
	}
	return static, proxy, load
}

func separateEndPoints(endPoints string) []string {
	var endPointsArr []string
	var current string
	for i := 0; i < len(endPoints); i++ {
		if string(endPoints[i]) == `"` {
			current = ""
			endPointsArr = append(endPointsArr, current)
			continue
		}
		current += string(endPoints[i])
	}
	return endPointsArr
}
