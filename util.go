package main

import (
	"sort"
	"time"
)

func GetTimeStamp(dateTime string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", dateTime)
	return t
}

func MergeData(resp1 Response, resp2 Response) (Response, error) {
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
	sort.Slice(resp.Data, func(i, j int) bool {
		return GetTimeStamp(resp.Data[i].Date).Before(GetTimeStamp(resp.Data[j].Date))
	})
	return resp, nil
}
