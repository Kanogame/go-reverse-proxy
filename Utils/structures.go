package utils

type Http struct {
	Port      string
	LogFolder string
}

type UndefinedLocation struct {
	WebPath string
	Utype   string
	Path    string
}

type StaticLocations struct {
	WebPath  string
	FilePath string
}

type ProxyLocations struct {
	WebPath  string
	EndPoint string
}

type LoadLocations struct {
	WebPath   string
	EndPoints map[string]int
}

type Locations struct {
	static []StaticLocations
	proxy  []ProxyLocations
	load   []LoadLocations
}
