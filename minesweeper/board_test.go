package minesweeper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurroundingTiles(t *testing.T) {
	// tiles is a 3x3 representation of tiles
	tiles := [][]*Tile{
		{&Tile{row: 0, col: 0}, &Tile{row: 0, col: 1}, &Tile{row: 0, col: 2}},
		{&Tile{row: 1, col: 0}, &Tile{row: 1, col: 1}, &Tile{row: 1, col: 2}},
		{&Tile{row: 2, col: 0}, &Tile{row: 2, col: 1}, &Tile{row: 2, col: 2}},
	}

	testCases := []struct {
		name     string
		row, col int
		expected []*Tile
	}{
		{
			name: "top left",
			row:  0, col: 0,
			expected: []*Tile{
				// first row - 1
				/* {row: 0, col: 0},*/ {row: 0, col: 1}, /*{row: 0, col: 2},*/
				// second row - 0, 1
				{row: 1, col: 0}, {row: 1, col: 1}, /*{row: 1, col: 2},*/
				// third row - none
				/* {row: 2, col: 0}, {row: 2, col: 1}, {row: 2, col: 2}, */
			},
		},
		{
			name: "top center",
			row:  0, col: 1,
			expected: []*Tile{
				// first row - 0, 2
				{row: 0, col: 0} /*{row: 0, col: 1},*/, {row: 0, col: 2},
				// second row - all
				{row: 1, col: 0}, {row: 1, col: 1}, {row: 1, col: 2},
				// third row - none
				/* {row: 2, col: 0}, {row: 2, col: 1}, {row: 2, col: 2}, */
			},
		},
		{
			name: "top right",
			row:  0, col: 2,
			expected: []*Tile{
				// first row - 1
				/* {row: 0, col: 0}, */ {row: 0, col: 1}, /* {row: 0, col: 2}, */
				// second row - 1, 2
				/* {row: 1, col: 0}, */ {row: 1, col: 1}, {row: 1, col: 2},
				// third row - none
				/* {row: 2, col: 0}, {row: 2, col: 1}, {row: 2, col: 2}, */
			},
		},
		{
			name: "middle left",
			row:  1, col: 0,
			expected: []*Tile{
				// first row - 0, 1
				{row: 0, col: 0}, {row: 0, col: 1}, /* {row: 0, col: 2}, */
				// second row - 1
				/* {row: 1, col: 0}, */ {row: 1, col: 1}, /* {row: 1, col: 2}, */
				// third row - 0, 1
				{row: 2, col: 0}, {row: 2, col: 1}, /* {row: 2, col: 2}, */
			},
		},
		{
			name: "middle middle",
			row:  1, col: 1,
			expected: []*Tile{
				// first row - all
				{row: 0, col: 0}, {row: 0, col: 1}, {row: 0, col: 2},
				// second row - 0, 2
				{row: 1, col: 0} /* {row: 1, col: 1}, */, {row: 1, col: 2},
				// third row - all
				{row: 2, col: 0}, {row: 2, col: 1}, {row: 2, col: 2},
			},
		},
		{
			name: "middle right",
			row:  1, col: 2,
			expected: []*Tile{
				// first row - 1, 2
				/* {row: 0, col: 0}, */ {row: 0, col: 1}, {row: 0, col: 2},
				// second row - 1
				/* {row: 1, col: 0}, */ {row: 1, col: 1}, /* {row: 1, col: 2}, */
				// third row - 1, 2
				/* {row: 2, col: 0}, */ {row: 2, col: 1}, {row: 2, col: 2},
			},
		},
		{
			name: "bottom left",
			row:  2, col: 0,
			expected: []*Tile{
				// first row - none
				/* {row: 0, col: 0}, {row: 0, col: 1}, {row: 0, col: 2}, */
				// second row - 0, 1
				{row: 1, col: 0}, {row: 1, col: 1}, /* {row: 1, col: 2}, */
				// third row - 1
				/* {row: 2, col: 0}, */ {row: 2, col: 1}, /* {row: 2, col: 2}, */
			},
		},
		{
			name: "bottom middle",
			row:  2, col: 1,
			expected: []*Tile{
				// first row - none
				/* {row: 0, col: 0}, {row: 0, col: 1}, {row: 0, col: 2}, */
				// second row - 0, 1, 2
				{row: 1, col: 0}, {row: 1, col: 1}, {row: 1, col: 2},
				// third row - 0, 2
				{row: 2, col: 0} /* {row: 2, col: 1}, */, {row: 2, col: 2},
			},
		},
		{
			name: "bottom right",
			row:  2, col: 2,
			expected: []*Tile{
				// first row - none
				/* {row: 0, col: 0}, {row: 0, col: 1}, {row: 0, col: 2}, */
				// second row - 1, 2
				/* {row: 1, col: 0}, */ {row: 1, col: 1}, {row: 1, col: 2},
				// third row - 1
				/* {row: 2, col: 0}, */ {row: 2, col: 1}, /* {row: 2, col: 2}, */
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := surroundingTiles(tc.row, tc.col, tiles)

			for _, tile := range actual {
				fmt.Printf("%d:%d:%v\n", tile.row, tile.col, tile.bomb)
			}

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestClearZeroSurrounds(t *testing.T) {
	testCases := []struct {
		name        string
		surrounding []*Tile
		tiles       [][]*Tile
		expected    [][]*Tile
	}{
		// {
		// 	name: "single 0S surrounded",
		// 	surrounding: []*Tile{
		// 		{row: 1, col: 1},
		// 		{row: 1, col: 2},
		// 		{row: 1, col: 3},
		// 		{row: 2, col: 1},
		// 		{row: 2, col: 3},
		// 		{row: 3, col: 1},
		// 		{row: 3, col: 2},
		// 		{row: 3, col: 3},
		// 	},
		// 	tiles: [][]*Tile{
		// 		// | B | B | B | B | B |
		// 		// | B |   |   |   | B |
		// 		// | B |   | 0 |   | B |
		// 		// | B |   |   |   | B |
		// 		// | B | B | B | B | B |
		// 		{
		// 			&Tile{row: 0, col: 0, bomb: true},
		// 			&Tile{row: 0, col: 1, bomb: true},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3, bomb: true},
		// 			&Tile{row: 0, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0, bomb: true},
		// 			&Tile{row: 1, col: 1},
		// 			&Tile{row: 1, col: 2},
		// 			&Tile{row: 1, col: 3},
		// 			&Tile{row: 1, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0, bomb: true},
		// 			&Tile{row: 2, col: 1},
		// 			&Tile{row: 2, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 2, col: 3},
		// 			&Tile{row: 2, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0, bomb: true},
		// 			&Tile{row: 3, col: 1},
		// 			&Tile{row: 3, col: 2},
		// 			&Tile{row: 3, col: 3},
		// 			&Tile{row: 3, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0, bomb: true},
		// 			&Tile{row: 4, col: 1, bomb: true},
		// 			&Tile{row: 4, col: 2, bomb: true},
		// 			&Tile{row: 4, col: 3, bomb: true},
		// 			&Tile{row: 4, col: 4, bomb: true},
		// 		},
		// 	},
		// 	expected: [][]*Tile{
		// 		// | B | B | B | B | B |
		// 		// | B | 5 | 3 | 5 | B |
		// 		// | B | 3 | 0 | 3 | B |
		// 		// | B | 5 | 3 | 5 | B |
		// 		// | B | B | B | B | B |
		// 		{
		// 			&Tile{row: 0, col: 0, bomb: true},
		// 			&Tile{row: 0, col: 1, bomb: true},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3, bomb: true},
		// 			&Tile{row: 0, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0, bomb: true},
		// 			&Tile{row: 1, col: 1, surrounding: 5, state: cleared},
		// 			&Tile{row: 1, col: 2, surrounding: 3, state: cleared},
		// 			&Tile{row: 1, col: 3, surrounding: 5, state: cleared},
		// 			&Tile{row: 1, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0, bomb: true},
		// 			&Tile{row: 2, col: 1, surrounding: 3, state: cleared},
		// 			&Tile{row: 2, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 2, col: 3, surrounding: 3, state: cleared},
		// 			&Tile{row: 2, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0, bomb: true},
		// 			&Tile{row: 3, col: 1, surrounding: 5, state: cleared},
		// 			&Tile{row: 3, col: 2, surrounding: 3, state: cleared},
		// 			&Tile{row: 3, col: 3, surrounding: 5, state: cleared},
		// 			&Tile{row: 3, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0, bomb: true},
		// 			&Tile{row: 4, col: 1, bomb: true},
		// 			&Tile{row: 4, col: 2, bomb: true},
		// 			&Tile{row: 4, col: 3, bomb: true},
		// 			&Tile{row: 4, col: 4, bomb: true},
		// 		},
		// 	},
		// },
		// {
		// 	name: "single 0S in corner",
		// 	surrounding: []*Tile{
		// 		{row: 0, col: 1},
		// 		{row: 1, col: 0},
		// 		{row: 1, col: 1},
		// 		{row: 1, col: 2},
		// 	},
		// 	tiles: [][]*Tile{
		// 		// | 0 |   | B |
		// 		// |   |   |   |
		// 		// | B |   |   |
		// 		{
		// 			&Tile{row: 0, col: 0, surrounding: 0, state: cleared},
		// 			&Tile{row: 0, col: 1},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0},
		// 			&Tile{row: 1, col: 1},
		// 			&Tile{row: 1, col: 2},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0, bomb: true},
		// 			&Tile{row: 2, col: 1},
		// 			&Tile{row: 2, col: 2},
		// 		},
		// 	},
		// 	expected: [][]*Tile{
		// 		// | 0 |   | B |
		// 		// |   |   |   |
		// 		// | B |   |   |
		// 		{
		// 			&Tile{row: 0, col: 0, surrounding: 0, state: cleared},
		// 			&Tile{row: 0, col: 1, surrounding: 1, state: cleared},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0, surrounding: 1, state: cleared},
		// 			&Tile{row: 1, col: 1, surrounding: 2, state: cleared},
		// 			&Tile{row: 1, col: 2, surrounding: 1, state: cleared},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0, bomb: true},
		// 			&Tile{row: 2, col: 1},
		// 			&Tile{row: 2, col: 2},
		// 		},
		// 	},
		// },
		// {
		// 	name: "three 0S horizontal line left to right",
		// 	surrounding: []*Tile{
		// 		{row: 1, col: 1},
		// 		{row: 1, col: 2},
		// 		{row: 1, col: 3},
		// 		{row: 2, col: 1},
		// 		{row: 2, col: 3},
		// 		{row: 3, col: 1},
		// 		{row: 3, col: 2},
		// 		{row: 3, col: 3},
		// 	},
		// 	tiles: [][]*Tile{
		// 		// |   | B |  B  |     |  B  |
		// 		// |   |   |     |     |     |
		// 		// | B |   | 0#0 | 0#0 | 0#0 |
		// 		// |   |   |     |     |     |
		// 		// |   | B |     |     |  B  |
		// 		{
		// 			&Tile{row: 0, col: 0},
		// 			&Tile{row: 0, col: 1, bomb: true},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3},
		// 			&Tile{row: 0, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0},
		// 			&Tile{row: 1, col: 1},
		// 			&Tile{row: 1, col: 2},
		// 			&Tile{row: 1, col: 3},
		// 			&Tile{row: 1, col: 4},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0, bomb: true},
		// 			&Tile{row: 2, col: 1},
		// 			&Tile{row: 2, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 2, col: 3},
		// 			&Tile{row: 2, col: 4},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0},
		// 			&Tile{row: 3, col: 1},
		// 			&Tile{row: 3, col: 2},
		// 			&Tile{row: 3, col: 3},
		// 			&Tile{row: 3, col: 4},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0},
		// 			&Tile{row: 4, col: 1, bomb: true},
		// 			&Tile{row: 4, col: 2},
		// 			&Tile{row: 4, col: 3},
		// 			&Tile{row: 4, col: 4, bomb: true},
		// 		},
		// 	},
		// 	expected: [][]*Tile{
		// 		// |   | B |  B  |     |  B  |
		// 		// |   |   |     |     |     |
		// 		// | B |   | 0#0 | 0#0 | 0#0 |
		// 		// |   |   |     |     |     |
		// 		// |   | B |     |     |  B  |
		// 		{
		// 			&Tile{row: 0, col: 0},
		// 			&Tile{row: 0, col: 1, bomb: true},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3},
		// 			&Tile{row: 0, col: 4, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0},
		// 			&Tile{row: 1, col: 1, surrounding: 3, state: cleared},
		// 			&Tile{row: 1, col: 2, surrounding: 2, state: cleared},
		// 			&Tile{row: 1, col: 3, surrounding: 2, state: cleared},
		// 			&Tile{row: 1, col: 4, surrounding: 1, state: cleared},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0, bomb: true},
		// 			&Tile{row: 2, col: 1, surrounding: 1, state: cleared},
		// 			&Tile{row: 2, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 2, col: 3, surrounding: 0, state: cleared},
		// 			&Tile{row: 2, col: 4, surrounding: 0, state: cleared},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0},
		// 			&Tile{row: 3, col: 1, surrounding: 2, state: cleared},
		// 			&Tile{row: 3, col: 2, surrounding: 1, state: cleared},
		// 			&Tile{row: 3, col: 3, surrounding: 1, state: cleared},
		// 			&Tile{row: 3, col: 4, surrounding: 1, state: cleared},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0},
		// 			&Tile{row: 4, col: 1, bomb: true},
		// 			&Tile{row: 4, col: 2},
		// 			&Tile{row: 4, col: 3},
		// 			&Tile{row: 4, col: 4, bomb: true},
		// 		},
		// 	},
		// },
		// {
		// 	name: "many in different directions",
		// 	surrounding: []*Tile{
		// 		{row: 1, col: 1},
		// 		{row: 1, col: 2},
		// 		{row: 1, col: 3},
		// 		{row: 2, col: 1},
		// 		{row: 2, col: 3},
		// 		{row: 3, col: 1},
		// 		{row: 3, col: 2},
		// 		{row: 3, col: 3},
		// 	},
		// 	tiles: [][]*Tile{
		// 		// |     |     |  B  |     |     |     |     |
		// 		// |  B  |  s  |  s  |  s  |  B  |     |     |
		// 		// |     |  s  | 0#0 |  s  |  B  |     |  B  |
		// 		// |  B  |  s  | 0#1 |  s  |     |     |  B  |
		// 		// |     |     | 0#2 | 0#3 | 0#4 |     |  B  |
		// 		// |     | 0#5 | 0#6 |     |     |     |     |
		// 		// |     |     | 0#7 |     |  B  |     |     |
		// 		// |  B  |     |     |     |  B  |     |     |
		// 		{
		// 			&Tile{row: 0, col: 0},
		// 			&Tile{row: 0, col: 1},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3},
		// 			&Tile{row: 0, col: 4},
		// 			&Tile{row: 0, col: 5},
		// 			&Tile{row: 0, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0, bomb: true},
		// 			&Tile{row: 1, col: 1},
		// 			&Tile{row: 1, col: 2},
		// 			&Tile{row: 1, col: 3},
		// 			&Tile{row: 1, col: 4, bomb: true},
		// 			&Tile{row: 1, col: 5},
		// 			&Tile{row: 1, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0},
		// 			&Tile{row: 2, col: 1},
		// 			&Tile{row: 2, col: 2, surrounding: 0},
		// 			&Tile{row: 2, col: 3},
		// 			&Tile{row: 2, col: 4, bomb: true},
		// 			&Tile{row: 2, col: 5},
		// 			&Tile{row: 2, col: 6, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0, bomb: true},
		// 			&Tile{row: 3, col: 1},
		// 			&Tile{row: 3, col: 2},
		// 			&Tile{row: 3, col: 3},
		// 			&Tile{row: 3, col: 4},
		// 			&Tile{row: 3, col: 5},
		// 			&Tile{row: 3, col: 6, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0},
		// 			&Tile{row: 4, col: 1},
		// 			&Tile{row: 4, col: 2},
		// 			&Tile{row: 4, col: 3},
		// 			&Tile{row: 4, col: 4},
		// 			&Tile{row: 4, col: 5},
		// 			&Tile{row: 4, col: 6, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 5, col: 0},
		// 			&Tile{row: 5, col: 1},
		// 			&Tile{row: 5, col: 2},
		// 			&Tile{row: 5, col: 3},
		// 			&Tile{row: 5, col: 4},
		// 			&Tile{row: 5, col: 5},
		// 			&Tile{row: 5, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 6, col: 0},
		// 			&Tile{row: 6, col: 1},
		// 			&Tile{row: 6, col: 2},
		// 			&Tile{row: 6, col: 3},
		// 			&Tile{row: 6, col: 4, bomb: true},
		// 			&Tile{row: 6, col: 5},
		// 			&Tile{row: 6, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 7, col: 0, bomb: true},
		// 			&Tile{row: 7, col: 1},
		// 			&Tile{row: 7, col: 2},
		// 			&Tile{row: 7, col: 3},
		// 			&Tile{row: 7, col: 4, bomb: true},
		// 			&Tile{row: 7, col: 5},
		// 			&Tile{row: 7, col: 6},
		// 		},
		// 	},
		// 	expected: [][]*Tile{
		// 		// |     |     |  B  |     |     |     |     | *
		// 		// |  B  |  2  |  1  |  3  |  B  |     |     | *
		// 		// |     |  2  | 0#0 |  2  |  B  |     |  B  | *
		// 		// |  B  |  1  | 0#1 |  1  |  1  |  4  |  B  | *
		// 		// |  1  |  1  | 0#2 | 0#3 | 0#4 |  2  |  B  | *
		// 		// | 0#5 | 0#6 | 0#7 |  1  |  1  |  2  |     | *
		// 		// |  1  |  1  | 0#8 |  2  |  B  |     |     |
		// 		// |  B  |  1  | 0#9 |  2  |  B  |     |     |
		// 		{
		// 			&Tile{row: 0, col: 0},
		// 			&Tile{row: 0, col: 1},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3},
		// 			&Tile{row: 0, col: 4},
		// 			&Tile{row: 0, col: 5},
		// 			&Tile{row: 0, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0, bomb: true},
		// 			&Tile{row: 1, col: 1, surrounding: 2, state: cleared},
		// 			&Tile{row: 1, col: 2, surrounding: 1, state: cleared},
		// 			&Tile{row: 1, col: 3, surrounding: 3, state: cleared},
		// 			&Tile{row: 1, col: 4, bomb: true},
		// 			&Tile{row: 1, col: 5},
		// 			&Tile{row: 1, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0},
		// 			&Tile{row: 2, col: 1, surrounding: 2, state: cleared},
		// 			&Tile{row: 2, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 2, col: 3, surrounding: 2, state: cleared},
		// 			&Tile{row: 2, col: 4, bomb: true},
		// 			&Tile{row: 2, col: 5},
		// 			&Tile{row: 2, col: 6, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0, bomb: true},
		// 			&Tile{row: 3, col: 1, surrounding: 1, state: cleared},
		// 			&Tile{row: 3, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 3, col: 3, surrounding: 1, state: cleared},
		// 			&Tile{row: 3, col: 4, surrounding: 1, state: cleared},
		// 			&Tile{row: 3, col: 5, surrounding: 4, state: cleared},
		// 			&Tile{row: 3, col: 6, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0, surrounding: 1, state: cleared},
		// 			&Tile{row: 4, col: 1, surrounding: 1, state: cleared},
		// 			&Tile{row: 4, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 4, col: 3, surrounding: 0, state: cleared},
		// 			&Tile{row: 4, col: 4, surrounding: 0, state: cleared},
		// 			&Tile{row: 4, col: 5, surrounding: 2, state: cleared},
		// 			&Tile{row: 4, col: 6, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 5, col: 0, surrounding: 0, state: cleared},
		// 			&Tile{row: 5, col: 1, surrounding: 0, state: cleared},
		// 			&Tile{row: 5, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 5, col: 3, surrounding: 1, state: cleared},
		// 			&Tile{row: 5, col: 4, surrounding: 1, state: cleared},
		// 			&Tile{row: 5, col: 5, surrounding: 2, state: cleared},
		// 			&Tile{row: 5, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 6, col: 0, surrounding: 1, state: cleared},
		// 			&Tile{row: 6, col: 1, surrounding: 1, state: cleared},
		// 			&Tile{row: 6, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 6, col: 3, surrounding: 2, state: cleared},
		// 			&Tile{row: 6, col: 4, bomb: true},
		// 			&Tile{row: 6, col: 5},
		// 			&Tile{row: 6, col: 6},
		// 		},
		// 		{
		// 			&Tile{row: 7, col: 0, bomb: true},
		// 			&Tile{row: 7, col: 1, surrounding: 1, state: cleared},
		// 			&Tile{row: 7, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 7, col: 3, surrounding: 2, state: cleared},
		// 			&Tile{row: 7, col: 4, bomb: true},
		// 			&Tile{row: 7, col: 5},
		// 			&Tile{row: 7, col: 6},
		// 		},
		// 	},
		// },
		// {
		// 	name: "only consider adjacent tiles for new 0S",
		// 	surrounding: []*Tile{
		// 		{row: 1, col: 1},
		// 		{row: 1, col: 2},
		// 		{row: 1, col: 3},
		// 		{row: 2, col: 1},
		// 		{row: 2, col: 3},
		// 		{row: 3, col: 1},
		// 		{row: 3, col: 2},
		// 		{row: 3, col: 3},
		// 	},
		// 	tiles: [][]*Tile{
		// 		// |     |     |  B  |     |     |     | *
		// 		// |  B  |  s  |  s  |  s  |  B  |     | *
		// 		// |     |  s  | 0#0 |  s  |     |     | *
		// 		// |  B  |  s  |  s  | 0#1 |     |  B  | *
		// 		// |     |  B  |     |     |     |     | *
		// 		// |     |     |     |  B  |     |     | *
		// 		{
		// 			&Tile{row: 0, col: 0},
		// 			&Tile{row: 0, col: 1},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3},
		// 			&Tile{row: 0, col: 4},
		// 			&Tile{row: 0, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0, bomb: true},
		// 			&Tile{row: 1, col: 1},
		// 			&Tile{row: 1, col: 2},
		// 			&Tile{row: 1, col: 3},
		// 			&Tile{row: 1, col: 4, bomb: true},
		// 			&Tile{row: 1, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0},
		// 			&Tile{row: 2, col: 1},
		// 			&Tile{row: 2, col: 2, surrounding: 0},
		// 			&Tile{row: 2, col: 3},
		// 			&Tile{row: 2, col: 4},
		// 			&Tile{row: 2, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0, bomb: true},
		// 			&Tile{row: 3, col: 1},
		// 			&Tile{row: 3, col: 2},
		// 			&Tile{row: 3, col: 3},
		// 			&Tile{row: 3, col: 4},
		// 			&Tile{row: 3, col: 5, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0},
		// 			&Tile{row: 4, col: 1, bomb: true},
		// 			&Tile{row: 4, col: 2},
		// 			&Tile{row: 4, col: 3},
		// 			&Tile{row: 4, col: 4},
		// 			&Tile{row: 4, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 5, col: 0},
		// 			&Tile{row: 5, col: 1},
		// 			&Tile{row: 5, col: 2},
		// 			&Tile{row: 5, col: 3, bomb: true},
		// 			&Tile{row: 5, col: 4},
		// 			&Tile{row: 5, col: 5},
		// 		},
		// 	},
		// 	expected: [][]*Tile{
		// 		// |     |     |  B  |     |     |     | *
		// 		// |  B  |  2  |  1  |  2  |  B  |     | *
		// 		// |     |  2  | 0#0 |  1  |  2  |     | *
		// 		// |  B  |  2  |  1  | 0#1 |  1  |  B  | *
		// 		// |     |  B  |  2  |  1  |  2  |     | *
		// 		// |     |     |     |  B  |     |     | *
		// 		{
		// 			&Tile{row: 0, col: 0},
		// 			&Tile{row: 0, col: 1},
		// 			&Tile{row: 0, col: 2, bomb: true},
		// 			&Tile{row: 0, col: 3},
		// 			&Tile{row: 0, col: 4},
		// 			&Tile{row: 0, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 1, col: 0, bomb: true},
		// 			&Tile{row: 1, col: 1, surrounding: 2, state: cleared},
		// 			&Tile{row: 1, col: 2, surrounding: 1, state: cleared},
		// 			&Tile{row: 1, col: 3, surrounding: 2, state: cleared},
		// 			&Tile{row: 1, col: 4, bomb: true},
		// 			&Tile{row: 1, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 2, col: 0},
		// 			&Tile{row: 2, col: 1, surrounding: 2, state: cleared},
		// 			&Tile{row: 2, col: 2, surrounding: 0, state: cleared},
		// 			&Tile{row: 2, col: 3, surrounding: 1, state: cleared},
		// 			&Tile{row: 2, col: 4, surrounding: 2, state: cleared},
		// 			&Tile{row: 2, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 3, col: 0, bomb: true},
		// 			&Tile{row: 3, col: 1, surrounding: 2, state: cleared},
		// 			&Tile{row: 3, col: 2, surrounding: 1, state: cleared},
		// 			&Tile{row: 3, col: 3, surrounding: 0, state: cleared},
		// 			&Tile{row: 3, col: 4, surrounding: 1, state: cleared},
		// 			&Tile{row: 3, col: 5, bomb: true},
		// 		},
		// 		{
		// 			&Tile{row: 4, col: 0},
		// 			&Tile{row: 4, col: 1, bomb: true},
		// 			&Tile{row: 4, col: 2, surrounding: 2, state: cleared},
		// 			&Tile{row: 4, col: 3, surrounding: 1, state: cleared},
		// 			&Tile{row: 4, col: 4, surrounding: 2, state: cleared},
		// 			&Tile{row: 4, col: 5},
		// 		},
		// 		{
		// 			&Tile{row: 5, col: 0},
		// 			&Tile{row: 5, col: 1},
		// 			&Tile{row: 5, col: 2},
		// 			&Tile{row: 5, col: 3, bomb: true},
		// 			&Tile{row: 5, col: 4},
		// 			&Tile{row: 5, col: 5},
		// 		},
		// 	},
		// },
		{
			name: "weird observation",
			surrounding: []*Tile{
				{row: 1, col: 1},
				{row: 1, col: 2},
				{row: 1, col: 3},
				{row: 2, col: 1},
				{row: 2, col: 3},
				{row: 3, col: 1},
				{row: 3, col: 2},
				{row: 3, col: 3},
			},
			tiles: [][]*Tile{
				// |  1  |  1  |  2  |  B  |  C  |     |  1  |
				// |  1  |  B  |  2  |  1  |  2  |  2  |  2  |
				// |  1  |  1  |  1  |     |     |  1  |     |
				// |     |     |     |     |     |  2  |     |
				{
					&Tile{row: 0, col: 0, surrounding: 1, state: cleared},
					&Tile{row: 0, col: 1, surrounding: 1, state: cleared},
					&Tile{row: 0, col: 2, surrounding: 2, state: cleared},
					&Tile{row: 0, col: 3, bomb: true},
					&Tile{row: 0, col: 4, surrounding: 2, state: cleared},
					&Tile{row: 0, col: 5},
				},
				{
					&Tile{row: 1, col: 0, bomb: true},
					&Tile{row: 1, col: 1},
					&Tile{row: 1, col: 2},
					&Tile{row: 1, col: 3},
					&Tile{row: 1, col: 4, bomb: true},
					&Tile{row: 1, col: 5},
				},
				{
					&Tile{row: 2, col: 0},
					&Tile{row: 2, col: 1},
					&Tile{row: 2, col: 2, surrounding: 0},
					&Tile{row: 2, col: 3},
					&Tile{row: 2, col: 4},
					&Tile{row: 2, col: 5},
				},
				{
					&Tile{row: 3, col: 0, bomb: true},
					&Tile{row: 3, col: 1},
					&Tile{row: 3, col: 2},
					&Tile{row: 3, col: 3},
					&Tile{row: 3, col: 4},
					&Tile{row: 3, col: 5, bomb: true},
				},
				{
					&Tile{row: 4, col: 0},
					&Tile{row: 4, col: 1, bomb: true},
					&Tile{row: 4, col: 2},
					&Tile{row: 4, col: 3},
					&Tile{row: 4, col: 4},
					&Tile{row: 4, col: 5},
				},
				{
					&Tile{row: 5, col: 0},
					&Tile{row: 5, col: 1},
					&Tile{row: 5, col: 2},
					&Tile{row: 5, col: 3, bomb: true},
					&Tile{row: 5, col: 4},
					&Tile{row: 5, col: 5},
				},
			},
			expected: [][]*Tile{
				// |     |     |  B  |     |     |     | *
				// |  B  |  2  |  1  |  2  |  B  |     | *
				// |     |  2  | 0#0 |  1  |  2  |     | *
				// |  B  |  2  |  1  | 0#1 |  1  |  B  | *
				// |     |  B  |  2  |  1  |  2  |     | *
				// |     |     |     |  B  |     |     | *
				{
					&Tile{row: 0, col: 0},
					&Tile{row: 0, col: 1},
					&Tile{row: 0, col: 2, bomb: true},
					&Tile{row: 0, col: 3},
					&Tile{row: 0, col: 4},
					&Tile{row: 0, col: 5},
				},
				{
					&Tile{row: 1, col: 0, bomb: true},
					&Tile{row: 1, col: 1, surrounding: 2, state: cleared},
					&Tile{row: 1, col: 2, surrounding: 1, state: cleared},
					&Tile{row: 1, col: 3, surrounding: 2, state: cleared},
					&Tile{row: 1, col: 4, bomb: true},
					&Tile{row: 1, col: 5},
				},
				{
					&Tile{row: 2, col: 0},
					&Tile{row: 2, col: 1, surrounding: 2, state: cleared},
					&Tile{row: 2, col: 2, surrounding: 0, state: cleared},
					&Tile{row: 2, col: 3, surrounding: 1, state: cleared},
					&Tile{row: 2, col: 4, surrounding: 2, state: cleared},
					&Tile{row: 2, col: 5},
				},
				{
					&Tile{row: 3, col: 0, bomb: true},
					&Tile{row: 3, col: 1, surrounding: 2, state: cleared},
					&Tile{row: 3, col: 2, surrounding: 1, state: cleared},
					&Tile{row: 3, col: 3, surrounding: 0, state: cleared},
					&Tile{row: 3, col: 4, surrounding: 1, state: cleared},
					&Tile{row: 3, col: 5, bomb: true},
				},
				{
					&Tile{row: 4, col: 0},
					&Tile{row: 4, col: 1, bomb: true},
					&Tile{row: 4, col: 2, surrounding: 2, state: cleared},
					&Tile{row: 4, col: 3, surrounding: 1, state: cleared},
					&Tile{row: 4, col: 4, surrounding: 2, state: cleared},
					&Tile{row: 4, col: 5},
				},
				{
					&Tile{row: 5, col: 0},
					&Tile{row: 5, col: 1},
					&Tile{row: 5, col: 2},
					&Tile{row: 5, col: 3, bomb: true},
					&Tile{row: 5, col: 4},
					&Tile{row: 5, col: 5},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clearZeroSurrounds(tc.surrounding, tc.tiles)

			assert.Equal(t, tc.expected, tc.tiles)
		})
	}
}
