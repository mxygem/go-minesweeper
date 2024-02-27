package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mxygem/go-minesweeper/minesweeper"
)

const (
	diffEasy = iota
	diffMedium
	diffHard
)

func main() {
	ebiten.SetWindowSize(minesweeper.ScreenWidth, minesweeper.ScreenHeight)
	ebiten.SetWindowTitle("Minesweeper")

	if err := ebiten.RunGame(minesweeper.NewGame(diffEasy)); err != nil {
		log.Fatal(err)
	}
}
