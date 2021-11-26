package scrapper

import (
	"cryptorate/storage"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func GetCurrentPrice(requestParams url.Values, saveToDB bool) (response []byte, err error) {
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
	var cryptocurrency storage.PriceMultyFull
	err = json.Unmarshal(body, &cryptocurrency)
	if err != nil {
		log.Print("Unmarshall error: ", err)
		return
	}
	response, err = StructureProcessing(cryptocurrency)
	if err != nil {
		log.Print("Structure processing error: ", err)
	}
	err = json.Unmarshal(response, &cryptocurrency)
	if err != nil {
		log.Print("Unmarshall respons to cryptocurrency error: ", err)
	}
	if saveToDB {
		if cryptocurrency.Display == nil || cryptocurrency.Raw == nil {
			log.Print("cryptocurrencies is nil")
			return
		}
		err = storage.Insert(cryptocurrency)
		if err != nil {
			log.Print("Insert error: ", err)
			return
		}
	}
	return
}

//StructureProcessing for cutting trash structures fields from response
func StructureProcessing(cryptocurrency storage.PriceMultyFull) (processedResponse []byte, err error) {
	var raw map[string]map[string]storage.CurrencyRaw
	var display map[string]map[string]storage.CurrencyDisplay

	//Unmarshall to cuted structures that we need as map
	err = json.Unmarshal(cryptocurrency.Raw, &raw)
	if err != nil {
		log.Print("Unmarshall raw error: ", err)
		return nil, err
	}
	err = json.Unmarshal(cryptocurrency.Display, &display)
	if err != nil {
		log.Print("Unmarshall display error: ", err)
	}
	//Marshall back to json
	r, err := json.Marshal(raw)
	if err != nil {
		log.Print("MARSHAL MAP ERROR: ", err)
		return nil, err
	}
	d, err := json.Marshal(display)
	if err != nil {
		log.Print("MARSHAL MAP ERROR: ", err)
		return nil, err
	}
	//Build necessary response from raw and display
	processedResponse, err = json.Marshal(storage.PriceMultyFull{Raw: r, Display: d})
	return processedResponse, err
}
