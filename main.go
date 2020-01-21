package main

import (
	"fmt"

	"github.com/ekr-paolo-carraro/algorithm/utils"
)

func main() {
	names := []string{
		"paolo",
		"tiziana",
		"francesco",
		"aurora",
		"franco",
		"paolo",
		"tiziana",
		"francesco",
		"aurora",
		"franco",
		"paolo",
		"tiziana",
		"zorro",
		"francesco",
		"aurora",
		"franco",
		"paolo",
		"tiziana",
		"francesco",
		"aurora",
		"zorro",
		"franco",
		"paolo",
		"tiziana",
		"francesco",
		"aurora",
		"franco",
		"zorro",
	}

	pos, err := utils.SearchItem("paolo", utils.Sort(names))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pos)
}
