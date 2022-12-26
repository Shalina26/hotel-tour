package csvFile

import (
	"bufio"
	"os"
)

// CountLines counts the lines of a given .csv file
func CountLines() int {
	var lineCount int

	file, _ := os.Open("Hotels.csv")
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lineCount++
	}
	file.Close()
	return lineCount
}
