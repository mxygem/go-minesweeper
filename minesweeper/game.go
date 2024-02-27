package minesweeper

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 400.0
	ScreenHeight = 400.0
)

type Game struct {
	board      *Board
	input      *Input
	boardImage *ebiten.Image
}

func NewGame(diff int) *Game {
	g := &Game{
		board: NewBoard(ScreenWidth, ScreenHeight, diff),
		input: NewInput(),
	}

	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		fmt.Println("attempting to restart game")
		g.board = NewBoard(g.board.width, g.board.height, g.board.difficulty)
	}

	g.input.Update()
	g.board.Update(g.input)

	return nil
}

// Draw draws the current state of the game
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)

	// op := &ebiten.DrawImageOptions{}
	// sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	// bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	// x := (sw - bw) / 2
	// y := (sh - bh) / 2
	// op.GeoM.Translate(float64(x), float64(y))
	// screen.DrawImage(g.boardImage, op)
	screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})
}
