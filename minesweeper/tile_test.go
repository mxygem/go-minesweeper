package minesweeper

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTiles(t *testing.T) {
	testCases := []struct {
		name                    string
		width, height, tileSize int
		expected                [][]*Tile
	}{
		{
			name:     "50 hw 10 t",
			width:    50,
			height:   50,
			tileSize: 10,
			expected: [][]*Tile{
				{&Tile{row: 0, col: 0}, &Tile{row: 0, col: 1}, &Tile{row: 0, col: 2}, &Tile{row: 0, col: 3}, &Tile{row: 0, col: 4}},
				{&Tile{row: 1, col: 0}, &Tile{row: 1, col: 1}, &Tile{row: 1, col: 2}, &Tile{row: 1, col: 3}, &Tile{row: 1, col: 4}},
				{&Tile{row: 2, col: 0}, &Tile{row: 2, col: 1}, &Tile{row: 2, col: 2}, &Tile{row: 2, col: 3}, &Tile{row: 2, col: 4}},
				{&Tile{row: 3, col: 0}, &Tile{row: 3, col: 1}, &Tile{row: 3, col: 2}, &Tile{row: 3, col: 3}, &Tile{row: 3, col: 4}},
				{&Tile{row: 4, col: 0}, &Tile{row: 4, col: 1}, &Tile{row: 4, col: 2}, &Tile{row: 4, col: 3}, &Tile{row: 4, col: 4}},
			},
		},
		{
			name:     "round down when not equally divisible",
			width:    31,
			height:   31,
			tileSize: 10,
			expected: [][]*Tile{
				{&Tile{row: 0, col: 0}, &Tile{row: 0, col: 1}, &Tile{row: 0, col: 2}},
				{&Tile{row: 1, col: 0}, &Tile{row: 1, col: 1}, &Tile{row: 1, col: 2}},
				{&Tile{row: 2, col: 0}, &Tile{row: 2, col: 1}, &Tile{row: 2, col: 2}},
			},
		},
		{
			name:     "round down when not equally divisible",
			width:    39,
			height:   39,
			tileSize: 10,
			expected: [][]*Tile{
				{&Tile{row: 0, col: 0}, &Tile{row: 0, col: 1}, &Tile{row: 0, col: 2}},
				{&Tile{row: 1, col: 0}, &Tile{row: 1, col: 1}, &Tile{row: 1, col: 2}},
				{&Tile{row: 2, col: 0}, &Tile{row: 2, col: 1}, &Tile{row: 2, col: 2}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tiles(tc.width, tc.height, tc.tileSize)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPlaceBombs(t *testing.T) {
	testCases := []struct {
		name       string
		difficulty int
		tiles      [][]*Tile
		expected   [][]*Tile
	}{
		{
			name:       "5x5 easy",
			difficulty: 0,
			tiles:      tiles(50, 50, 10),
			expected: [][]*Tile{
				{&Tile{row: 0, col: 0, bomb: true}, &Tile{row: 0, col: 1, bomb: true}, &Tile{row: 0, col: 2}, &Tile{row: 0, col: 3}, &Tile{row: 0, col: 4, bomb: true}},
				{&Tile{row: 1, col: 0, bomb: true}, &Tile{row: 1, col: 1}, &Tile{row: 1, col: 2}, &Tile{row: 1, col: 3}, &Tile{row: 1, col: 4}},
				{&Tile{row: 2, col: 0}, &Tile{row: 2, col: 1}, &Tile{row: 2, col: 2}, &Tile{row: 2, col: 3}, &Tile{row: 2, col: 4, bomb: true}},
				{&Tile{row: 3, col: 0}, &Tile{row: 3, col: 1}, &Tile{row: 3, col: 2}, &Tile{row: 3, col: 3, bomb: true}, &Tile{row: 3, col: 4}},
				{&Tile{row: 4, col: 0, bomb: true}, &Tile{row: 4, col: 1}, &Tile{row: 4, col: 2}, &Tile{row: 4, col: 3}, &Tile{row: 4, col: 4}},
			},
		},
		{
			name:       "5x5 medium",
			difficulty: 1,
			tiles:      tiles(50, 50, 10),
			expected: [][]*Tile{
				{&Tile{row: 0, col: 0, bomb: true}, &Tile{row: 0, col: 1, bomb: true}, &Tile{row: 0, col: 2}, &Tile{row: 0, col: 3}, &Tile{row: 0, col: 4, bomb: true}},
				{&Tile{row: 1, col: 0, bomb: true}, &Tile{row: 1, col: 1, bomb: true}, &Tile{row: 1, col: 2}, &Tile{row: 1, col: 3}, &Tile{row: 1, col: 4}},
				{&Tile{row: 2, col: 0}, &Tile{row: 2, col: 1, bomb: true}, &Tile{row: 2, col: 2}, &Tile{row: 2, col: 3}, &Tile{row: 2, col: 4, bomb: true}},
				{&Tile{row: 3, col: 0}, &Tile{row: 3, col: 1}, &Tile{row: 3, col: 2}, &Tile{row: 3, col: 3, bomb: true}, &Tile{row: 3, col: 4}},
				{&Tile{row: 4, col: 0, bomb: true}, &Tile{row: 4, col: 1}, &Tile{row: 4, col: 2}, &Tile{row: 4, col: 3, bomb: true}, &Tile{row: 4, col: 4, bomb: true}},
			},
		},
		{
			name:       "10x10 easy",
			difficulty: 0,
			tiles:      tiles(100, 100, 10),
			expected: [][]*Tile{
				{&Tile{row: 0, col: 0}, &Tile{row: 0, col: 1, bomb: true}, &Tile{row: 0, col: 2}, &Tile{row: 0, col: 3, bomb: true}, &Tile{row: 0, col: 4}, &Tile{row: 0, col: 5, bomb: true}, &Tile{row: 0, col: 6}, &Tile{row: 0, col: 7}, &Tile{row: 0, col: 8, bomb: true}, &Tile{row: 0, col: 9}},
				{&Tile{row: 1, col: 0}, &Tile{row: 1, col: 1}, &Tile{row: 1, col: 2}, &Tile{row: 1, col: 3}, &Tile{row: 1, col: 4}, &Tile{row: 1, col: 5, bomb: true}, &Tile{row: 1, col: 6}, &Tile{row: 1, col: 7}, &Tile{row: 1, col: 8}, &Tile{row: 1, col: 9, bomb: true}},
				{&Tile{row: 2, col: 0, bomb: true}, &Tile{row: 2, col: 1, bomb: true}, &Tile{row: 2, col: 2, bomb: true}, &Tile{row: 2, col: 3}, &Tile{row: 2, col: 4}, &Tile{row: 2, col: 5}, &Tile{row: 2, col: 6}, &Tile{row: 2, col: 7}, &Tile{row: 2, col: 8}, &Tile{row: 2, col: 9}},
				{&Tile{row: 3, col: 0}, &Tile{row: 3, col: 1}, &Tile{row: 3, col: 2, bomb: true}, &Tile{row: 3, col: 3, bomb: true}, &Tile{row: 3, col: 4}, &Tile{row: 3, col: 5}, &Tile{row: 3, col: 6, bomb: true}, &Tile{row: 3, col: 7}, &Tile{row: 3, col: 8, bomb: true}, &Tile{row: 3, col: 9}},
				{&Tile{row: 4, col: 0, bomb: true}, &Tile{row: 4, col: 1}, &Tile{row: 4, col: 2, bomb: true}, &Tile{row: 4, col: 3}, &Tile{row: 4, col: 4}, &Tile{row: 4, col: 5, bomb: true}, &Tile{row: 4, col: 6}, &Tile{row: 4, col: 7}, &Tile{row: 4, col: 8}, &Tile{row: 4, col: 9, bomb: true}},
				{&Tile{row: 5, col: 0}, &Tile{row: 5, col: 1}, &Tile{row: 5, col: 2}, &Tile{row: 5, col: 3}, &Tile{row: 5, col: 4}, &Tile{row: 5, col: 5}, &Tile{row: 5, col: 6, bomb: true}, &Tile{row: 5, col: 7}, &Tile{row: 5, col: 8}, &Tile{row: 5, col: 9, bomb: true}},
				{&Tile{row: 6, col: 0, bomb: true}, &Tile{row: 6, col: 1, bomb: true}, &Tile{row: 6, col: 2}, &Tile{row: 6, col: 3}, &Tile{row: 6, col: 4, bomb: true}, &Tile{row: 6, col: 5}, &Tile{row: 6, col: 6, bomb: true}, &Tile{row: 6, col: 7}, &Tile{row: 6, col: 8}, &Tile{row: 6, col: 9}},
				{&Tile{row: 7, col: 0, bomb: true}, &Tile{row: 7, col: 1}, &Tile{row: 7, col: 2, bomb: true}, &Tile{row: 7, col: 3}, &Tile{row: 7, col: 4}, &Tile{row: 7, col: 5}, &Tile{row: 7, col: 6}, &Tile{row: 7, col: 7}, &Tile{row: 7, col: 8}, &Tile{row: 7, col: 9, bomb: true}},
				{&Tile{row: 8, col: 0}, &Tile{row: 8, col: 1}, &Tile{row: 8, col: 2}, &Tile{row: 8, col: 3}, &Tile{row: 8, col: 4}, &Tile{row: 8, col: 5}, &Tile{row: 8, col: 6}, &Tile{row: 8, col: 7, bomb: true}, &Tile{row: 8, col: 8}, &Tile{row: 8, col: 9}},
				{&Tile{row: 9, col: 0}, &Tile{row: 9, col: 1}, &Tile{row: 9, col: 2, bomb: true}, &Tile{row: 9, col: 3, bomb: true}, &Tile{row: 9, col: 4}, &Tile{row: 9, col: 5, bomb: true}, &Tile{row: 9, col: 6}, &Tile{row: 9, col: 7}, &Tile{row: 9, col: 8}, &Tile{row: 9, col: 9}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := rand.New(rand.NewSource(2113))

			placeBombs(tc.difficulty, tc.tiles, r)

			if !assert.Equal(t, tc.expected, tc.tiles) {
				for i, r := range tc.tiles {
					fmt.Printf("r%d: %v\n", i, r)
				}
			}
		})
	}
}
