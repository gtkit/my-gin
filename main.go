package main

import (
	"runtime"

	"my_gin/cmd"
)

func main() {
	maxProces := runtime.NumCPU()
	if maxProces > 1 {
		maxProces--
	}
	runtime.GOMAXPROCS(maxProces)
	cmd.Execute()
}
