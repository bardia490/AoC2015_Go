package day6

import (
	"fmt"
	"testing"
)

func TestUpdateLights(t *testing.T) {
	lights := make(map[point]bool, 1000*1000)
	InitializeLights(&lights)

	UpdateLights(&lights, "turn on 0,0 through 999,999")
	if CountLitLights(lights) != 1000000 {
		fmt.Println("the first test failed")
	}

	UpdateLights(&lights, "toggle 0,0 through 999,0")
	if CountLitLights(lights) != 1000000 {
		fmt.Println("the second test failed")
	}
}
