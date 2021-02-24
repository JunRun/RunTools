package remd

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed hello.txt
var s string

func TestEmbed_1(t *testing.T) {
	fmt.Println(s)
}

//go:embed *
var f embed.FS

func TestEmbed_2(t *testing.T) {
	confFile, _ := f.ReadFile("p/conf")

	fmt.Println(string(confFile))
}
