package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type WeekdayValuePairList []WeekdayValuePair

func (p WeekdayValuePairList) Len() int           { return len(p) }
func (p WeekdayValuePairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p WeekdayValuePairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func DisplayMap(title string, weekdayMap map[time.Weekday]float32, currentPrice float32) {
	fmt.Println("==" + title + "==")
	var l WeekdayValuePairList
	for weekday := range weekdayMap {
		val := weekdayMap[weekday]
		if currentPrice == 0 {
			l = append(l, WeekdayValuePair{
				Weekday: weekday,
				Value:   val,
			})
		} else {
			l = append(l, WeekdayValuePair{
				Weekday:    weekday,
				Value:      val,
				Percentage: (currentPrice - val) * 100 / val,
			})
		}

	}
	sort.Sort(l)
	for i := range l {
		if currentPrice == 0 {
			fmt.Println(l[i].Weekday.String() + ": \t" + fmt.Sprint(math.Floor(float64(l[i].Value))))
		} else {
			fmt.Println(l[i].Weekday.String() + ": \t" + fmt.Sprint(math.Floor(float64(l[i].Value))) + " \t" + fmt.Sprint(math.Floor(float64(l[i].Percentage))) + "%")
		}
	}
}
