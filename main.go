package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mxygem/go-minesweeper/minesweeper"
)

func main() {
	ebiten.SetWindowSize(minesweeper.ScreenWidth, minesweeper.ScreenHeight+100)
	ebiten.SetWindowTitle("Minesweeper")

	if err := ebiten.RunGame(minesweeper.NewGame(minesweeper.Easy)); err != nil {
		log.Fatal(err)
	}
}
