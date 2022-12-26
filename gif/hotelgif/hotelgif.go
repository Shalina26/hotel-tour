package hotelgif

import (
	"hotel-tour/gif/drawline"
	"hotel-tour/gif/numbersgif"
	"hotel-tour/gif/rectangle"
	"hotel-tour/hotel"
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
	"strconv"
)

// createGifFile creates a new .gif file, opens and returns it.
func createGifFile() *os.File {
	file, err := os.Create("Hotels.gif")
	if err != nil {
		panic(err)
	}
	return file
}

// getGifWidth takes in the created tour and its length and calculates the width for the gif.
// It returns the calculated width.
func getGifWidth(tour []hotel.HotelType, tourLen int) int {
	var totalLineLength int

	for i := 0; i < tourLen-1; i++ {
		totalLineLength += 3
		st := strconv.Itoa(tour[i].Distance)
		for i := 0; i < len(st); i++ {
			totalLineLength += 5
		}
	}
	width := 43*tourLen + totalLineLength

	return width
}

// createBasicImage creates a colour-palette and a basic white gif.
// It returns the paletted image and animated gif.
func createBasicImage(file io.Writer, width int, height int, tour []hotel.HotelType, tourLen int) (*image.Paletted, gif.GIF) {

	palette := []color.Color{color.White, color.Black, color.RGBA{200, 100, 150, 255}}
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	anim := gif.GIF{Delay: []int{0}, Image: []*image.Paletted{img}}

	return img, anim
}

// drawWindows takes the image and coordinates for the windows.
// It draws 17 windows for one hotel.
func drawWindows(img *image.Paletted, windowX0 int) {
	windowY0 := 16
	start := windowX0

	for row := 1; row <= 4; row++ {
		windowX0 = start
		for column := 1; column <= 5; column++ {
			//avoid overlapping of door and windows
			if row != 4 {
				rectangle.DrawRect(img, windowX0, windowY0, windowX0+3, windowY0+4, 1)
			} else {
				if column == 1 || column == 5 {
					rectangle.DrawRect(img, windowX0, windowY0, windowX0+3, windowY0+4, 1)
				}
			}
			windowX0 += 6
		}
		windowY0 += 7
	}
}

// drawHotel takes in all the coordinates needed for the hotel, the tour and the number of the current hotel.
// It draws the entire hotel including the number of the hotel and its windows.
func drawHotel(img *image.Paletted, mainX0 int, doorX0 int, doorlineX int, tour []hotel.HotelType, i int, hotelNrX0 int, windowColX0 int) {
	//main building
	rectangle.DrawRect(img, mainX0, 13, mainX0+33, 44, 1)

	//door
	rectangle.DrawRect(img, doorX0, 37, doorX0+10, 44, 1)
	drawline.DrawLine(img, doorlineX, 37, doorlineX, 44, 1)

	//number of the hotel
	num := tour[i].HotelNr
	numbersgif.DrawNumber(img, num, hotelNrX0, 5, hotelNrX0+3, 11, 8)

	//windows
	drawWindows(img, windowColX0)
}

// drawDistance takes in coordinates for the line and information about the direction of the given tour.
// It draws a line and the distance between two hotels and returns the length of it.
func drawDistance(img *image.Paletted, lineX0 int, tour []hotel.HotelType, i int, x0 int, reversed bool) int {
	var st string
	if reversed {
		st = strconv.Itoa(tour[i+1].Distance)
	} else {
		st = strconv.Itoa(tour[i].Distance)
	}
	for _, number := range st {
		num, _ := strconv.Atoi(string(number))
		numbersgif.DrawNumber(img, num, x0, 24, x0+3, 30, 27)
		x0 += 5
	}

	lineX1 := lineX0 + 5*len(st) + 2

	drawline.DrawLine(img, lineX0, 30, lineX0, 34, 1) //left
	drawline.DrawLine(img, lineX0, 32, lineX1, 32, 1) //line
	drawline.DrawLine(img, lineX1, 30, lineX1, 34, 1) //right
	length := lineX1 - lineX0
	return length
}

// drawHotelRoute takes in the tour, its length and information about its direction.
// It draws all the hotels and their distances from the given tour in the right order.
func drawHotelRoute(img *image.Paletted, tour []hotel.HotelType, tourLen int, reversed bool) {
	var mainX0 int = 5
	var doorX0 int = 17
	var doorlineX int = 22
	var lineX0 int = 43
	var distanceX0 int = 45
	var hotelNrX0 int = 20
	var windowX0 int = 8

	drawHotel(img, mainX0, doorX0, doorlineX, tour, 0, hotelNrX0, windowX0)
	for i := 0; i < tourLen-1; i++ {
		lengthLine := drawDistance(img, lineX0, tour, i, distanceX0, reversed)
		mainX0 += 43 + lengthLine
		doorX0 += 43 + lengthLine
		doorlineX += 43 + lengthLine
		lineX0 += 43 + lengthLine
		distanceX0 += 43 + lengthLine
		hotelNrX0 += 43 + lengthLine
		windowX0 += 43 + lengthLine
		drawHotel(img, mainX0, doorX0, doorlineX, tour, i+1, hotelNrX0, windowX0)
	}
}

// CreateGif creates the final gif of hotels.
func CreateGif(tour []hotel.HotelType, reversed bool) {
	file := createGifFile()

	tourLen := len(tour)
	width := getGifWidth(tour, tourLen)

	img, anim := createBasicImage(file, width, 50, tour, tourLen)

	drawHotelRoute(img, tour, tourLen, reversed)

	gif.EncodeAll(file, &anim)

	file.Close()
}
