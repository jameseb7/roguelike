package types

type Direction uint8

const (
	HERE Direction = iota
	NORTH
	SOUTH
	EAST
	WEST

	NORTHEAST
	NORTHWEST
	SOUTHEAST
	SOUTHWEST

	UP
	UPNORTH
	UPSOUTH
	UPEAST
	UPWEST

	UPNORTHEAST
	UPNORTHWEST
	UPSOUTHEAST
	UPSOUTHWEST
	DOWN
	DOWNNORTH
	DOWNSOUTH
	DOWNEAST
	DOWNWEST
	DOWNNORTHEAST
	DOWNNORTHWEST
	DOWNSOUTHEAST
	DOWNSOUTHWEST
)

type Triple struct{ X, Y, Z int }

var Directions = [...]Triple{
	HERE:          {0, 0, 0},
	NORTH:         {0, -1, 0},
	SOUTH:         {0, 1, 0},
	EAST:          {1, 0, 0},
	WEST:          {-1, 0, 0},
	NORTHEAST:     {1, -1, 0},
	NORTHWEST:     {-1, -1, 0},
	SOUTHEAST:     {1, 1, 0},
	SOUTHWEST:     {-1, 1, 0},
	UP:            {0, 0, 1},
	UPNORTH:       {0, -1, 1},
	UPSOUTH:       {0, 1, 1},
	UPEAST:        {1, 0, 1},
	UPWEST:        {-1, 0, 1},
	UPNORTHEAST:   {1, -1, 1},
	UPNORTHWEST:   {-1, -1, 1},
	UPSOUTHEAST:   {1, 1, 1},
	UPSOUTHWEST:   {-1, 1, 1},
	DOWN:          {0, 0, -1},
	DOWNNORTH:     {0, -1, -1},
	DOWNSOUTH:     {0, 1, -1},
	DOWNEAST:      {1, 0, -1},
	DOWNWEST:      {-1, 0, -1},
	DOWNNORTHEAST: {1, -1, -1},
	DOWNNORTHWEST: {-1, -1, -1},
	DOWNSOUTHEAST: {1, 1, -1},
	DOWNSOUTHWEST: {-1, 1, -1},
}
