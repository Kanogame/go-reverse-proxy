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
	t.Log("Testing GetValue")

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
	t.Log("Testing GetValue")

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
