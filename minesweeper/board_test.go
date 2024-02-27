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
