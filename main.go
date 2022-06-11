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

var symbolListIdx30 = []string{
	"ADRO.XIDX",
	"ANTM.XIDX",
	"ASII.XIDX",
	"BBCA.XIDX",
	"BBNI.XIDX",
	"BBRI.XIDX",
	"BBTN.XIDX",
	"BMRI.XIDX",
	"BRPT.XIDX",
	"BUKA.XIDX",
	"CPIN.XIDX",
	"EMTK.XIDX",
	"EXCL.XIDX",
	"ICBP.XIDX",
	"INCO.XIDX",
	"INDF.XIDX",
	"INKP.XIDX",
	"KLBF.XIDX",
	"MDKA.XIDX",
	"MIKA.XIDX",
	"PGAS.XIDX",
	"PTBA.XIDX",
	"SMGR.XIDX",
	"TBIG.XIDX",
	"TINS.XIDX",
	"TLKM.XIDX",
	"TOWR.XIDX",
	"UNTR.XIDX",
	"UNVR.XIDX",
	"WSKT.XIDX",
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

	totalWeekdayResult := make([]ResultWeekdayRank, 0)
	totalPartOfMonthResult := make([]ResultPartOfMonthRank, 0)

	//symbolList := symbolListIdx
	//symbolList := symbolListNasdaq
	symbolList := symbolListIdx30

	for _, symbol := range symbolList {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Symbol:" + symbol)

		resp, err = LoadFile("data", symbol)
		if err != nil {
			fmt.Println("Fetch data online")
			resp, err = FetchEodData(symbol)
			if err != nil {
				log.Fatalln(err)
			}
			err = SaveFile(symbol, resp)
			if err != nil {
				log.Fatalln(err)
			}
		}

		weekdayResult := FindAvgWeekday(resp.Data)
		totalWeekdayResult = append(totalWeekdayResult, weekdayResult)

		partOfMonthResult := FindAvgPartOfMonth(resp.Data)
		totalPartOfMonthResult = append(totalPartOfMonthResult, partOfMonthResult)
	}

	totalWeekdayRank := FindWeekdayTotalRank(totalWeekdayResult)
	fmt.Println("")
	fmt.Println("===WEEKDAY RANK===")
	for i := range totalWeekdayRank {
		fmt.Println(totalWeekdayRank[i].Weekday.String(), ":", totalWeekdayRank[i].Value)
	}

	totalPartOfMonthRank := FindPartOfMonthTotalRank(totalPartOfMonthResult)
	fmt.Println("")
	fmt.Println("===PART OF MONTH RANK===")
	for i := range totalPartOfMonthRank {
		fmt.Println(PartOfMonthToString(totalPartOfMonthRank[i].PartOfMonth), ":", totalPartOfMonthRank[i].Value)
	}
}
