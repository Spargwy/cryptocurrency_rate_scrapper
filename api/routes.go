package api

import (
	"cryptorate/scrapper"
	"fmt"
	"log"
	"net/http"
)

func price(w http.ResponseWriter, r *http.Request) {
	parsedRequestParams := scrapper.ParseRequestParams(r.URL.Query())
	unavailableParam, necesseryParamIsNotAvailable := scrapper.CheckParamsAvailability(parsedRequestParams)
	if !necesseryParamIsNotAvailable {
		fmt.Fprintf(w, "fsyms or tsyms param is empty or null")
		return
	}
	if unavailableParam != "" {
		log.Printf("Unavailable param: %s", unavailableParam)
		fmt.Fprintf(w, "Unavailable param %s", unavailableParam)
		return
	}
	response, err := scrapper.GetCurrentPrice(r.URL.Query(), false, false)
	if err != nil {
		log.Print("get current price error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func SetupRoutes() {
	http.HandleFunc("/price", price)
}
