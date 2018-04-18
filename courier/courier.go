package courier

import (
	"fmt"
	"log"

	"github.com/courier_journey/utils"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// OptimizeRoute get the all coordinates(Lattitude and langitude)
// we are using external google direction api(https://github.com/googlemaps/google-maps-services-go)
// googlemaps api will accept only 23 waypoints including source and destination
// hence we are splitting the the routes according to limits
// for every request we get the optimized route for 23 points
// map all the routes with source and destination as the key
// TODO: merge all the splitted routes and make single optimized route.
func OptimizeRoute() {
	points := utils.ReadCSV()
	i := 0
	routeMap := make(map[string][]maps.Route)

	// here we are splitting the total points multiple chunks
	// each chunk will have 23 waypoints including source and destination
	for i < len(points) {
		source := points[i]
		var destination string
		var waypoints []string

		if i+22 < len(points) {
			destination = points[i+22]
			waypoints = points[i+1 : i+21]
		} else {
			destination = points[len(points)-1]
			waypoints = points[i+1:]
		}
		i = i + 22

		route := getRoute(source, destination, waypoints)
		routeMap[fmt.Sprintf("%s,%s", source, destination)] = route
		pretty.Println(routeMap)
	}

}

// it accepts the source, destination and waypoints, returs the optimised route
func getRoute(source, destination string, waypoints []string) []maps.Route {

	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyBePdA4iCPtU6h13Ea-eXGh2Nm7zu3BesQ"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.DirectionsRequest{
		Origin:      source,
		Destination: destination,
		Waypoints:   waypoints,
		Optimize:    true,
	}

	route, ponts, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	pretty.Println(ponts)
	return route
}
