package main

import (
	"FishingSpotMarker/core"
)

func main() {
	// 初始化核心模块
	if err := core.Init(); err != nil {
		panic(err)
	}

	// 运行核心模块
	if err := core.Run(); err != nil {
		panic(err)
	}
}
