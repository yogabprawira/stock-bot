package main

import "time"

const GlobalScheme = "http"
const GlobalHost = "api.marketstack.com"

type StockData struct {
	Open        float32 `json:"open"`
	High        float32 `json:"high"`
	Low         float32 `json:"low"`
	Close       float32 `json:"close"`
	Volume      float32 `json:"volume"`
	AdjHigh     float32 `json:"adj_high"`
	AdjLow      float32 `json:"adj_low"`
	AdjClose    float32 `json:"adj_close"`
	AdjOpen     float32 `json:"adj_open"`
	AdjVolume   float32 `json:"adj_volume"`
	SplitFactor float32 `json:"split_factor"`
	Dividend    float32 `json:"dividend"`
	Date        string  `json:"date"`
	Symbol      string  `json:"symbol"`
	Exchange    string  `json:"exchange"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
	Total  int `json:"total"`
}

type Response struct {
	Pagination Pagination
	Data       []StockData `json:"data"`
}

type WeekdayValuePair struct {
	Weekday    time.Weekday
	Value      float32
	Percentage float32
}
