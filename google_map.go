package main

import (
	"context"
	"log"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyAs5TX4krGb0hwwM7Ib1kbiND7xgqHF2FE"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      "松山區",
		Destination: "台北市松山區八德路3段155巷2之2號7樓25坪",
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	pretty.Println(route)
	//fmt.Println(route[0].Legs[0].EndLocation.Lat)

}
