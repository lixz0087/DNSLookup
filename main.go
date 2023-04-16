package main

import (
	"DNSLookup/dao"
	"DNSLookup/server"
	"DNSLookup/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dnsLookupDao := dao.CreateDNSLookupDAO()

	http.HandleFunc(utils.ApiPrefix+"/dig", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			errMsg := "ERROR: Method not allowed"
			http.Error(w, errMsg, http.StatusMethodNotAllowed)
			log.Println(errMsg, "API", "/dig", "Method", r.Method)
			return
		}

		digInput := &server.DigInput{}

		err := json.NewDecoder(r.Body).Decode(digInput)
		if err != nil {
			errMsg := fmt.Sprintf("ERROR: can't decode input with error msg: %v", err.Error())
			http.Error(w, errMsg, http.StatusBadRequest)
			log.Println(errMsg)
			return
		}

		log.Printf("INFO: get request on /dig API with input: %v", digInput)

		digResult, err := server.DigHandler(digInput, dnsLookupDao)
		if err != nil {
			http.Error(w, fmt.Sprintf(err.Error()), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		_, err = fmt.Fprintf(w, "DNS lookup return for domain: %v \n%v", digInput.Domain, digResult)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
	})

	http.HandleFunc(utils.ApiPrefix+"/dig/history", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			errMsg := "ERROR: Method not allowed"
			http.Error(w, errMsg, http.StatusMethodNotAllowed)
			log.Println(errMsg, "API", "/dig/history", "Method", r.Method)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		log.Println("INFO: get request on /dig/history API")

		err := json.NewEncoder(w).Encode(server.DigHistoryHandler(dnsLookupDao))
		if err != nil {
			errMsg := fmt.Sprintf("ERROR: can't decode input with error msg: %v", err.Error())
			http.Error(w, errMsg, http.StatusBadRequest)
			log.Println(errMsg)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
