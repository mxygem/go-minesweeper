package minesweeper

import (
	"fmt"
	"image/color"
)

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 255}
	cRed            = color.RGBA{255, 0, 0, 255}
	cOrange         = color.RGBA{245, 120, 0, 255}
	cYellow         = color.RGBA{255, 255, 0, 255}
	cGreen          = color.RGBA{0, 255, 0, 255}
	cBlue           = color.RGBA{0, 0, 255, 255}
	cCyan           = color.RGBA{0, 255, 255, 255}
	cViolet         = color.RGBA{125, 0, 255, 255}
	cPink           = color.RGBA{255, 0, 255, 255}
	cBrown          = color.RGBA{125, 50, 0, 255}
	cBlack          = color.RGBA{0, 0, 0, 255}
	cGrey20         = color.RGBA{51, 51, 51, 255}
	cGrey50         = color.RGBA{127, 127, 127, 255}
	cGrey80         = color.RGBA{204, 204, 204, 255}
	cWhite          = color.RGBA{255, 255, 255, 255}
)

func numColor(num int) color.RGBA {
	switch num {
	case 1:
		return cBlue
	case 2:
		return cGreen
	case 3:
		return cYellow
	case 4:
		return cOrange
	case 5:
		return cRed
	case 6:
		return cPink
	case 7:
		return cCyan
	case 8:
		return cBrown
	}
	panic(fmt.Sprintf("couldn't match number %d for color", num))
}
