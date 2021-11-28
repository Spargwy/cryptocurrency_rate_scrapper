package storage

import "encoding/json"

//PriceMultiFull - describes response structure from cryptocompare api
type PriceMultiFull struct {
	//Raw and Display are the only one constant structures
	//that contains currency and cryptocurrency structures that
	//differ depending on the request params
	Raw     json.RawMessage `json:"RAW"`
	Display json.RawMessage `json:"DISPLAY"`
}

//Raw - describes structure of currency and
//cryptocurrency that included into Raw and Display field of response
type Raw struct {
	Change24Hour    float64 `json:"CHANGE24HOUR"`
	ChangePCT24Hour float64 `json:"CHANGEPCT24HOUR"`
	Open24Hour      float64 `json:"OPEN24HOUR"`
	Volume24Hour    float64 `json:"VOLUME24HOUR"`
	Voulume24HourTo float64 `json:"VOLUME24HOURTO"`
	Low24Hour       float64 `json:"LOW24HOUR"`
	High24Hour      float64 `json:"HIGH24HOUR"`
	Price           float64 `json:"PRICE"`
	LastUpdate      int64   `json:"LASTUPDATE"`
	Supply          float64 `json:"SUPPLY"`
	Mktcap          float64 `json:"MKTCAP"`
}

//Display - describes the second field of response that
//contains processed data from Raw
type Display struct {
	Change24Hour    string `json:"CHANGE24HOUR"`
	ChangePCT24Hour string `json:"CHANGEPCT24HOUR"`
	Open24Hour      string `json:"OPEN24HOUR"`
	Volume24Hour    string `json:"VOLUME24HOUR"`
	Voulume24HourTo string `json:"VOLUME24HOURTO"`
	High24Hour      string `json:"HIGH24HOUR"`
	Price           string `json:"PRICE"`
	FromSymbol      string `json:"FROMSYMBOL"`
	ToSymbol        string `json:"TOSYMBOL"`
	LastUpdate      string `json:"LASTUPDATE"`
	Supply          string `json:"SUPPLY"`
	Mktcap          string `json:"MKTCAP"`
}
