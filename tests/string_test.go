package tests

import (
	`strings`
	`testing`
)

func TestString(t *testing.T) {
	f := strings.Split("index.min.json", ".")
	t.Log(len(f), f[len(f)-1])
}
