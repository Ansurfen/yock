package cmd

import "testing"

func TestHTTP(t *testing.T) {
	HTTP(HttpOpt{
		Method: "GET",
		Save:   true,
		Debug:  true,
		Dir:    ".",
		Filename: func(s string) string {
			return s
		},
	}, []string{"https://www.github.com"})
}
