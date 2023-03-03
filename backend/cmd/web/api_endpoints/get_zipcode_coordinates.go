package api_endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Codes Data structure containing the zipcode and its coordinates
type Codes struct {
	Pc4         string `json:"pc4"`
	Coordinates any    `json:"coordinates"`
}

// Nl Data structure for all the zipcodes
type Nl struct {
	Nl []Codes `json:"nl"`
}

// CodesRequested Data structure for storing the requested zipcodes
type CodesRequested struct {
	Pc4Codes []string `json:"pc4"`
}

// openFile
//
//	@Description: Opens the file containing all the zipcodes and returns a data structure
//				  containing them and the request code
//	@param path
//	@param w
//	@param r
//	@return Nl
//	@return CodesRequested
func openFile(path string, w http.ResponseWriter, r *http.Request) (Nl, CodesRequested) {
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
	}

	return nl, reqCode
}

// checkZipcodesValidity
//
//	@Description: Boolean function that checks if the inputted zipcodes are valid
//	@param pc4
//	@param nl
//	@param w
//	@return bool
func checkZipcodesValidity(pc4 []string, nl Nl, w http.ResponseWriter) bool {
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
		return false
	} else {
		return true
	}
}

// returnCoordinates
//
//	@Description: Encodes the coordinates of the requested zipcodes
//	@param pc4
//	@param nl
//	@param w
func returnCoordinates(pc4 []string, nl Nl, w http.ResponseWriter) {
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

// GetAreaPoint
//
//	@Description: This is the main function of this API endpoint
//	@param w
//	@param r
func GetAreaPoint(w http.ResponseWriter, r *http.Request) {

	var path = "./data/nlnasty.json"

	nl, reqCode := openFile(path, w, r)
	pc4 := reqCode.Pc4Codes

	if checkZipcodesValidity(pc4, nl, w) == false {
		fmt.Println("Some zipcodes are not in the dataset!")
		return
	} else {
		returnCoordinates(pc4, nl, w)
	}

	return
}
