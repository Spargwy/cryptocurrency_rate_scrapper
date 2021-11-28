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

//GetCurrentPrice get real-time data from cryptocompare.com api, processing response and return it
func GetCurrentPrice(requestParams url.Values, saveToDB bool, periodic bool) (responseBody []byte, err error) {
	//If request failed - execute data from db
	res, err := setupAndExecuteRequest(requestParams)
	if err != nil {
		log.Print("setupAndExecuteRequest error: ", err)

		//If it is cron task we don't need to execute data from db
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
		log.Print("Unmarshall error: ", err)
		return
	}
	err = json.Unmarshal(display, &displayMap)
	if err != nil {
		log.Print("Unmarshall error: ", err)
		return
	}

	//Delete all unnecessary structs
	for fsym := range rawMap {
		if !elementInArray(fsym, requestParams["fsyms"]) {
			delete(rawMap, fsym)
			continue
		}
		for tsym := range rawMap[fsym] {
			if !elementInArray(tsym, requestParams["tsyms"]) {
				delete(rawMap[fsym], tsym)
			}
		}
	}
	for fsym := range displayMap {
		if !elementInArray(fsym, requestParams["fsyms"]) {
			delete(displayMap, fsym)
			continue
		}
		for tsym := range displayMap[fsym] {
			if !elementInArray(tsym, requestParams["tsyms"]) {
				delete(displayMap[fsym], tsym)
			}
		}
	}
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

//responseProcessing process the response, edit source response struct and return it
func responseProcessing(res *http.Response, saveToDB bool) (body []byte, cryptocurrency storage.PriceMultyFull, err error) {
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &cryptocurrency)
	if err != nil {
		log.Print("Unmarshall error: ", err)
		return
	}
	raw, display, err := structureProcessing(cryptocurrency)
	if err != nil {
		log.Print("Structure processing error: ", err)
	}
	body = responseFromStructure(raw, display)
	err = json.Unmarshal(body, &cryptocurrency)
	if err != nil {
		log.Print("Unmarshall respons to cryptocurrency error: ", err)
	}

	return
}

//structureProcessing for cutting trash structures fields from response
func structureProcessing(cryptocurrency storage.PriceMultyFull) (r, d []byte, err error) {
	var raw map[string]map[string]storage.Raw
	var display map[string]map[string]storage.Display

	//Unmarshall to cuted structures that we need as map
	err = json.Unmarshal(cryptocurrency.Raw, &raw)
	if err != nil {
		log.Print("Unmarshall raw error: ", err)
		return
	}
	err = json.Unmarshal(cryptocurrency.Display, &display)
	if err != nil {
		log.Print("Unmarshall display error: ", err)
	}
	//Marshall back to json
	r, err = json.Marshal(raw)
	if err != nil {
		log.Print("MARSHAL MAP ERROR: ", err)
		return
	}
	d, err = json.Marshal(display)
	if err != nil {
		log.Print("MARSHAL MAP ERROR: ", err)
		return
	}
	//Build necessary response from raw and display
	return r, d, err
}

//responseFromStructure creating json response from two different
//parts of one response - RAW and DISPLAY
func responseFromStructure(r []byte, d []byte) []byte {
	//Build necessary response from raw and display
	processedResponse, err := json.Marshal(storage.PriceMultyFull{Raw: r, Display: d})
	if err != nil {
		log.Print("Marshall error: ", err)
	}
	return processedResponse
}
