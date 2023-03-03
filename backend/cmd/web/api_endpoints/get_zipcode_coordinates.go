package api_endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetAreaPoint(w http.ResponseWriter, r *http.Request) {

	type Codes struct {
		Pc4         string `json:"pc4"`
		Coordinates any    `json:"coordinates"`
	}

	type Nl struct {
		Nl []Codes `json:"nl"`
	}

	type CodesRequested struct {
		Pc4Codes []string `json:"pc4"`
	}

	var path = "./data/nlnasty.json"
	jsonNl, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	fmt.Println("Successfully Opened newnl.json")
	defer jsonNl.Close()

	byteValue, err := io.ReadAll(jsonNl)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	var nl Nl

	err = json.Unmarshal(byteValue, &nl)
	if err != nil {
		fmt.Println(err.Error())
		panic("unmarshal")
	}

	var reqCode CodesRequested
	err = json.NewDecoder(r.Body).Decode(&reqCode)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pc4 := reqCode.Pc4Codes
	var checkZipcodes []string

	var found bool
	var testCode string

	for i := 0; i < len(pc4); i++ {
		found = false
		testCode = pc4[i]

		for i := 0; i < len(nl.Nl); i++ {
			if nl.Nl[i].Pc4 == testCode {
				found = true
			}
		}

		if found == false {
			checkZipcodes = append(checkZipcodes, testCode)
		}
	}

	if len(checkZipcodes) > 0 {
		json.NewEncoder(w).Encode(append(checkZipcodes, "not in the dataset"))
		return
	}

	var coordinates []any

	for i := 0; i < len(nl.Nl); i++ {
		for _, zipcode := range pc4 {
			if nl.Nl[i].Pc4 == zipcode {
				coordinates = append(coordinates, nl.Nl[i].Coordinates)
				w.WriteHeader(http.StatusOK)
				fmt.Printf("Coordinates for zipcode %s were successfully collected!\n", zipcode)
			}
		}
	}

	json.NewEncoder(w).Encode(coordinates)
	fmt.Println("All coordinates were collected and returned!")

	w.WriteHeader(http.StatusNotFound)
	return
}
