package minesweeper

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

type mouseButton int

const (
	mouseButtonNone mouseButton = iota
	mouseButtonLeft
	mouseButtonMiddle
	mouseButtonRight
	// mouseButtonBoth ??
)

type Input struct {
	mouseState  mouseState
	mousePosX   int
	mousePosY   int
	mouseButton mouseButton
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() {
	switch i.mouseState {
	case mouseStateNone:
		switch {
		case ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft):
			i.mouseButton = mouseButtonLeft
			i.mouseState = mouseStatePressing
			// x, y := ebiten.CursorPosition()
			// fmt.Printf("left click pressed x: %d, y: %d\n", x, y)
			// i.mousePosX = x
			// i.mousePosY = y
		case ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle):
			i.mouseButton = mouseButtonMiddle
			i.mouseState = mouseStatePressing
			// x, y := ebiten.CursorPosition()
			// fmt.Printf("middle click pressed x: %d, y: %d\n", x, y)
			// i.mousePosX = x
			// i.mousePosY = y
		case ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight):
			i.mouseButton = mouseButtonRight
			i.mouseState = mouseStatePressing
			// x, y := ebiten.CursorPosition()
			// fmt.Printf("right pressed x: %d, y: %d\n", x, y)
			// i.mousePosX = x
			// i.mousePosY = y
		}
	case mouseStatePressing:
		switch i.mouseButton {
		case mouseButtonLeft:
			if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				x, y := ebiten.CursorPosition()
				fmt.Printf("left released x: %d, y: %d\n", x, y)
				i.mousePosX = x
				i.mousePosY = y
				i.mouseState = mouseStateSettled
			}
		case mouseButtonMiddle:
			if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
				x, y := ebiten.CursorPosition()
				fmt.Printf("middle released x: %d, y: %d\n", x, y)
				i.mousePosX = x
				i.mousePosY = y
				i.mouseState = mouseStateSettled
			}
		case mouseButtonRight:
			if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
				x, y := ebiten.CursorPosition()
				fmt.Printf("right released x: %d, y: %d\n", x, y)
				i.mousePosX = x
				i.mousePosY = y
				i.mouseState = mouseStateSettled
			}
		}
	case mouseStateSettled:
		i.mouseButton = mouseButtonNone
		i.mouseState = mouseStateNone
	}
}
