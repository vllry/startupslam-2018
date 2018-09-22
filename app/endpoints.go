package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type apiResponse struct {
	Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("MAKE_WORK") != "" {
		for x := 0; x <= 1000; x++ {
			for y := 0; y <= 100; y++ {
				fmt.Println(x*y)
			}
		}
	}

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
