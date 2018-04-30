package courier

import (
	"log"
	"strconv"
	"strings"

	"github.com/courier_journey/haversine"
	"github.com/courier_journey/utils"
)

const (
	unit = 1.0
)

// OptimizeRoute get the all coordinates(Lattitude and langitude)
// calculating the flight path distance (0, len(points)-1) using the haversine formula gives great-circle distances between two points
// on a sphere from their longitudes and latitudes
// next ignoring the erraneous index
func OptimizeRoute() {

	points := utils.ReadCSV()

	finalPoints := removeDuplicate(points)
	p, q := getCoordinates(finalPoints, 0, len(finalPoints)-1)
	_, flightPathDistance := haversine.Distance(p, q)
	log.Println(flightPathDistance, "Kilometer")
	var errnoeusindex []int
	errnoeusindex = getRoute(finalPoints, flightPathDistance)
	if len(errnoeusindex) == 0 {
		err := utils.WriteCSV(finalPoints)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("No Erroneus coordinates found")
		return
	}
	optimizedData := make([]string, 0)
	var j, k = 0, 0
	for i := 0; i < len(finalPoints); i++ {
		if i == errnoeusindex[j] && j != len(errnoeusindex)-1 {
			j++
		} else {
			optimizedData = append(optimizedData, finalPoints[i])
			k++
		}
	}
	err := utils.WriteCSV(finalPoints)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Optiizedpoints.csv contains the optimized points")
	return

}

// it accepts the points and flightPathDistance and returns the errnoeus index
func getRoute(points []string, flightPathDistance float64) (errnoeusindex []int) {
	i := 0
	errnoeusindex = make([]int, 0)
	// Recursive logic  needs to be implemented
	for {
		if i >= len(points)-2 {
			break
		}
		p, q := getCoordinates(points, i, i+1)
		_, distance1 := haversine.Distance(p, q)
		p, q = getCoordinates(points, i+1, i+2)
		_, distance2 := haversine.Distance(p, q)
		if distance1+distance2 >= (0.5 * flightPathDistance) {
			errnoeusindex = append(errnoeusindex, i+1)
		}
		i++

	}
	return
}

//removeDuplicate will remove the duplicated values from the list of points
func removeDuplicate(points []string) []string {
	var finalPoints []string
	flag := false
	for i := 0; i < len(points); i++ {
		if i != 0 {
			for j := 0; j < len(finalPoints); j++ {

				if finalPoints[j] == points[i] {
					flag = false
					break
				} else {
					flag = true
				}
			}

		} else {
			finalPoints = append(finalPoints, points[i])
		}
		if flag {
			finalPoints = append(finalPoints, points[i])
		}
	}
	return finalPoints
}

func getCoordinates(points []string, i, j int) (p, q haversine.Coord) {
	points1 := strings.Split(points[i], ",")
	points2 := strings.Split(points[j], ",")
	p = haversine.Coord{
		Lat: parse(points1[0]),
		Lon: parse(points1[1]),
	}
	q = haversine.Coord{
		Lat: parse(points2[0]),
		Lon: parse(points2[1]),
	}
	return
}

func parse(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("failed to parse", err)
	}
	return f
}
