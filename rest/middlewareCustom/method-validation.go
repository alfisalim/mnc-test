package middlewareCustom

import (
	"encoding/json"
	"mnc-test/util"

	//"fmt"
	//"mnc-test/model"
	"log"
	"net/http"
)

func GETMethodValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			response := util.ResponseBuilder(http.StatusMethodNotAllowed, "Only Method GET that allowed", nil)

			jsonResp, _ := json.Marshal(response)
			log.Println("Method not allowed")

			w.Write(jsonResp)
		}

		next.ServeHTTP(w,r)
	})
}
