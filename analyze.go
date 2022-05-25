package main

import (
	"sort"
	"time"
)

func CreateWeekdayMap(stockDatas []StockData) map[time.Weekday][]StockData {
	weekdayMap := make(map[time.Weekday][]StockData, 0)
	for _, stockData := range stockDatas {
		weekDay := GetTimeStamp(stockData.Date).Weekday()
		weekdayMap[weekDay] = append(weekdayMap[weekDay], stockData)
	}
	return weekdayMap
}

func FindAvgWeekday(stockDatas []StockData) {
	sort.Slice(stockDatas, func(i, j int) bool {
		return GetTimeStamp(stockDatas[i].Date).Before(GetTimeStamp(stockDatas[j].Date))
	})
	currentStockPrice := stockDatas[len(stockDatas)-1].Close

	weekdayMap := CreateWeekdayMap(stockDatas)

	openMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Open
	})
	DisplayMap("Open Price", openMap, currentStockPrice)

	lowMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Low
	})
	DisplayMap("Low Price", lowMap, currentStockPrice)

	highMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.High
	})
	DisplayMap("High Price", highMap, currentStockPrice)

	closeMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Close
	})
	DisplayMap("Close Price", closeMap, currentStockPrice)

	VolMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Volume
	})
	DisplayMap("Volume", VolMap, 0)
}

func FindAvgFromMap(weekdayMap map[time.Weekday][]StockData, getVal func(StockData) float32) map[time.Weekday]float32 {
	weekdayMapAvg := make(map[time.Weekday]float32, 0)
	for weekDay := range weekdayMap {
		var sum float32 = 0
		count := 0
		for _, stockData := range weekdayMap[weekDay] {
			sum += getVal(stockData)
			count++
		}
		weekdayMapAvg[weekDay] = sum / float32(count)
	}
	return weekdayMapAvg
}
