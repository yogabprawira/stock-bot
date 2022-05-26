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

func FindAvgWeekday(stockDatas []StockData) ResultRank {
	sort.Slice(stockDatas, func(i, j int) bool {
		return GetTimeStamp(stockDatas[i].Date).Before(GetTimeStamp(stockDatas[j].Date))
	})
	currentStockPrice := stockDatas[len(stockDatas)-1].Close

	weekdayMap := CreateWeekdayMap(stockDatas)

	var result ResultRank

	openMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Open
	})
	openRank := CreateRank(openMap, currentStockPrice)
	result.Open = openRank
	DisplayMap("Open Price", openRank, currentStockPrice)

	lowMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Low
	})
	lowRank := CreateRank(lowMap, currentStockPrice)
	result.Low = lowRank
	DisplayMap("Low Price", lowRank, currentStockPrice)

	highMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.High
	})
	highRank := CreateRank(highMap, currentStockPrice)
	result.High = highRank
	DisplayMap("High Price", highRank, currentStockPrice)

	closeMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Close
	})
	closeRank := CreateRank(closeMap, currentStockPrice)
	result.Close = closeRank
	DisplayMap("Close Price", closeRank, currentStockPrice)

	VolMap := FindAvgFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Volume
	})
	volRank := CreateRank(VolMap, currentStockPrice)
	result.Vol = volRank
	DisplayMap("Volume", volRank, 0)

	return result
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

func CreateRank(weekdayMap map[time.Weekday]float32, currentPrice float32) WeekdayValuePairList {
	var l WeekdayValuePairList
	for weekday := range weekdayMap {
		val := weekdayMap[weekday]
		var w WeekdayValuePair
		w.Weekday = weekday
		w.Value = val
		if currentPrice != 0 {
			w.Percentage = (currentPrice - val) * 100 / val
		}
		l = append(l, w)
	}
	sort.Sort(l)
	return l
}

func FindTotalRank(resultList []ResultRank) WeekdayValuePairList {
	rankMap := make(map[time.Weekday]float32)
	for i := range resultList {
		for j, w := range resultList[i].Open {
			rankMap[w.Weekday] += float32(j + 1)
		}
		for j, w := range resultList[i].Close {
			rankMap[w.Weekday] += float32(j + 1)
		}
		for j, w := range resultList[i].Low {
			rankMap[w.Weekday] += float32(j + 1)
		}
		for j, w := range resultList[i].High {
			rankMap[w.Weekday] += float32(j + 1)
		}
	}
	return CreateRank(rankMap, 0)
}
