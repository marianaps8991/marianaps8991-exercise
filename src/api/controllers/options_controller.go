package controllers

import (
	"api/responses"
	"fmt"
	"net/http"
)

func AddCORSheader(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n\n\n\n\n\n Definido \n\n\n\n\n\n\n")
	w.Header().Set("Access-Control-Allow-Headers:", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.WriteHeader(http.StatusOK)

	responses.JSON(w, http.StatusOK, "all set")
}
