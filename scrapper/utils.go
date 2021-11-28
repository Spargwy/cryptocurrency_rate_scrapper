package scrapper

import (
	"cryptorate/storage"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//CheckParamsAvailability - checking for unavailable params and
//availability of necessery params
func CheckParamsAvailability(params url.Values) (string, bool) {
	if len(params["fsyms"]) == 0 {
		return "", false
	} else if len(params["fsyms"]) == 1 && params["fsyms"][0] == "" {
		return "", false
	}
	if len(params["tsyms"]) == 0 {
		return "", false
	} else if len(params["tsyms"]) == 1 && params["tsyms"][0] == "" {
		return "", false
	}
	availableParams := ParseAvailableParams()
	for param, paramWords := range params {
		for _, paramWord := range paramWords {
			if !elementInArray(paramWord, availableParams[param]) {
				return paramWord, true
			}
		}
	}
	return "", true
}

//ParseAvailableParams - get params from env variables
func ParseAvailableParams() map[string][]string {
	params := make(map[string][]string)
	availableFsyms := os.Getenv("AVAILABLE_FSYMS")
	availableTsyms := os.Getenv("AVAILABLE_TSYMS")

	availableFsyms = strings.ReplaceAll(availableFsyms, " ", "")
	availableTsyms = strings.ReplaceAll(availableTsyms, " ", "")

	params["fsyms"] = strings.Split(availableFsyms, ",")
	params["tsyms"] = strings.Split(availableTsyms, ",")
	return params
}

//ParseRequestParams - separates params that included in one
//place of array
func ParseRequestParams(params url.Values) url.Values {
	for param, paramWords := range params {
		var separatedParams []string
		for _, paramWord := range paramWords {
			paramWord = strings.ReplaceAll(paramWord, " ", "")
			separatedParams = append(separatedParams, strings.Split(paramWord, ",")...)
		}
		params[param] = separatedParams
	}
	return params
}

func elementInArray(element string, array []string) bool {
	elementInArray := false
	for i := range array {
		if array[i] == element {
			elementInArray = true
			break
		}

	}
	return elementInArray
}

//responseProcessing process the response, edit source response struct and return it
func responseProcessing(res *http.Response, saveToDB bool) (body []byte, cryptocurrency storage.PriceMultiFull, err error) {
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

//structureProcessing - conversions response structure into
//Raw and Display structures
func structureProcessing(cryptocurrency storage.PriceMultiFull) (r, d []byte, err error) {
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
		log.Print("Marshal raw error: ", err)
		return
	}
	d, err = json.Marshal(display)
	if err != nil {
		log.Print("Marshal display error: ", err)
		return
	}
	//Build necessary response from raw and display
	return r, d, err
}

//because rows in db contains response with all available
//params, we need cut all except of structures which correspond
//params from user's request
func processingStructureFromDB(raw map[string]map[string]storage.Raw,
	display map[string]map[string]storage.Display,
	params url.Values) (map[string]map[string]storage.Raw,
	map[string]map[string]storage.Display) {
	//Delete all unnecessary response parts
	for fsym := range raw {
		if !elementInArray(fsym, params["fsyms"]) {
			delete(raw, fsym)
			continue
		}
		for tsym := range raw[fsym] {
			if !elementInArray(tsym, params["tsyms"]) {
				delete(raw[fsym], tsym)
			}
		}
	}
	for fsym := range display {
		if !elementInArray(fsym, params["fsyms"]) {
			delete(display, fsym)
			continue
		}
		for tsym := range display[fsym] {
			if !elementInArray(tsym, params["tsyms"]) {
				delete(display[fsym], tsym)
			}
		}
	}
	return raw, display
}

//responseFromStructure creating json response from two different
//parts of one response - RAW and DISPLAY
func responseFromStructure(r []byte, d []byte) []byte {
	//Build necessary response from raw and display
	processedResponse, err := json.Marshal(storage.PriceMultiFull{Raw: r, Display: d})
	if err != nil {
		log.Print("Marshall error: ", err)
	}
	return processedResponse
}
