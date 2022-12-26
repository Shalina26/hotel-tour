package numbersgif

import (
	"hotel-tour/gif/drawline"
	"hotel-tour/gif/rectangle"
	"image"
)

// draw number zero
func Zero(img *image.Paletted, x0 int, y0 int, x1 int, y1 int) {
	rectangle.DrawRect(img, x0, y0, x1, y1, 1) //whole circle
}

// draw number one
func One(img *image.Paletted, y0 int, x1 int, y1 int) {
	drawline.DrawLine(img, x1, y0, x1, y1, 1) //right side
}

// draw number two
func Two(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, halfY int) {
	drawline.DrawLine(img, x0, y0, x1, y0, 1)       //top
	drawline.DrawLine(img, x0, halfY, x1, halfY, 1) //middle
	drawline.DrawLine(img, x0, y1, x1, y1, 1)       //bottom
	drawline.DrawLine(img, x1, y0, x1, halfY, 1)    //top right
	drawline.DrawLine(img, x0, halfY, x0, y1, 1)    //bottom left
}

// draw number three
func Three(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, halfY int) {
	drawline.DrawLine(img, x1, y0, x1, y1, 1)       //right side
	drawline.DrawLine(img, x0, y0, x1, y0, 1)       //top
	drawline.DrawLine(img, x0, halfY, x1, halfY, 1) //middle
	drawline.DrawLine(img, x0, y1, x1, y1, 1)       //bottom
}

// draw number four
func Four(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, halfY int) {
	drawline.DrawLine(img, x1, y0, x1, y1, 1)       //right side
	drawline.DrawLine(img, x0, y0, x0, halfY, 1)    //top left
	drawline.DrawLine(img, x0, halfY, x1, halfY, 1) //middle
}

// draw number five
func Five(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, halfY int) {
	drawline.DrawLine(img, x0, y0, x1, y0, 1)       //top
	drawline.DrawLine(img, x0, halfY, x1, halfY, 1) //middle
	drawline.DrawLine(img, x0, y1, x1, y1, 1)       //bottom
	drawline.DrawLine(img, x0, y0, x0, halfY, 1)    //top left
	drawline.DrawLine(img, x1, halfY, x1, y1, 1)    //bottom right
}

// draw number six
func Six(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, halfY int) {
	drawline.DrawLine(img, x0, y0, x1, y0, 1)     //top
	drawline.DrawLine(img, x0, y0, x0, halfY, 1)  //top left
	rectangle.DrawRect(img, x0, halfY, x1, y1, 1) //bottom circle
}

// draw number seven
func Seven(img *image.Paletted, x0 int, y0 int, x1 int, y1 int) {
	drawline.DrawLine(img, x0, y0, x1, y0, 1) //top
	drawline.DrawLine(img, x1, y0, x1, y1, 1) //right side
}

// draw number eight
func Eight(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, halfY int) {
	rectangle.DrawRect(img, x0, y0, x1, y1, 1)      //whole circle
	drawline.DrawLine(img, x0, halfY, x1, halfY, 1) //middle
}

// draw number nine
func Nine(img *image.Paletted, x0 int, y0 int, x1 int, y1 int, halfY int) {
	rectangle.DrawRect(img, x0, y0, x1, halfY, 1) //top circle
	drawline.DrawLine(img, x1, halfY, x1, y1, 1)  //bottom right
	drawline.DrawLine(img, x0, y1, x1, y1, 1)     //bottom
}

// DrawNumber takes in coordinates and draws a number according to the input that was given.
func DrawNumber(img *image.Paletted, num int, x0 int, y0 int, x1 int, y1 int, halfY int) {
	switch num {
	case 0:
		Zero(img, x0, y0, x1, y1)
	case 1:
		One(img, y0, x1, y1)
	case 2:
		Two(img, x0, y0, x1, y1, halfY)
	case 3:
		Three(img, x0, y0, x1, y1, halfY)
	case 4:
		Four(img, x0, y0, x1, y1, halfY)
	case 5:
		Five(img, x0, y0, x1, y1, halfY)
	case 6:
		Six(img, x0, y0, x1, y1, halfY)
	case 7:
		Seven(img, x0, y0, x1, y1)
	case 8:
		Eight(img, x0, y0, x1, y1, halfY)
	case 9:
		Nine(img, x0, y0, x1, y1, halfY)
	}
}
