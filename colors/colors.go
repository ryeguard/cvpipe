package colors

import "image/color"

var (
	White = color.RGBA{255, 255, 255, 255}
	Black = color.RGBA{0, 0, 0, 255}

	Red   = color.RGBA{255, 0, 0, 255}
	Green = color.RGBA{0, 255, 0, 255}
	Blue  = color.RGBA{0, 0, 255, 255}

	Grey   = color.RGBA{125, 125, 125, 255}
	Cyan   = color.RGBA{0, 255, 255, 255}
	Purple = color.RGBA{255, 0, 255, 255}
	Yellow = color.RGBA{255, 255, 0, 255}
)

var All []color.RGBA = []color.RGBA{
	White,
	Black,
	Grey,
	Red, Green, Blue,
	Cyan, Purple, Yellow,
}

var Full []color.RGBA = []color.RGBA{
	Red, Green, Blue,
	Cyan, Purple, Yellow,
}
