package minesweeper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 500.0
	ScreenHeight = 500.0
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type Game struct {
	board      *Board
	input      *Input
	boardImage *ebiten.Image
}

func NewGame(diff Difficulty) *Game {
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
	// restart game when R is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.board = NewBoard(g.board.width, g.board.height, g.board.difficulty)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
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

	screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})
}
