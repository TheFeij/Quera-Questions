package main

import (
	"fmt"
	"sort"
)

// Table represents a table
type Table struct {
	// d distance from the door
	d int
	// price for sticking the table to the ground
	t int
}

// subPrices is a map to store answer to sub problems
// when we stick a table to the ground its price and price of
// tables after that becomes independent of the previous tables,
// so it can be stored here and used later to prevent recalculating
// the same sub-problem
var subPrices map[int]int

// calculate calculates the minimum price that needs to be paid for this problem
// this function uses backtracking and dynamic programming approaches
//
// inputs:
// tables: stores info about tables
// priceSoFar: total price we have to pay till now
// lastStuck: position of the last table stuck to the ground (field d of Table is stored in it)
// index: current table we are working on
func calculate(tables []Table, priceSoFar int, lastStuck int, index int) int {
	tablesLength := len(tables)

	if index == tablesLength {
		return priceSoFar
	}

	var stickRes int

	subProblemRes, ok := subPrices[tables[index].d]

	if ok {
		stickRes = subProblemRes + priceSoFar
	} else {
		stickRes = calculate(tables, tables[index].t, tables[index].d, index+1)
		subPrices[tables[index].d] = stickRes
		stickRes += priceSoFar
	}

	noStickRes := calculate(tables, tables[index].d-lastStuck, lastStuck, index+1)
	noStickRes += priceSoFar

	res := stickRes
	if noStickRes < stickRes {
		res = noStickRes
	}
	return res
}

func main() {
	// receive number of tables in the restaurant
	var n int
	fmt.Scanf("%d\n", &n)

	// receive each tables information
	var tables = make([]Table, n)
	for i := 0; i < n; i++ {
		var d int
		if i < n-1 {
			fmt.Scanf("%d", &d)
		} else {
			fmt.Scanf("%d\n", &d)
		}
		tables[i] = Table{d: d}
	}
	for i := 0; i < n; i++ {
		var t int
		if i < n-1 {
			fmt.Scanf("%d", &t)
		} else {
			fmt.Scanf("%d\n", &t)
		}
		tables[i].t = t
	}

	// sort tables based on their position (position of a table: Table.d)
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].d < tables[j].d
	})

	// initialize subPrices map
	subPrices = make(map[int]int)

	// calculate the total price
	cost := calculate(tables, tables[0].t, tables[0].d, 1)

	// print the result
	fmt.Println(cost)
}
