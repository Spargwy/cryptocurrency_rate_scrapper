package scrapper

import (
	"net/url"
	"os"
	"strings"
)

func CheckParamsAvailability(params url.Values) string {
	availableParams := ParseAvailableParams()
	for param, paramWords := range params {
		for _, paramWord := range paramWords {
			if !elementInArray(paramWord, availableParams[param]) {
				return paramWord
			}
		}
	}
	return ""
}

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

//In request we have map that has as key array of strings
//But this strings can be locate in same array
//  so we cant check separate parameter
//Here we'r solving this problems
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
