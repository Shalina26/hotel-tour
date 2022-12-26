package main

import (
	"fmt"
	"hotel-tour/csvFile"
	"hotel-tour/gif/hotelgif"
	"hotel-tour/hotel"
)

// getStartAndEndHotel takes in the hotels and asks the user for the names of a start-hotel and an end-hotel.
// It returns both hotels.
func getStartAndEndHotel(hotels []hotel.HotelType) (hotel.HotelType, hotel.HotelType) {
	var hotel1 string
	var hotel2 string
	var number1 int
	var number2 int

	hotel.PrintHotels(hotels)
	for {
		fmt.Printf("\nWelches Hotel soll das Starthotel sein? Hotelname: ")
		fmt.Scanf("%s\n", &hotel1)
		res, err := hotel.FindHotelNumberByName(hotels, hotel1, number1)
		number1 = res
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	current := hotels[number1]

	for {
		fmt.Printf("\nWelches Hotel ist das Ziel? Hotelname: ")
		fmt.Scanf("%s\n", &hotel2)
		res, err := hotel.FindHotelNumberByName(hotels, hotel2, number2)
		number2 = res
		if err != nil {
			fmt.Println(err)

		} else {
			if number2 != number1 {
				break
			}
			fmt.Println("Ihr Start- und Zielhotel stimmen überein.")
		}
	}
	end := hotels[number2]

	return current, end
}

// calcDistance takes in the start and end hotel.
// It adds each hotel on the route from A to B to the tour and calculates the distance between them.
// It returns the tour and the total distance from A to B.
func calcDistance(hotels []hotel.HotelType, current hotel.HotelType, end hotel.HotelType) (int, []hotel.HotelType) {
	var distance int
	tour := []hotel.HotelType{}

	for current != end {
		tour = append(tour, current)
		distance += current.Distance
		current = hotels[current.NextHotelNr-1]
	}
	tour = append(tour, end)
	return distance, tour
}

// calcShortestDistance calculates the distance from A to B and from B to A.
// hotels are managed like a circle, moving clockwise the distance from A to B and from B to A might differ.
// The shortest distance and the corresponding tour will be returned.
func calcShortestDistance(hotels []hotel.HotelType, current hotel.HotelType, end hotel.HotelType) ([]hotel.HotelType, bool) {
	distance1, tour1 := calcDistance(hotels, current, end)
	distance2, tour2 := calcDistance(hotels, end, current)

	if distance1 <= distance2 {
		fmt.Printf("\nDer kürzeste Weg von %s nach %s ist %d km lang.", current.Name, end.Name, distance1)
		return tour1, false
	} else {
		fmt.Printf("\nDer kürzeste Weg von %s nach %s ist %d km lang.", current.Name, end.Name, distance2)
		reversedTour := hotel.ReverseHotelSlice(tour2)
		return reversedTour, true
	}
}

// main creates a slice of hotel-structs by reading from a given .csv file,
// calculates the distance of the shortest route from hotel A to hotel B and creates a GIF which
// shows the hotels that are part of the tour and their distance.
func main() {
	lineCount := csvFile.CountLines()
	hotels, err := hotel.GetHotels(lineCount)
	if err != nil {
		fmt.Println(err)
		return
	}
	current, end := getStartAndEndHotel(hotels)
	tour, reversed := calcShortestDistance(hotels, current, end)

	hotelgif.CreateGif(tour, reversed)
}
