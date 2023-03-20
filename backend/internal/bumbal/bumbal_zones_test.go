package bumbal

import (
	"bzone/backend/internal/models"
	"fmt"
	"reflect"
	"testing"
)

func TestGetZoneListReponse(t *testing.T) {

	t.Run("Happy path", func(t *testing.T) {

		t.Skip("Skipping test for now, need to mock the response from bumbal")

		jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NzkwNTA2MTYsImV4cCI6MTY3OTA3OTQxNiwiaXNzIjoic2VwMjAyMzAyLmJ1bWJhbC5ldSIsImxvZ2dlZF9vdXRfdXVpZCI6IjJkdmVzamh0LWk3M3ctcjhjOC1nNW4wLWMyOXNsNmE5M3hjZyIsInV1aWQiOiJqcXVid2oxeC1sdzZ2LXM1cjUtbDE5Ni10Y3JvNmNzMm13MGsiLCJ1c2VyX2lkIjoiMTEiLCJyb2xlX2lkIjoiNSIsImRldmljZV9pZCI6bnVsbCwidXNlcl91dWlkIjoiaXB3dnVjbzctangzbC05NTV3LWgzMXctMnZwZ2hjaW50NWE4IiwidXNlcl9wYXJ0eV9pZCI6IjEifQ.GkfKUWTPFN4_6ZEF7sviceF76iitgZt5vkqUjkA1qhM"

		// TODO mock response from bumbal

		zoneResponse, err := GetZoneListReponse(jwt)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(zoneResponse, err)

		for _, zone := range zoneResponse.Items {

			fmt.Println("hellllooooo")
			fmt.Println(zone.Name)

			for _, zoneRange := range *zone.ZoneRanges {
				fmt.Printf("\t %+v \n", zoneRange)
			}
		}

	})

	t.Run("Invalid jwt - empty string", func(t *testing.T) {

		jwt := ""

		_, err := GetZoneListReponse(jwt)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

	})
}

func TestZoneListResponseToListOfZoneModels(t *testing.T) {
	testCases := []struct {
		name             string
		mockZoneResponse BumbalZoneListResponse
		expectedZoneList []models.ZoneModel
		expectedError    bool
	}{
		{
			name: "Happy path",
			mockZoneResponse: BumbalZoneListResponse{
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
			},
			expectedZoneList: []models.ZoneModel{
				{
					ZoneRanges: []models.ZoneRangeModel{
						{ZoneRangeId: 0, ZipcodeFrom: 1000, ZipcodeTo: 2000, IsoCountry: "US"},
					},
					ZoneFuelCost:    0,
					ZoneDrivingTime: 0,
				},
				{
					ZoneRanges: []models.ZoneRangeModel{
						{ZoneRangeId: 0, ZipcodeFrom: 3000, ZipcodeTo: 4000, IsoCountry: "CA"},
					},
					ZoneFuelCost:    0,
					ZoneDrivingTime: 0,
				},
			},
			expectedError: false,
		},
		{
			name: "Invalid zipcode from, not a number",
			mockZoneResponse: BumbalZoneListResponse{
				Items: []BumbalZoneModel{
					{
						ZoneRanges: &[]BumbalZoneRangeModel{
							{
								ZipcodeFrom: "invalid",
								ZipcodeTo:   "54321",
								IsoCountry:  "US",
							},
						},
					},
				},
			},
			expectedZoneList: nil,
			expectedError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneList, err := tc.mockZoneResponse.ToRangeModelList()

			// check for error first
			if tc.expectedError && (err == nil) {
				t.Error("Expected error but got nil")
			} else if !tc.expectedError && err != nil {
				t.Errorf("Expected nil error but got %v", err)
			}

			if tc.expectedError && !reflect.DeepEqual(tc.expectedZoneList, zoneList) {
				t.Errorf("expected zone list %v, but got %v", tc.expectedZoneList, zoneList)
			}

		})
	}

}

func TestFlowFromBumbalRequestToPlotModel(t *testing.T) {

	// skip as GetZoneListReponse is not mocked and always calls bumbal's API
	t.Skip()

	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NzkyNTIzMDIsImV4cCI6MTY3OTI4MTEwMiwiaXNzIjoic2VwMjAyMzAyLmJ1bWJhbC5ldSIsImxvZ2dlZF9vdXRfdXVpZCI6IjkzM2sybGgwLTN0bmYtMzg3MS01NmY1LWxjNTk3dmE4NnRzOCIsInV1aWQiOiJqOGhwYXBrMC1vZm8wLTExYWItbjV4cS1zZjF1aHBzNzAwMDQiLCJ1c2VyX2lkIjoiMTEiLCJyb2xlX2lkIjoiNSIsImRldmljZV9pZCI6bnVsbCwidXNlcl91dWlkIjoiaXB3dnVjbzctangzbC05NTV3LWgzMXctMnZwZ2hjaW50NWE4IiwidXNlcl9wYXJ0eV9pZCI6IjEifQ.OEglHHbh9mNU0VSVgfU3Hca2gfiWLEabP9H_IXiQU1s"

	zoneResponse, err := GetZoneListReponse(jwt)
	if err != nil {
		t.Fatal(err)
	}

	p, err := zoneResponse.ToPlotModel()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", *p)

}
