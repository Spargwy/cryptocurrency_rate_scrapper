package scrapper

import (
	"cryptorate/storage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/mitchellh/mapstructure"
)

func GetCurrentPrice(requestParams url.Values, saveToDB bool) (cryptocurrency storage.PriceMultyFull, err error) {
	apiKey := os.Getenv("API_KEY")
	url := "https://min-api.cryptocompare.com/data/pricemultifull"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	values := req.URL.Query()
	values.Add("api_key", apiKey)
	for param, paramWords := range requestParams {
		for _, paramWord := range paramWords {
			values.Add(param, paramWord)
		}
	}
	req.URL.RawQuery = values.Encode()
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &cryptocurrency)
	if err != nil {
		log.Print("Unmarshall error: ", err)
		return
	}
	err = storage.Insert(cryptocurrency)
	if err != nil {
		log.Print("Insert error: ", err)
		return
	}
	// structureProcessing(cryptocurrency)
	return
}

func StructureProcessing(cryptocurrency storage.PriceMultyFull) {
	var raw map[string]map[string]map[string]interface{}

	err := json.Unmarshal(cryptocurrency.Raw, &raw)
	if err != nil {
		log.Print("Unmarshall error: ", err)
	}
	var crc storage.Currency
	for cryptocurrency := range raw {
		for _, value := range raw[cryptocurrency] {
			err := mapstructure.Decode(value, &crc)
			if err != nil {
				log.Print("DECODE mapstructure ERROR: ", err)
			}
		}
	}
	fmt.Println(crc)
}
