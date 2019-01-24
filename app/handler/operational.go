package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JAG-UK/numerologySolver/structures"
)

func GetValueOfWord(w http.ResponseWriter, r *http.Request) {
	var conversionQuery = new(structures.NumerologyQuery)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(conversionQuery); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	//TODO look up in the precomp table?
	numerologicalValue := BasicConversion(conversionQuery.Word)
	responseBody := fmt.Sprintf("Numerological value of %s: %v\n", conversionQuery.Word, numerologicalValue)
	SuccessResponse(w, http.StatusOK, map[string]string{"SUCCESS": responseBody})
}

func GetAllWordsOfValue(w http.ResponseWriter, r *http.Request) {
	var conversionQuery = new(structures.TexterologyQuery)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(conversionQuery); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	//TODO fix the client to pass real json so that this can be a uint32, not a string with later explicit conversion
	valv, err := strconv.ParseUint(conversionQuery.Number, 10, 32)
	if err != nil {
		//fmt.Printf("**DBG** Cannot convert '%s' to uint", conversionQuery.Number)
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	vall, err := strconv.ParseUint(conversionQuery.Length, 10, 32)
	if err != nil {
		//fmt.Printf("**DBG** Cannot convert '%s' to uint", conversionQuery.Number)
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	words := FindAll(uint32(valv), uint32(vall))
	responseBody := fmt.Sprintf("There are %v known words with numerological value %v: %v\n", len(words), conversionQuery.Number, words)

	SuccessResponse(w, http.StatusOK, map[string]string{"SUCCESS": responseBody})
}
