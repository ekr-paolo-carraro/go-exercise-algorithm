package main

import (
	"fmt"

	"github.com/ekr-paolo-carraro/go-exercise-algorithm/utils"
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

	utils.Sort(names)
	pos, err := utils.SearchItem("paolo", names)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pos)

}
