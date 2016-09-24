package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func serve(port int) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, index)
	})

	http.HandleFunc("/program", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		liftName := r.Form["lift"][0]
		maxStr := r.Form["1rm"][0]
		max, err := strconv.Atoi(maxStr)
		if err != nil {
			fmt.Fprintf(w, "unable to parse number out of: %v", maxStr)
			return
		}

		weeksStr := r.Form["weeks"][0]
		weeks, err := strconv.Atoi(weeksStr)
		if err != nil {
			fmt.Fprintf(w, "unable to parse number out of: %v", weeksStr)
			return
		}

		incrementStr := r.Form["increment"][0]
		increment, err := strconv.Atoi(incrementStr)
		if err != nil {
			fmt.Fprintf(w, "unable to parse number out of: %v", incrementStr)
			return
		}

		fmt.Fprintf(w, "<html><body><pre>")
		SmolovJr.Print(w, liftName, max, weeks, increment)
		fmt.Fprintf(w, "</pre></body></html>")
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

var index = `<html>
<body>
    <h1>Smolov Jr.</h1>
    <form action="program">
        Lift:<br>
        <input type="text" name="lift" value="Squat"><br>
        Current 1RM:<br>
        <input type="text" name="1rm" value="100"><br>
        Weeks to run:<br>
        <input type="text" name="weeks" value="3"><br>
        Increment:<br>
        <input type="text" name="increment" value="2"><br><br>
        <input type="submit" value="Submit">
    </form>
</body>
</html>`
