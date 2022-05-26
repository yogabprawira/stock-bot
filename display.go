package main

import (
	"fmt"
	"math"
)

func DisplayMap(title string, l WeekdayValuePairList, currentPrice float32) {
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
