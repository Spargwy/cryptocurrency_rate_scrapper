package storage

import "encoding/json"

type PriceMultyFull struct {
	Raw     json.RawMessage `json:"RAW"`
	Display json.RawMessage `json:"DISPLAY"`
}

type Raw struct {
	CryptoCurrency CryptoCurrency
}

type CryptoCurrency struct {
	Currency Currency
}

type Currency struct {
	Change24Hour    float64 `mapstructure:"CHANGE24HOUR"`
	ChangePCT24Hour float64 `mapstructure:"CHANGEPCT24HOUR"`
	Open24Hour      float64 `mapstructure:"OPEN24HOUR"`
	Volume24Hour    float64 `mapstructure:"VOLUME24HOUR"`
	Voulume24HourTo float64 `mapstructure:"VOLUME24HOURTO"`
	Low24Hour       float64 `mapstructure:"LOW24HOUR"`
	High24Hour      float64 `mapstructure:"HIGH24HOUR"`
	Price           float64 `mapstructure:"PRICE"`
	Supply          int     `mapstructure:"SUPPLY"`
	Mktcap          float64 `mapstructure:"MKTCAP"`
}
