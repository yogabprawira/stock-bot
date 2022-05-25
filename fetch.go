package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func FetchEodData(symbols string) (Response, error) {
	return FetchData(symbols, "/v1/eod", 500, 0)
}

func FetchData(symbols string, path string, limit int, offset int) (Response, error) {
	accessKey := os.Getenv("ACCESS_KEY")
	if len(accessKey) <= 0 {
		return Response{}, fmt.Errorf("please insert access key")
	}

	httpClient := http.Client{}

	u := url.URL{
		Scheme: GlobalScheme,
		Host:   GlobalHost,
		Path:   path,
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Response{}, err
	}

	q := req.URL.Query()
	q.Add("access_key", accessKey)
	q.Add("symbols", symbols)
	q.Add("limit", fmt.Sprint(limit))
	q.Add("offset", fmt.Sprint(offset))
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	var apiResponse Response
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return Response{}, err
	}
	return apiResponse, nil
}
