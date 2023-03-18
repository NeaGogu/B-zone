package bumbal

import (
	"bytes"
	"bzone/backend/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

const baseUrl = "https://sep202302.bumbal.eu/api/v2/"
const (
	zonesUrl = "zone"
)

func (b *BumbalZoneListResponse) ToRangeModelList() ([]models.ZoneModel, error) {

	zoneList := make([]models.ZoneModel,0 , len(b.Items))

	for _, bumbalZone := range b.Items {
		// for each bumbal zone create a Zone Model
		var newZone models.ZoneModel

		// for each zone range in the bumbal zone
		for _, bumbalZoneRange := range *bumbalZone.ZoneRanges {

			// create a new models.ZoneRangeModel
			// we only care about the zipcodes and the country
			var newZoneRange models.ZoneRangeModel
			z, err := strconv.Atoi(bumbalZoneRange.ZipcodeFrom)
			if err != nil {
				return nil, err 
			}
			
			newZoneRange.ZipcodeFrom = z
			
			z, err = strconv.Atoi(bumbalZoneRange.ZipcodeTo)
			if err != nil {
				return nil, err 
			}

			newZoneRange.ZipcodeTo = z
			newZoneRange.IsoCountry = bumbalZoneRange.IsoCountry

			// add the new zone range to the zone
			newZone.ZoneRanges = append(newZone.ZoneRanges, newZoneRange)
		}


		zoneList = append(zoneList, newZone)
	}

	return zoneList, nil
}

func GetZoneListReponse(jwt string) (*BumbalZoneListResponse, error) { 

	// get the zones from bumbal
	requestBody := []byte(`{
	  "options": {
	   "include_zone_ranges": true,
	   "include_brands": false
	  }
	}`)

	url := baseUrl + zonesUrl
	
	bodyReader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var resModel BumbalZoneListResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode( &resModel)
	if err != nil {
		return nil, err
	}

	return &resModel, nil
} 
