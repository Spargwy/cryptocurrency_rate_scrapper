package storage

import "encoding/json"

type PriceMultyFull struct {
	Raw     json.RawMessage `json:"RAW"`
	Display json.RawMessage `json:"DISPLAY"`
}

type CurrencyRaw struct {
	Change24Hour    float64 `json:"CHANGE24HOUR"`
	ChangePCT24Hour float64 `json:"CHANGEPCT24HOUR"`
	Open24Hour      float64 `json:"OPEN24HOUR"`
	Volume24Hour    float64 `json:"VOLUME24HOUR"`
	Voulume24HourTo float64 `json:"VOLUME24HOURTO"`
	Low24Hour       float64 `json:"LOW24HOUR"`
	High24Hour      float64 `json:"HIGH24HOUR"`
	Price           float64 `json:"PRICE"`
	Supply          float64 `json:"SUPPLY"`
	Mktcap          float64 `json:"MKTCAP"`
}

type CurrencyDisplay struct {
	Change24Hour    string `json:"CHANGE24HOUR"`
	ChangePCT24Hour string `json:"CHANGEPCT24HOUR"`
	Open24Hour      string `json:"OPEN24HOUR"`
	Volume24Hour    string `json:"VOLUME24HOUR"`
	Voulume24HourTo string `json:"VOLUME24HOURTO"`
	Low24Hour       string `json:"LOW24HOUR"`
	High24Hour      string `json:"HIGH24HOUR"`
	Price           string `json:"PRICE"`
	Supply          string `json:"SUPPLY"`
	Mktcap          string `json:"MKTCAP"`
}
