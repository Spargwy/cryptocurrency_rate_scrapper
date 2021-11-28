package scrapper

import (
	"cryptorate/storage"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

//GetCurrentPrice get real-time data from cryptocompare.com api, processing response and return it
func GetCurrentPrice(requestParams url.Values, saveToDB bool, periodic bool) (responseBody []byte, err error) {
	//If request failed - extract data from db
	res, err := setupAndExecuteRequest(requestParams)
	if err != nil {
		log.Print("setupAndExecuteRequest error: ", err)

		//If it is a cron task we don't need to extract data from db
		if !periodic {
			responseBody, err = getCurrencyFromDB(requestParams)
			if err != nil {
				log.Print("getCurrencyFromDB error: ", err)
			}
		}
		return
	}

	defer res.Body.Close()
	if err != nil {
		log.Print("Cant select data: ", err)
	}
	responseBody, cryptocurrency, err := responseProcessing(res, saveToDB)
	if err != nil {
		log.Print("responseProcessing error: ", err)
		return
	}
	if saveToDB {
		err = storage.Insert(cryptocurrency)
		if err != nil {
			log.Print("Insert error: ", err)
			return
		}
	}
	return responseBody, nil
}

//getCurrencyFromDB - executes data from db if request to api above is not available
func getCurrencyFromDB(requestParams url.Values) (body []byte, err error) {
	requestParams = ParseRequestParams(requestParams)
	raw, display, err := storage.Select()
	if err != nil {
		log.Print("Select from db error: ", err)
	}
	var rawMap map[string]map[string]storage.Raw
	var displayMap map[string]map[string]storage.Display
	err = json.Unmarshal(raw, &rawMap)
	if err != nil {
		log.Print("Unmarshall to rawMap error: ", err)
		return
	}
	err = json.Unmarshal(display, &displayMap)
	if err != nil {
		log.Print("Unmarshall to displayMap error: ", err)
		return
	}

	//Delete all unnecessary response parts
	rawMap, displayMap = processingStructureFromDB(rawMap, displayMap, requestParams)
	raw, err = json.Marshal(rawMap)
	if err != nil {
		log.Print("rawMap Marshal error: ", err)
		return
	}
	display, err = json.Marshal(displayMap)
	if err != nil {
		log.Print("displayMap Marshal error: ", err)
		return
	}
	body, err = json.Marshal(storage.PriceMultyFull{Raw: raw, Display: display})
	return
}

//setupAndExecuteRequest - creates request that included necessery params and executing it
func setupAndExecuteRequest(requestParams url.Values) (res *http.Response, err error) {
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
	res, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
