package main

import (
	"sort"
	"time"
)

func GetTimeStamp(dateTime string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", dateTime)
	return t
}

func ConvertDayToPartOfMonth(day int) int {
	if day <= 10 {
		return StartOfMonth
	} else if day > 20 {
		return EndOfMonth
	} else {
		return MiddleOfMonth
	}
}

func PartOfMonthToString(partOfMonth int) string {
	switch partOfMonth {
	case StartOfMonth:
		return "Start of Month"
	case MiddleOfMonth:
		return "Middle of Month"
	case EndOfMonth:
		return "End of Month"
	default:
		return ""
	}
}

func SortStockDatas(stockDatas []StockData) []StockData {
	sort.Slice(stockDatas, func(i, j int) bool {
		return GetTimeStamp(stockDatas[i].Date).Before(GetTimeStamp(stockDatas[j].Date))
	})
	return stockDatas
}

func MergeData(resp1 Response, resp2 Response) (Response, error) {
	resp1.Data = SortStockDatas(resp1.Data)
	resp2.Data = SortStockDatas(resp2.Data)
	var resp Response
	resp = resp1
	minDate := GetTimeStamp(resp.Data[0].Date)
	maxDate := GetTimeStamp(resp.Data[len(resp.Data)-1].Date)
	for _, stockData := range resp2.Data {
		if GetTimeStamp(stockData.Date).Before(minDate) || GetTimeStamp(stockData.Date).After(maxDate) {
			resp.Data = append(resp.Data, stockData)
			resp.Pagination.Count++
			resp.Pagination.Total++
		}
	}
	resp.Data = SortStockDatas(resp.Data)
	return resp, nil
}
