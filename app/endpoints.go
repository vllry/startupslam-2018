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
	// Pretend this endpoint is computationally heavy.
	for x := 0; x <= 500; x++ {
		for y := 0; y <= 100; y++ {
			fmt.Println(x*y)
		}
	}

	responseStruct := apiResponse{
		"API is running, env config is: " + os.Getenv("RETURN_MSG"),
	}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal error")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(response))
}
