package api

import (
	"cryptorate/scrapper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func price(w http.ResponseWriter, r *http.Request) {
	parsedRequestParams := scrapper.ParseRequestParams(r.URL.Query())
	unavailableParam := scrapper.CheckParamsAvailability(parsedRequestParams)
	if unavailableParam != "" {
		log.Printf("Unavailable param: %s", unavailableParam)
		fmt.Fprintf(w, "Unavailable param %s", unavailableParam)
		return
	}
	cryptocurrency, err := scrapper.GetCurrentPrice(r.URL.Query(), false)
	if err != nil {
		log.Print("get current price error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(cryptocurrency)
	if err != nil {
		log.Print("MARSHAL ERROR IN price: ", err)
	}
	w.Write(jsonResp)
}

func SetupRoutes() {
	http.HandleFunc("/price", price)
}
