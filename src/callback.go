// From: Jonathan Kans

package main

import (
	"image"
	"image/png"
	"os"
)

func DrawSquare(left, top, base int, draw func(x, y int)) {
	if draw == nil {
		return
	}

	right := left + base
	bottom := top + base

	for x := left; x <= right; x++ {
		draw(x, top)
		draw(x, bottom)
	}

	for y := top; y <= bottom; y++ {
		draw(left, y)
		draw(right, y)
	}
}

func main() {

	var thickFrame bool

	rgba := image.NewRGBA(300, 200)
	if rgba == nil {
		return
	}

	redColor := image.RGBAColor{255, 0, 0, 255}
	blueColor := image.RGBAColor{0, 0, 255, 255}

	pngFile, err := os.Open("squares.png", os.O_WRONLY|os.O_CREATE, 0644)
	if pngFile == nil || err != nil {
		return
	}
	defer pngFile.Close()

	frameSquare := func(x, y int) {
		// closure effortlessly passes local variables to callback
		if thickFrame {
			// draw 3 x 3 pixel block for thicker rectangle
			for x0 := x - 1; x0 <= x+1; x0++ {
				for y0 := y - 1; y0 <= y+1; y0++ {
					rgba.Set(x0, y0, redColor)
				}
			}
		} else {
			rgba.Set(x, y, blueColor)
		}
	}

	thickFrame = true
	DrawSquare(30, 20, 50, frameSquare)
	thickFrame = false
	DrawSquare(100, 120, 35, frameSquare)

	png.Encode(pngFile, rgba)
}
