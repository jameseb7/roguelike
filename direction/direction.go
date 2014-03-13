package direction

type Direction uint8

const (
	Here Direction = iota
	North
	South
	East
	West
	NorthEast
	NorthWest
	SouthWast
	SouthWest

	Up
	UpNorth
	UpSouth
	UpEast
	UpWest
	UpNorthEast
	UpNorthWest
	UpSouthEast
	UpSouthWest
	
	Down
	DownNorth
	DownSouth
	DownEast
	DownWest
	DownNorthEast
	DownNorthWest
	DownSouthEast
	DownSouthWest
	
	Special1
	Special2
	Special3

	NumDirections
)

type Triple struct{ X, Y, Z int }

var Directions = [...]Triple{
	Here:          {0, 0, 0},
	North:         {0, -1, 0},
	South:         {0, 1, 0},
	East:          {1, 0, 0},
	West:          {-1, 0, 0},
	NorthEast:     {1, -1, 0},
	NorthWest:     {-1, -1, 0},
	SouthEast:     {1, 1, 0},
	SouthWest:     {-1, 1, 0},
	Up:            {0, 0, 1},
	UpNorth:       {0, -1, 1},
	UpSouth:       {0, 1, 1},
	UpEast:        {1, 0, 1},
	UpWest:        {-1, 0, 1},
	UpNorthEast:   {1, -1, 1},
	UpNorthWest:   {-1, -1, 1},
	UpSouthEast:   {1, 1, 1},
	UpSouthWest:   {-1, 1, 1},
	Down:          {0, 0, -1},
	DownNorth:     {0, -1, -1},
	DownSouth:     {0, 1, -1},
	DownEast:      {1, 0, -1},
	DownWest:      {-1, 0, -1},
	DownNorthEast: {1, -1, -1},
	DownNorthWest: {-1, -1, -1},
	DownSouthEast: {1, 1, -1},
	DownSouthWest: {-1, 1, -1},
	Special1:      {0, 0, 0},
	Special2:      {0, 0, 0},
	Special3:      {0, 0, 0},
}