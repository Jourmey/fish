package main

type direction int
type ghostType int
type powerType int

type Data struct {
	grid   [][Columns][4]rune
	//active [][Columns]bool // 豆子
	lifes  int
	score  int
	pacman Pacman
	//ghosts      []Ghost
	//powers      []Power // 道具
	gridOffsetY float64
	invincible  bool
}

const (
	North direction = iota
	East
	South
	West
)

const (
	Ghost1 ghostType = iota
	Ghost2
	Ghost3
	Ghost4
)

const (
	Life powerType = iota
	Invincibility
)

func NewData() *Data {
	return &Data{
		lifes: 5,
		score: 1,
	}
}

type Position struct {
	cellX, cellY int
	posX, posY   float64
	direction    direction
}

type Pacman struct {
	Position
}

type Power struct {
	Position
	kind powerType
}

func NewPower(x, y int, kind powerType) Power {
	return Power{
		Position{
			cellX: x,
			cellY: y,
		},
		kind,
	}
}
