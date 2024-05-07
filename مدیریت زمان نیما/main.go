package main

import "fmt"

func main() {
	// scan values
	w, s, i := scanValues()

	// calculate r
	r := calculateFreeTime(w, s, i)

	// print r
	fmt.Println(r)
}

// calculateFreeTime calculates nima's free time
func calculateFreeTime(w, s, i int) int {
	return 24 - (w + s - i)
}

// scanValues scans values for w, s, i and returns them
// w: hours only working
// s: hours only studying
// i: hours doing both studying and working
func scanValues() (w int, s int, i int) {
	fmt.Scanf("%d %d %d", &s, &w, &i)
	return w, s, i
}
