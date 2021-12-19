package main

import "fmt"
import "time"

func main() {

	publisher := "DizzyBooks Publishing Inc."
	writer := "Tracey Hatchet"
	artist := "Jewel Tampson"
	title := "Mr. GoToSleep"

	var year int16 = 1997
	var pageNumber int16 = 14
	var grade float32 = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist,
		"year", year, "number of pages", pageNumber, "grade", grade, "publisher by", publisher)

	publisher = "DizzyBooks Publishing Inc."
	writer = "Ryan N. Shawn."
	artist = "Phoebe Paperclips"
	title = "Epic Vol. 1"
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist,
		"year", year, "number of pages", pageNumber, "grade", grade, "publisher by", publisher)

	fmt.Println(time.Now())

}
