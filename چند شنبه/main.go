// question: https://quera.org/problemset/183351
// difficulty: hard

package main

import "fmt"

type Month struct {
	name  string
	index int
}

type Day struct {
	name  string
	index int
}

// Shift shifts Day k times
// for example: day = jome, k = 1 -> day = shanbe
// for example: day = jome, k = -1 -> day = 5shanbe
func (d *Day) Shift(k int) {
	index := (d.index + k) % 7
	if index < 0 {
		index = 7 + index
	}
	fmt.Print()

	for i, day := range days {
		if i == index {
			d.index = day.index
			d.name = day.name
		}
	}
}

type Date struct {
	month Month
	day   int
}

// ConvertToNumber converts a Date to a number in [1, 365]
// for example date = shahrivar 3rd -> 34
// for example date = esfand 29th -> 365
func (d Date) ConvertToNumber() int {
	days := d.day
	if d.month.index < 6 {
		days += d.month.index * 31
	} else if d.month.index >= 6 && d.month.index < 11 {
		days += 6 * 31
		days += (d.month.index - 6) * 30
	} else if d.month.index == 11 {
		days += 6 * 31
		days += 5 * 30
	}

	return days
}

// months stores all months of the year
var months []Month

// days stores all days of the week
var days []Day

func main() {
	// adding all months of the year
	months = make([]Month, 12)
	months[0] = Month{name: "Farvardin", index: 0}
	months[1] = Month{name: "Ordibehesht", index: 1}
	months[2] = Month{name: "Khordad", index: 2}
	months[3] = Month{name: "Tir", index: 3}
	months[4] = Month{name: "Mordad", index: 4}
	months[5] = Month{name: "Shahrivar", index: 5}
	months[6] = Month{name: "Mehr", index: 6}
	months[7] = Month{name: "Aban", index: 7}
	months[8] = Month{name: "Azar", index: 8}
	months[9] = Month{name: "Dey", index: 9}
	months[10] = Month{name: "Bahman", index: 10}
	months[11] = Month{name: "Esfand", index: 11}

	// adding all days of the week
	days = make([]Day, 7)
	days[0] = Day{name: "shanbe", index: 0}
	days[1] = Day{name: "1shanbe", index: 1}
	days[2] = Day{name: "2shanbe", index: 2}
	days[3] = Day{name: "3shanbe", index: 3}
	days[4] = Day{name: "4shanbe", index: 4}
	days[5] = Day{name: "5shanbe", index: 5}
	days[6] = Day{name: "jome", index: 6}

	// number of scenarios
	var t int
	fmt.Scanf("%d\n", &t)

	// slice to store result of each scenario
	res := make([]string, t)

	// iterate over each scenario
	for i := 0; i < t; i++ {
		var sourceDayOfMonth int
		var sourceMonthName, sourceDayName string
		fmt.Scanf("%d%s%s\n", &sourceDayOfMonth, &sourceMonthName, &sourceDayName)

		var targetDayOfMonth int
		var targetMonthName string
		fmt.Scanf("%d%s\n", &targetDayOfMonth, &targetMonthName)

		var targetMonth Month
		for _, month := range months {
			if month.name == targetMonthName {
				targetMonth = month
				break
			}
		}
		targetDate := Date{
			month: targetMonth,
			day:   targetDayOfMonth,
		}

		var sourceMonth Month
		for _, month := range months {
			if month.name == sourceMonthName {
				sourceMonth = month
				break
			}
		}
		sourceDate := Date{
			month: sourceMonth,
			day:   sourceDayOfMonth,
		}

		var day Day
		for _, item := range days {
			if item.name == sourceDayName {
				day = item
			}
		}

		shiftAmount := (targetDate.ConvertToNumber() - sourceDate.ConvertToNumber()) % 7
		day.Shift(shiftAmount)
		res[i] = day.name
	}

	// print results
	for i := 0; i < t; i++ {
		fmt.Println(res[i])
	}
}
