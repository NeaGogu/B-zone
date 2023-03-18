package bumbal

import (
	"bzone/backend/internal/models"
	"fmt"
	"reflect"
	"testing"
)

func TestGetZoneListReponse(t *testing.T) {
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NzkwNTA2MTYsImV4cCI6MTY3OTA3OTQxNiwiaXNzIjoic2VwMjAyMzAyLmJ1bWJhbC5ldSIsImxvZ2dlZF9vdXRfdXVpZCI6IjJkdmVzamh0LWk3M3ctcjhjOC1nNW4wLWMyOXNsNmE5M3hjZyIsInV1aWQiOiJqcXVid2oxeC1sdzZ2LXM1cjUtbDE5Ni10Y3JvNmNzMm13MGsiLCJ1c2VyX2lkIjoiMTEiLCJyb2xlX2lkIjoiNSIsImRldmljZV9pZCI6bnVsbCwidXNlcl91dWlkIjoiaXB3dnVjbzctangzbC05NTV3LWgzMXctMnZwZ2hjaW50NWE4IiwidXNlcl9wYXJ0eV9pZCI6IjEifQ.GkfKUWTPFN4_6ZEF7sviceF76iitgZt5vkqUjkA1qhM"

	// TODO mock response from bumbal

	zoneResponse, err := GetZoneListReponse(jwt)
	if err != nil {
		t.Fatal(err)
	}


	for _, zone := range zoneResponse.Items{

		fmt.Println(zone.Name)

		for _, zoneRange := range *zone.ZoneRanges{
			fmt.Printf("\t %+v \n", zoneRange)
		}
	}

}

func TestZoneListResponseToListOfZoneModels(t *testing.T) {

	// mock response
	// zoneResponse := &BumbalZoneListResponse{
	//
	// 	Items: []BumbalZoneModel{
	//
	// 		{
	//
	// 			ZoneRanges: &[]BumbalZoneRangeModel{
	//
	// 				{
	//
	// 					ZipcodeFrom: "1000",
	// 					ZipcodeTo:   "2000",
	// 					IsoCountry: "NL",
	// 				},
	// 				{
	//
	// 					ZipcodeFrom: "3000",
	// 					ZipcodeTo:   "4000",
	// 					IsoCountry: "NL",
	// 				},
	// 				{
	//
	// 					ZipcodeFrom: "5000",
	// 					ZipcodeTo:   "7000",
	// 					IsoCountry: "NL",
	// 				},
	//
	// 			},
	// 		},
	// 		{
	//
	// 			ZoneRanges: &[]BumbalZoneRangeModel{
	//
	// 				{
	//
	// 					ZipcodeFrom: "1000",
	// 					ZipcodeTo:   "2000",
	// 					IsoCountry: "NL",
	// 				},
	// 				{
	//
	// 					ZipcodeFrom: "3000",
	// 					ZipcodeTo:   "4000",
	// 					IsoCountry: "NL",
	// 				},
	// 				{
	//
	// 					ZipcodeFrom: "5000",
	// 					ZipcodeTo:   "7000",
	// 					IsoCountry: "NL",
	// 				},
	//
	// 			},
	// 		},
	// 	},
	// }
	//
	// exceptedZoneList := []models.ZoneModel{
	// 	{
	// 	
	// 		ZoneRanges: []models.ZoneRangeModel{
	// 			{
	// 				ZipcodeFrom: 1000,
	// 				ZipcodeTo:   2000,
	// 				IsoCountry: "NL",
	// 			},
	// 			{
	// 				ZipcodeFrom: 3000,
	// 				ZipcodeTo:   4000,
	// 				IsoCountry: "NL",
	// 			},
	// 			{
	// 				ZipcodeFrom: 5000,
	// 				ZipcodeTo:   7000,
	// 				IsoCountry: "NL",
	// 			},
	// 		},
	// 	},
	// }

		zoneResponse := BumbalZoneListResponse{
		Items: []BumbalZoneModel{
			{
				Id:   "1",
				Name: "zone 1",
				ZoneRanges: &[]BumbalZoneRangeModel{
					{
						ZipcodeFrom: "1000",
						ZipcodeTo:   "2000",
						IsoCountry:  "US",
					},
				},
			},
			{
				Id:   "2",
				Name: "zone 2",
				ZoneRanges: &[]BumbalZoneRangeModel{
					{
						ZipcodeFrom: "3000",
						ZipcodeTo:   "4000",
						IsoCountry:  "CA",
					},
				},
			},
		},
	}

	expectedZoneList := []models.ZoneModel{
		{
			ZoneRanges:   []models.ZoneRangeModel{{ZoneRangeId: 0, ZipcodeFrom: 1000, ZipcodeTo: 2000, IsoCountry: "US"}},
			ZoneFuelCost: 0,
			ZoneDrivingTime: 0,
		},
		{
			ZoneRanges:   []models.ZoneRangeModel{{ZoneRangeId: 0, ZipcodeFrom: 3000, ZipcodeTo: 4000, IsoCountry: "CA"}},
			ZoneFuelCost: 0,
			ZoneDrivingTime: 0,
		},
	}

	zoneList, err := zoneResponse.ToRangeModelList()
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if !reflect.DeepEqual(expectedZoneList, zoneList) {
		t.Errorf("expected zone list %v, but got %v",expectedZoneList, zoneList)
	}


}
