package core

import (
	"fmt"
	"strings"

	"github.com/tilegame/tiledefs"
)

// ______________________________________________________________________
// Constants and Variables
// ======================================================================

const (
	GridSize = 30
)

var (
	nextId = 1336
)

const exampleMap = `
╔════════════════════════════╗
║,,,,,,,,,tttttttt,,,t,,,,,,,║
║,,,,t,,t,,,,,,,,,,t,,,,t,t,,║
║,,,,,ttt,,,,,,,ttttt,,,,t,,,║
║,,,,,ttt,,,,,,,,,tttt,,,~~,,║
║,,,t,,t,,t,t,,,t,ttt,,,~~~,,║
║,,,,t,,,,,,t,,t,,,,t,,,,,~~~║
║,,,tt,t,,,t,,,,,t,,,,t~,,,,,║
║,,,,,t,,,,,,,,,,,,,,~~~~,,,,║
║,,,t,,,,,t,,,,,t,,~~~~~~~,,,║
║,,,,,t,,,,,,,,,,,,,~~~~,,,,,║
║,,t,,,,,ttt,tt,,~~~~~,,,,,,,║
║,,,,,,,,,tttt,,,,,,~~~,,,,,,║
║,,,,tt,,,t,,,,t,,,,,,,,,,,,,║
║,,,,,,,,,,,,,,,,,,,,,,,,,,,,║
║,,t,,,,,,,,,,,,,,t,t,,,,,,,,║
║,,,,,,╔═══╗,,,,,,,,,,,t,,,,,║
║,,,,,,║,,,║,,,,,,,,,,,,,,,,,║
║,,,,,,║,,,╚══,,,,══╗,,t,,,,,║
║,,,,,,║,,,,,,,,,,,,║,,,,,,,,║
║,,,,,╔╝,,,,,,,,,,,,║,,,,,t,,║
║,,,,,║,,,,,~~~~~,,,,,,,,,,,,║
║,,,,,,,,,,,~~,,,~~~,,,,,,,,,║
║,,,,,,,,,,,~,,,,,,~~~,,,,,,,║
║,,,,,║,,,,,,,,,,,,,,,,,,,,,,║
║,,,,,╚═══,,,,,~~~,,,,,,,,,,,║
║,,,,,,t,,,,,,,~~~~,,,~~,,,,,║
║,,,,,,,,,,,,,~~~~,,,~~~,,,,,║
║,,,,t,,,,,,,,,t,,,,,,~~~,,,,║
╚════════════════════════════╝
`

func (g *Grid) loadMap(s string) {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Trim(s, "\n")
	row := 0
	col := 0
	for _, r := range s {
		if r == '\n' {
			row++
			col = 0
			continue
		}
		t, ok := tiledefs.DecodeSymbol(string(r))
		if !ok {
			e := fmt.Errorf("Unable to decode symbol: %s", string(r))
			panic(e)
		}
		g.tileMatrix[row][col] = t
		col++
	}
}

// ______________________________________________________________________
// Type Definitions
// ======================================================================

type Game struct {
	grid Grid
	ents EntMap
}

type Grid struct {
	tileMatrix [GridSize][GridSize]tiledefs.Tile
}

type Location struct {
	x, y int
}

type Entity struct {
	Kind    string
	Name    string
	Uid     int
	Pace    int
	Current Location
	Target  Location
	Health  int
}

type EntMap struct {
	data map[int]Entity
}

// ______________________________________________________________________
// Public Game Methods:  The Main API
// ======================================================================

func DefaultGame() *Game {
	return &Game{}
}

func (g *Game) NextTick() {
}

func (g *Game) AddEntity() {
}

func (g *Game) GetEntity() {
}

// ______________________________________________________________________
// Other Public Helper Functions
// ======================================================================

func Player(name string) *Entity {
	return &Entity{
		Kind:   "player",
		Name:   name,
		Uid:    nextUniqueId(),
		Pace:   300,
		Health: 100,
	}
}

// ______________________________________________________________________
// Private Helper Functions
// ======================================================================

func nextUniqueId() int {
	nextId++
	return nextId
}
