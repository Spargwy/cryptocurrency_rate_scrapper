package scrapper

import (
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
