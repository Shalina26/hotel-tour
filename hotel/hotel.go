package hotel

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type HotelType struct {
	HotelNr     int
	Name        string
	NextHotelNr int
	Distance    int
}

// GetHotels opens a .csv file and reads it.
// Each line represents a set of data for one hotel and will be stored in a struct.
// Each hotel-struct will be added to a slice of hotels.
func GetHotels(lineCount int) ([]HotelType, error) {
	var item HotelType
	hotels := []HotelType{}

	file, err := os.Open("Hotels.csv")
	if err != nil {
		return hotels, errors.New("CSV-Datei konnte nicht geöffnet werden.")
	}

	csvreader := csv.NewReader(file)

	for i := 0; i <= lineCount; i++ {
		dataset, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return hotels, errors.New(err.Error())
		}

		dataset0 := dataset[0]
		dataset1, _ := strconv.Atoi(dataset[1])

		if i+2 > lineCount {
			item = HotelType{i + 1, dataset0, 1, dataset1}
		} else {
			item = HotelType{i + 1, dataset0, i + 2, dataset1}
		}
		hotels = append(hotels, item)
	}
	file.Close()
	return hotels, nil
}

// PrintHotels takes in the hotel slice and prints each element in a readable format.
func PrintHotels(hotels []HotelType) {
	fmt.Println("Folgende Hotels gibt es:")
	fmt.Println("----------------------------------------")
	for i, item := range hotels {
		fmt.Printf("Hotel %d: %s\n", i+1, item.Name)
	}
	fmt.Println("----------------------------------------")
}

// FindHotelNumberByName takes in the name of a hotel and tries to find the corresponding number to it.
func FindHotelNumberByName(hotels []HotelType, hotelname string, number int) (int, error) {
	for i, element := range hotels {
		if element.Name == hotelname {
			return i, nil
		}
	}
	return -1, errors.New("Der eingegebene Name ist nicht gültig.")
}

// ReverseHotelSlice takes in the tour, reverses it and returns the reversed tour.
func ReverseHotelSlice(s []HotelType) []HotelType {
	reversed := []HotelType{}
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return reversed
}
