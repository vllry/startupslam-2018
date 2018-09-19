package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiResponse struct {
	Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	responseStruct := apiResponse{
		"API is running",
	}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal error")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(response))
}
