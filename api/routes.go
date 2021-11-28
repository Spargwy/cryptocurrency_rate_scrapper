package api

import "net/http"

//SetupRoutes - setups available endpoints
func SetupRoutes() {
	http.HandleFunc("/price", price)
}
