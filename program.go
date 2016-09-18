package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/CrowdSurge/banner"
)

type Program struct {
	Days []Day
}

type Day struct {
	Name       string
	Sets       int
	Reps       int
	Percentage float64
}

func (p Program) Print(liftName string, max int, weeks int, increment int) {
	banner.Print(liftName)
	fmt.Println(weeks, "week Smolov Jr. program")
	fmt.Println("==========================================")
	for i := 0; i < weeks; i++ {
		fmt.Println("\nWeek", i+1)
		fmt.Println("---------")

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		fmt.Fprintln(w, "\tSets\tReps\t Weight\tWorksheet\t")

		for _, day := range p.Days {
			fmt.Fprintf(w, "%s\t", day.Name)
			fmt.Fprintf(w, "%d\t", day.Sets)
			fmt.Fprintf(w, "%d\t", day.Reps)
			fmt.Fprintf(w, "%d\t  ", int(day.Percentage*float64(max))+(increment*i))
			for j := 0; j < day.Sets; j++ {
				fmt.Fprint(w, "| ")
			}
			fmt.Fprintf(w, "|\t\n")
		}

		w.Flush()
	}
}
