package main

import (
	"fmt"
	"math"
)

func DisplayWeekdayRank(title string, l WeekdayValuePairList, currentPrice float32) {
	fmt.Println("==" + title + "==")
	for i := range l {
		if currentPrice == 0 {
			fmt.Println(l[i].Weekday.String() + ": \t" + fmt.Sprint(math.Floor(float64(l[i].Value))))
		} else {
			fmt.Println(l[i].Weekday.String() + ": \t" + fmt.Sprint(math.Floor(float64(l[i].Value))) + " \t" +
				fmt.Sprint(math.Floor(float64(l[i].Percentage))) + "%")
		}
	}
}

func DisplayPartOfMonthRank(title string, l PartOfMonthValuePairList, currentPrice float32) {
	fmt.Println("==" + title + "==")
	for i := range l {
		if currentPrice == 0 {
			fmt.Println(PartOfMonthToString(l[i].PartOfMonth) + ": \t" + fmt.Sprint(math.Floor(float64(l[i].Value))))
		} else {
			fmt.Println(PartOfMonthToString(l[i].PartOfMonth) + ": \t" + fmt.Sprint(math.Floor(float64(l[i].Value))) + " \t" +
				fmt.Sprint(math.Floor(float64(l[i].Percentage))) + "%")
		}
	}
}
