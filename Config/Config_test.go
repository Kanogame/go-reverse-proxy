package config

import "testing"

func TestGetValue(t *testing.T) {
	t.Log("Testing GetValue")

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"test 1", `port: "123"`, "123"},
		{"test 1", `	log: "./latest.log"`, "./latest.log"},
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
