package main

import (
	"fmt"
	"log"
)

func main() {
	var err error
	var resp Response
	symbol := "BBKP.XIDX"

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

	FindAvgWeekday(resp.Data)

}
