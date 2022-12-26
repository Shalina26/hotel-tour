package rectangle

import (
	"hotel-tour/gif/drawline"
	"image"
)

// DrawRect draws a rectangle using given coordinates.
func DrawRect(img *image.Paletted, x0, y0, x1, y1 int, col uint8) {
	drawline.DrawLine(img, x0, y0, x1, y0, col)
	drawline.DrawLine(img, x0, y1, x1, y1, col)
	drawline.DrawLine(img, x1, y0, x1, y1, col)
	drawline.DrawLine(img, x0, y0, x0, y1, col)
}
