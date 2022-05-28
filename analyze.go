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

func FindAvgWeekday(stockDatas []StockData) ResultWeekdayRank {
	stockDatas = SortStockDatas(stockDatas)
	currentStockPrice := stockDatas[len(stockDatas)-1].Close

	weekdayMap := CreateWeekdayMap(stockDatas)

	var result ResultWeekdayRank

	openMap := FindAvgWeekdayFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Open
	})
	openRank := CreateWeekdayRank(openMap, currentStockPrice)
	result.Open = openRank
	DisplayWeekdayRank("Open Price", openRank, currentStockPrice)

	lowMap := FindAvgWeekdayFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Low
	})
	lowRank := CreateWeekdayRank(lowMap, currentStockPrice)
	result.Low = lowRank
	DisplayWeekdayRank("Low Price", lowRank, currentStockPrice)

	highMap := FindAvgWeekdayFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.High
	})
	highRank := CreateWeekdayRank(highMap, currentStockPrice)
	result.High = highRank
	DisplayWeekdayRank("High Price", highRank, currentStockPrice)

	closeMap := FindAvgWeekdayFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Close
	})
	closeRank := CreateWeekdayRank(closeMap, currentStockPrice)
	result.Close = closeRank
	DisplayWeekdayRank("Close Price", closeRank, currentStockPrice)

	VolMap := FindAvgWeekdayFromMap(weekdayMap, func(stockData StockData) float32 {
		return stockData.Volume
	})
	volRank := CreateWeekdayRank(VolMap, currentStockPrice)
	result.Vol = volRank
	DisplayWeekdayRank("Volume", volRank, 0)

	return result
}

func FindAvgWeekdayFromMap(weekdayMap map[time.Weekday][]StockData, getVal func(StockData) float32) map[time.Weekday]float32 {
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

func CreateWeekdayRank(weekdayMap map[time.Weekday]float32, currentPrice float32) WeekdayValuePairList {
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

func FindWeekdayTotalRank(resultList []ResultWeekdayRank) WeekdayValuePairList {
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
	return CreateWeekdayRank(rankMap, 0)
}

func CreatePartOfMonthMap(stockDatas []StockData) map[int][]StockData {
	partOfMonthMap := make(map[int][]StockData, 0)
	for _, stockData := range stockDatas {
		partOfMonth := ConvertDayToPartOfMonth(GetTimeStamp(stockData.Date).Day())
		partOfMonthMap[partOfMonth] = append(partOfMonthMap[partOfMonth], stockData)
	}
	return partOfMonthMap
}

func FindAvgPartOfMonth(stockDatas []StockData) ResultPartOfMonthRank {
	stockDatas = SortStockDatas(stockDatas)
	currentStockPrice := stockDatas[len(stockDatas)-1].Close

	partOfMonthMap := CreatePartOfMonthMap(stockDatas)

	var result ResultPartOfMonthRank

	openMap := FindAvgPartOfMonthFromMap(partOfMonthMap, func(stockData StockData) float32 {
		return stockData.Open
	})
	openRank := CreatePartOfMonthRank(openMap, currentStockPrice)
	result.Open = openRank
	DisplayPartOfMonthRank("Open Price", openRank, currentStockPrice)

	lowMap := FindAvgPartOfMonthFromMap(partOfMonthMap, func(stockData StockData) float32 {
		return stockData.Low
	})
	lowRank := CreatePartOfMonthRank(lowMap, currentStockPrice)
	result.Low = lowRank
	DisplayPartOfMonthRank("Low Price", lowRank, currentStockPrice)

	highMap := FindAvgPartOfMonthFromMap(partOfMonthMap, func(stockData StockData) float32 {
		return stockData.High
	})
	highRank := CreatePartOfMonthRank(highMap, currentStockPrice)
	result.High = highRank
	DisplayPartOfMonthRank("High Price", highRank, currentStockPrice)

	closeMap := FindAvgPartOfMonthFromMap(partOfMonthMap, func(stockData StockData) float32 {
		return stockData.Close
	})
	closeRank := CreatePartOfMonthRank(closeMap, currentStockPrice)
	result.Close = closeRank
	DisplayPartOfMonthRank("Close Price", closeRank, currentStockPrice)

	volMap := FindAvgPartOfMonthFromMap(partOfMonthMap, func(stockData StockData) float32 {
		return stockData.Volume
	})
	volRank := CreatePartOfMonthRank(volMap, currentStockPrice)
	result.Vol = volRank
	DisplayPartOfMonthRank("Volume", volRank, 0)

	return result
}

func FindAvgPartOfMonthFromMap(partOfMonthMap map[int][]StockData, getVal func(StockData) float32) map[int]float32 {
	partOfMonthMapAvg := make(map[int]float32, 0)
	for weekDay := range partOfMonthMap {
		var sum float32 = 0
		count := 0
		for _, stockData := range partOfMonthMap[weekDay] {
			sum += getVal(stockData)
			count++
		}
		partOfMonthMapAvg[weekDay] = sum / float32(count)
	}
	return partOfMonthMapAvg
}

func CreatePartOfMonthRank(partOfMonthMap map[int]float32, currentPrice float32) PartOfMonthValuePairList {
	var l PartOfMonthValuePairList
	for partOfMonth := range partOfMonthMap {
		val := partOfMonthMap[partOfMonth]
		var p PartOfMonthValuePair
		p.PartOfMonth = partOfMonth
		p.Value = val
		if currentPrice != 0 {
			p.Percentage = (currentPrice - val) * 100 / val
		}
		l = append(l, p)
	}
	sort.Sort(l)
	return l
}

func FindPartOfMonthTotalRank(resultList []ResultPartOfMonthRank) PartOfMonthValuePairList {
	rankMap := make(map[int]float32)
	for i := range resultList {
		for j, p := range resultList[i].Open {
			rankMap[p.PartOfMonth] += float32(j + 1)
		}
		for j, p := range resultList[i].Close {
			rankMap[p.PartOfMonth] += float32(j + 1)
		}
		for j, p := range resultList[i].Low {
			rankMap[p.PartOfMonth] += float32(j + 1)
		}
		for j, p := range resultList[i].High {
			rankMap[p.PartOfMonth] += float32(j + 1)
		}
	}
	return CreatePartOfMonthRank(rankMap, 0)
}
