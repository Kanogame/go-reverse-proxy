package config

import (
	utils "main/Utils"
	"testing"
)

func TestGetValue(t *testing.T) {
	t.Log("Testing GetValue")

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"test 1", `port: "123"`, "123"},
		{"test 2", `	log: "./latest.log"`, "./latest.log"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ConfigGetValue(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestGetLocation(t *testing.T) {
	t.Log("Testing Location")

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"test 1", `location(/app) {`, "/app"},
		{"test 2", `	location(/) {`, "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := GetLocationPath(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestParseLocation(t *testing.T) {
	t.Log("Testing ParsingLocation")

	var tests = []struct {
		name   string
		input1 []string
		input2 int
		want   utils.UndefinedLocation
	}{
		{"test 1", []string{"location(/) {", `type: "static";`, `path: "./static";`}, 0, utils.UndefinedLocation{Utype: "static", Path: "./static", WebPath: "/"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ConfigParseLocation(tt.input1, tt.input2)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestParseConfig(t *testing.T) {
	want1 := utils.Http{LogFolder: "./latest.log", Port: "80"}
	want2 := []utils.UndefinedLocation{utils.UndefinedLocation{Utype: "static", Path: "./static", WebPath: "/"}}
	t.Log("Testing ParcingConfig")
	t.Run("test", func(t *testing.T) {
		ans1, ans2 := ParseConfig([]string{"http {", `port: "80";`, `port: "./latest.log";`, "location(/) {", `type: "static";`, `path: "./static";`, "}", "}"})
		if ans1 != want1 && ans2[0] != want2[0] {
			t.Errorf("got %s, want %s", ans1, ans2)
		}
	})
}

func TestSepatateEndPoints(t *testing.T) {
	t.Log("Testing SepatateEndPoints")

	var tests = []struct {
		name  string
		input string
		want  []string
	}{
		{"test 1", `https://0.0.0.1:1"https://0.0.0.0:1`, []string{"https://0.0.0.1:1", "https://0.0.0.0:1"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := separateEndPoints(tt.input)
			for i, val := range ans {
				if val != tt.want[i] {
					t.Errorf("got %s, want %s", ans, tt.want[i])
				}
			}
		})
	}
}

func TestDefineServer(t *testing.T) {
	t.Log("Testing DefineServer")

	var tests = []struct {
		name  string
		input []utils.UndefinedLocation
		want1 []utils.StaticLocations
		want2 []utils.ProxyLocations
		want3 []utils.LoadLocations
	}{
		{"test 1",
			[]utils.UndefinedLocation{{Utype: "proxy", WebPath: "/", Path: "https://0.0.0.0:1"}},
			nil,
			[]utils.ProxyLocations{{WebPath: "/", EndPoint: "https://0.0.0.0:1"}},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans1, ans2, _ := DefineServers(tt.input)
			for i, val := range ans1 {
				if val != tt.want1[i] {
					t.Errorf("got %s, want %s", ans1, tt.want1[i])
				}
			}
			for i, val := range ans2 {
				if val != tt.want2[i] {
					t.Errorf("got %s, want %s", ans2, tt.want2[i])
				}
			}
			/*for i, val := range ans3 {
				if val != tt.want3[i] {
					t.Errorf("got %s, want %s", ans3, tt.want3[i])
				}
			}*/
		})
	}
}
