package main

import (
	"fmt"
	"log"
)

var symbolListIdx = []string{
	"ADMR.XIDX",
	"ADRO.XIDX",
	"ANTM.XIDX",
	"ARTO.XIDX",
	"ASII.XIDX",
	"BBCA.XIDX",
	"BBHI.XIDX",
	"BBKP.XIDX",
	"BBRI.XIDX",
	"BBTN.XIDX",
	"BBYB.XIDX",
	"BJTM.XIDX",
	"BUKA.XIDX",
	"CPIN.XIDX",
	"DOID.XIDX",
	"ELSA.XIDX",
	"ERAA.XIDX",
	"GOTO.XIDX",
	"ITMG.XIDX",
	"JPFA.XIDX",
	"LUCK.XIDX",
	"MAPI.XIDX",
	"MDKA.XIDX",
	"SIDO.XIDX",
	"SRTG.XIDX",
	"WIRG.XIDX",
}

var symbolListNasdaq = []string{
	"MSFT",
	"AAPL",
	"AMZN",
	"GOOGL",
	"BABA",
	"FB",
	"V",
	"JPM",
	"T",
	"INTC",
	"KO",
	"XOM",
	"ADBE",
	"WFC",
	"NVDA",
	"NFLX",
	"ORCL",
	"BA",
	"MCD",
	"NKE",
	"C",
	"TSLA",
	"HSBC",
	"PYPL",
	"PM",
	"ASML",
	"AVGO",
	"TXN",
}

func main() {
	var err error
	var resp Response

	totalResult := make([]ResultRank, 0)

	symbolList := symbolListIdx
	//symbolList = symbolListNasdaq

	for _, symbol := range symbolList {
		fmt.Println("Symbol:" + symbol)

		resp, err = FetchEodData(symbol)
		if err != nil {
			log.Fatalln(err)
		}

		err = SaveFile(symbol, resp)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err = LoadFile(symbol)
		if err != nil {
			log.Fatalln(err)
		}

		result := FindAvgWeekday(resp.Data)
		totalResult = append(totalResult, result)
	}

	totalRank := FindTotalRank(totalResult)
	fmt.Println("")
	fmt.Println("===DAY RANK===")
	for i := range totalRank {
		fmt.Println(totalRank[i].Weekday.String())
	}
}
