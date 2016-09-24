package main

import (
	"flag"
	"fmt"
	"os"
)

var SmolovJr = Program{
	Days: []Day{
		Day{
			Name:       "Monday",
			Sets:       6,
			Reps:       6,
			Percentage: 0.7,
		},
		Day{
			Name:       "Wednesday",
			Sets:       7,
			Reps:       5,
			Percentage: 0.75,
		},
		Day{
			Name:       "Friday",
			Sets:       8,
			Reps:       4,
			Percentage: 0.8,
		},
		Day{
			Name:       "Saturday",
			Sets:       10,
			Reps:       3,
			Percentage: 0.85,
		},
	},
}

var (
	port int
)

func parseFlags() {
	flag.IntVar(&port, "port", 0, "port to run webserver on")
	flag.Parse()
}

func main() {
	parseFlags()

	if port == 0 {
		if err := cli(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if err := serve(port); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
