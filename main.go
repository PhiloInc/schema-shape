package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	source      *string
	dest        *string
	srcun       *string
	srcpw       *string
	destun      *string
	destpw      *string
	numSeries   *int
	numValues   *int
	numParallel *int
)

func init() {
	source = flag.String("source", "http://localhost:8086", "hostname of inlfux server")
	dest = flag.String("dest", "http://localhost:8087", "hostname of inlfux server")
	srcun = flag.String("su", "", "influx auth username for source server")
	srcpw = flag.String("sp", "", "influx auth password for source server")
	destun = flag.String("du", "", "influx auth username for destination server")
	destpw = flag.String("dp", "", "influx auth password for destination server")
	numSeries = flag.Int("numSeries", 10000, "number of series to query at one time")
	numValues = flag.Int("numValues", 10000, "number of values to query at one time")
	numParallel = flag.Int("numParallel", 2, "number of parallel queries allowed")
	flag.Parse()
}

func main() {
	sc := NewSchamaShape(*numSeries, *numValues, *numParallel)
	sc.Hydrate()
	sc.MakeQueries()
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func iToS(face interface{}) string {
	return fmt.Sprintf("%v", face)
}
