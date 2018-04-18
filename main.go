package main

import (
	"github.com/courier_journey/courier"
)

func main() {
	// this is to optimize the waypoints between source and destination
	courier.OptimizeRoute()
}
