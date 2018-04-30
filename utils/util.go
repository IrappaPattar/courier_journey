package utils

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Coordinates represents the lattitude , langitude and timestamp
type Coordinates struct {
	Latitude  string
	Langitude string
	TimeStamp string
}

// ReadCSV reads the coordinates from the points.csv
func ReadCSV() (waypoints []string) {
	var coOrdinates []Coordinates
	csvFile, _ := os.Open("points.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		coOrdinates = append(coOrdinates, Coordinates{
			Latitude:  line[0],
			Langitude: line[1],
			TimeStamp: line[2],
		})
	}

	for _, value := range coOrdinates {
		waypoints = append(waypoints, fmt.Sprintf("%s,%s", value.Latitude, value.Langitude))
	}
	return

}

func WriteCSV(coOrdinates []string) (err error) {
	file, err := os.Create("OptimizedPoints.csv")
	if err != nil {
		err = errors.New("Failed to create file")
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < len(coOrdinates); i++ {
		points := strings.Split(coOrdinates[i], ",")
		line := []string{fmt.Sprintf("%s", points[0]), fmt.Sprintf("%s", points[1])}
		err = writer.Write(line)
		if err != nil {
			err = errors.New("Failed to write to file")
			return
		}
	}
	return
}
