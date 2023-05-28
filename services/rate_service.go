package services

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CoinGeckoResponse struct {
	Bitcoin map[string]float64 `json:"bitcoin"`
}

func GetBitcoinPrice() (float64, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah&precision=3")
	if err != nil {
		return 0, errors.New("third party API error: " + err.Error())
	}
	defer resp.Body.Close()

	var data CoinGeckoResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, errors.New("response decoding error: " + err.Error())
	}
	bitcoinPrice := data.Bitcoin["uah"]

	return bitcoinPrice, nil
}
