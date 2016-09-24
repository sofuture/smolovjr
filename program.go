package main

import (
	"fmt"
	"io"
	"text/tabwriter"
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

func (p Program) Print(w io.Writer, liftName string, max int, weeks int, increment int) {
	fmt.Fprintln(w, liftName)
	fmt.Fprintln(w, "================================================")
	fmt.Fprintln(w, weeks, "week Smolov Jr. program -- working 1RM:", max)
	fmt.Fprintln(w, "================================================")
	for i := 0; i < weeks; i++ {
		fmt.Fprintln(w)
		tab := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.AlignRight)
		fmt.Fprintf(tab, "Week %d\t", i+1)
		fmt.Fprintln(tab, "Sets\tReps\tWeight\tWorksheet\t")
		fmt.Fprintln(tab, "------\t----\t----\t------\t---------\t")

		for _, day := range p.Days {
			fmt.Fprintf(tab, "%s\t", day.Name)
			fmt.Fprintf(tab, "%d\t", day.Sets)
			fmt.Fprintf(tab, "%d\t", day.Reps)
			fmt.Fprintf(tab, "%d\t  [ ", int(day.Percentage*float64(max))+(increment*i))
			for j := 0; j < day.Sets-1; j++ {
				fmt.Fprint(tab, "| ")
			}
			fmt.Fprintf(tab, "]\t\n")
		}

		tab.Flush()
	}
}
