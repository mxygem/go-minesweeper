package minesweeper

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	width, height int
	tiles         [][]*Tile
	difficulty    Difficulty
}

// NewBoard returns a new Board with tiles populated base on the provided height, width, and
// difficulty.
func NewBoard(width, height int, diff Difficulty) *Board {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	b := &Board{width: width, height: height}
	b.tiles = tiles(width, height, tileSize)

	placeBombs(diff, b.tiles, r)

	return b
}

func (b *Board) Size() (int, int) {
	x := b.width*tileSize + (b.width + 1)
	y := b.height*tileSize + (b.height + 1)
	return x, y
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(cBlack)

	for _, row := range b.tiles {
		for _, t := range row {
			t.Draw(boardImage)
		}
	}
}

func (b *Board) Update(input *Input) {
	if input.mouseState != mouseStateSettled {
		return
	}

	// try to find match base on mouse location and tile areas
	match := maybeMatchTile(input.mousePosX, input.mousePosY, b.tiles)
	if match == nil {
		return
	}

	// update tile's state based on type of input and game mechanics
	// get surrounding tiles
	surrounding := surroundingTiles(match.row, match.col, b.tiles)
	match.Update(input.mouseButton, surrounding)

	// if a bomb was clicked, explode all of them
	if match.state == explode {
		explodeAll(b.tiles)
	}

	// if a tile with no surrounding bombs is clicked, attempt to clear any other adjacent tiles
	// with no surrounding bombs and clear the border tiles round them.
	if match.surrounding == 0 && !match.bomb {
		clearZeroSurrounds(surrounding, b.tiles)
	}
}

func maybeMatchTile(inputX, inputY int, tiles [][]*Tile) *Tile {
	for _, row := range tiles {
		for _, t := range row {
			x := inputX
			y := inputY
			tr := t.Rect
			if tr.Min.X <= x && x < tr.Max.X && tr.Min.Y <= y && y < tr.Max.Y {
				return t
			}
		}
	}

	return nil
}

func surroundingTiles(row, col int, tiles [][]*Tile) []*Tile {
	if len(tiles) == 0 || len(tiles[0]) == 0 {
		return nil
	}

	rowStart := row - 1
	if rowStart < 0 {
		rowStart = 0
	}
	rowEnd := row + 2
	if rowEnd > len(tiles) {
		rowEnd = len(tiles)
	}
	colStart := col - 1
	if colStart < 0 {
		colStart = 0
	}
	colEnd := col + 2
	if colEnd > len(tiles[0]) {
		colEnd = len(tiles[0])
	}

	tilesOut := make([]*Tile, 0, 8)
	for _, subRow := range tiles[rowStart:rowEnd] {
		for _, t := range subRow[colStart:colEnd] {
			if row == t.row && col == t.col {
				continue
			}
			tilesOut = append(tilesOut, t)
		}
	}

	// fmt.Printf("tile row: %d col: %d\n", row, col)
	// fmt.Printf("returning %d surrounding tiles\n", len(tilesOut))
	// for i, t := range tilesOut {
	// 	fmt.Printf("%d: row: %d col: %d bomb: %v, state: %d, surrounding: %d\n", i, t.row, t.col, t.bomb, t.state, t.surrounding)
	// }

	return tilesOut
}

func clearZeroSurrounds(surrounding []*Tile, tiles [][]*Tile) {
	zeros := make([]*Tile, 0, len(surrounding))

	for _, st := range surrounding {
		tile := tiles[st.row][st.col]
		subSurround := surroundingTiles(tile.row, tile.col, tiles)
		tile.Update(mouseButtonLeft, subSurround)

		if tile.surrounding == 0 && tile.state != cleared && !tile.bomb {
			zeros = append(zeros, tile)
		}
		tile.state = cleared
	}

	if len(zeros) > 0 {
		for _, z := range zeros {
			zSurrounding := surroundingTiles(z.row, z.col, tiles)
			clearZeroSurrounds(zSurrounding, tiles)
		}
	}
}
