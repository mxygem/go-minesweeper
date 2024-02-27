package minesweeper

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed FiraSans-Regular.ttf
var firaSansRegular []byte
var firaSansFaceSource *text.GoTextFaceSource

const (
	tileSize = 20
	fontSize = 20
)

var (
	diffFill  = map[int]float64{0: 0.25, 1: 0.45, 2: 0.6}
	tileImage = ebiten.NewImage(tileSize-2, tileSize-2)
)

type tileState int

const (
	base tileState = iota
	marked
	cleared
	explode
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(firaSansRegular))
	if err != nil {
		log.Fatal(err)
	}
	firaSansFaceSource = s
	tileImage.Fill(cGrey50)
}

type Tile struct {
	row, col    int
	X, Y        int
	Size        int
	Rect        image.Rectangle
	bomb        bool
	state       tileState
	surrounding int
}

func (t *Tile) Update(btn mouseButton, surrounding []*Tile) {
	switch btn {
	case mouseButtonLeft:
		if t.state == marked || t.state == cleared {
			return
		}

		if t.bomb {
			t.state = explode
			return
		}

		var bombCount int
		for _, s := range surrounding {
			if !s.bomb {
				continue
			}

			bombCount++
		}

		t.surrounding = bombCount
		t.state = cleared
	case mouseButtonRight:
		switch t.state {
		case base:
			t.state++
		case marked:
			t.state--
		}
	}
}

// Draw draws the current tile to the given boardImage.
func (t *Tile) Draw(boardImage *ebiten.Image) {
	// clicked
	// unsolved
	// cleared
	// 	* number
	// 	* blank
	// explosion (red X for now)
	// if t.explode {
	// 	fmt.Println("found explode")
	// 	tileImage.Fill(cRed)
	// }

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.X)+1, float64(t.Y)+1)

	switch t.state {
	case base:
		op.ColorScale.ScaleWithColor(cGrey80)
	case cleared:
		if t.surrounding == 0 {
			op.ColorScale.ScaleWithColor(cWhite)
		} else {
			op.ColorScale.ScaleWithColor(cGreen)
		}
	case explode:
		op.ColorScale.ScaleWithColor(cRed)
	}

	boardImage.DrawImage(tileImage, op)

	if t.state == base || t.state == explode {
		return
	}

	top := &text.DrawOptions{}
	top.GeoM.Translate(
		float64(t.X+5),
		float64(t.Y-1),
	)

	var msg string
	if t.state == marked {
		top.ColorScale.ScaleWithColor(cRed)
		msg = "X"
	}
	if t.surrounding > 0 {
		top.ColorScale.ScaleWithColor(color.White)
		msg = strconv.Itoa(t.surrounding)
	}

	if msg != "" {
		text.Draw(boardImage, msg, &text.GoTextFace{
			Source: firaSansFaceSource,
			Size:   fontSize,
		}, top)
	}
}

func tiles(width, height, tileSize int) [][]*Tile {
	rowCount := int(height) / tileSize
	colCount := int(width) / tileSize

	tiles := [][]*Tile{}
	for i := 0; i < rowCount; i++ {
		row := []*Tile{}
		for j := 0; j < colCount; j++ {
			t := &Tile{
				row: i, col: j,
				Size: tileSize,
				X:    j * tileSize,
				Y:    i * tileSize,
			}
			t.Rect = image.Rect(t.X, t.Y, t.X+t.Size, t.Y+t.Size)

			row = append(row, t)
		}
		tiles = append(tiles, row)
	}

	return tiles
}

func placeBombs(difficulty int, tiles [][]*Tile, r *rand.Rand) {
	if len(tiles) == 0 {
		return
	}

	rowCount := len(tiles)
	colCount := len(tiles[0])
	tileCount := rowCount * colCount
	bombCount := int(math.Floor(float64(tileCount) * diffFill[difficulty]))

	for i := 0; i < bombCount; i++ {
		row := r.Intn(rowCount)
		col := r.Intn(colCount)

		if tiles[row][col].bomb {
			i--
			continue
		}
		tiles[row][col].bomb = true
	}

	// tiles[0][0].bomb = true
}
