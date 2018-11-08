package core

import (
	"testing"
)

func nop(interface{}) {}

func TestMain(t *testing.T) {
	game := Game{}
	nop(game)
}

func TestGrid(t *testing.T) {
	game := Game{}
	game.grid.loadMap(exampleMap)
	t.Log("map loaded, printing...")
	for _, line := range game.grid.tileMatrix {
		t.Log(line)
	}
}
