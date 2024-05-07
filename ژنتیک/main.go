package main

import "fmt"

// calculate calculates the minimum cost of creating a desired sub-dna in a dna string
// by inserting characters to it
func calculate(dnaStr string, subStr string) int {
	dnaLenght := len(dnaStr)
	subStrLength := len(subStr)

	cost := 0

	for i := 0; i < subStrLength; i++ {
		cost += insertCost[rune(subStr[i])]
	}

	minmumCost := cost

	for i := 0; i < dnaLenght; i++ {
		// initialize cost with cost of substr
		tempCost := cost

		dnaPointer := i
		subStrPointer := 0
		for dnaPointer < dnaLenght && subStrPointer < subStrLength {
			if dnaStr[dnaPointer] == subStr[subStrPointer] {
				tempCost -= insertCost[rune(subStr[subStrPointer])]
				dnaPointer++
				subStrPointer++
			} else if dnaStr[dnaPointer] != subStr[subStrPointer] {
				subStrPointer++
			}
		}

		if tempCost < minmumCost {
			minmumCost = tempCost
		}
	}

	return minmumCost
}

// insertCost contains cost of inserting characters into the dna
// example insertCost['A'] = 34 -> cost of inserting A into the dna string would be 34
var insertCost map[rune]int

func main() {
	// receive dna string and the sub-dna string
	var dnaStr, subDnaStr string
	fmt.Scanf("%s\n", &dnaStr)
	fmt.Scanf("%s\n", &subDnaStr)

	// initialize insertCost map
	insertCost = make(map[rune]int, 4)

	// receive insert cost of each character
	var aCost, cCost, gCost, tCost int
	fmt.Scanf("%d%d%d%d\n", &aCost, &cCost, &gCost, &tCost)
	insertCost['A'] = aCost
	insertCost['C'] = cCost
	insertCost['G'] = gCost
	insertCost['T'] = tCost

	// calculate the cost and print it
	fmt.Println(calculate(dnaStr, subDnaStr))
}
