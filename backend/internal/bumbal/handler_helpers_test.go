package bumbal

import (
	"bzone/backend/internal/models"
	"reflect"
	"testing"
)

func TestFilterResp(t *testing.T) {
	// Define testing variables
	zipcodePloiesti, latitudePloiesti, longitudePloiesti := "100051", "44.933334", "26.033333"
	zipcodeSurani, latitudeSurani, longitudeSurani := "107545", "45.1955", "26.1853"
	zipcodeCraiova, latitudeCraiova, longitudeCraiova := "200051", "44.319305", "23.800678"
	zipcodeCaracal, latitudeCaracal, longitudeCaracal := "253200", "44.115745", "24.342475"
	zipcodeGalati, latitudeGalati, longitudeGalati := "800017", "45.435321", "28.007994"
	zipcodeBacau, latitudeBacau, longitudeBacau := "600087", "46.568825", "26.916025"
	// Define test cases
	tests := []struct {
		name         string
		respModel    []models.ActivityModelBumbal
		expectedResp []models.ActivityModelBumbal
	}{
		// Test cases here...
		{
			name:         "Empty input slice",
			respModel:    []models.ActivityModelBumbal{},
			expectedResp: []models.ActivityModelBumbal{},
		},
		{
			name: "Input slice with no matching elements",
			respModel: []models.ActivityModelBumbal{
				{
					Id:             nil,
					AddressApplied: nil,
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeBacau, Latitude: &latitudeBacau, Longitude: &longitudeBacau,
					},
				},
				{
					Id:             nil,
					AddressApplied: nil,
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeGalati, Latitude: &latitudeGalati, Longitude: &longitudeGalati,
					},
				},
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodeGalati, Latitude: &latitudeGalati, Longitude: &longitudeGalati,
					},
					DepotAddress: nil,
				},
				{
					Id:             nil,
					AddressApplied: nil,
					DepotAddress:   nil,
				},
			},
			expectedResp: []models.ActivityModelBumbal{},
		},
		{
			name: "Input slice with matching elements",
			respModel: []models.ActivityModelBumbal{
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodePloiesti, Latitude: &latitudePloiesti, Longitude: &longitudePloiesti,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeSurani, Latitude: &latitudeSurani, Longitude: &longitudeSurani,
					},
				},
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodeCraiova, Latitude: &latitudeCraiova, Longitude: &longitudeCraiova,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeCaracal, Latitude: &latitudeCaracal, Longitude: &longitudeCaracal,
					},
				},
			},
			expectedResp: []models.ActivityModelBumbal{
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodePloiesti, Latitude: &latitudePloiesti, Longitude: &longitudePloiesti,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeSurani, Latitude: &latitudeSurani, Longitude: &longitudeSurani,
					},
				},
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodeCraiova, Latitude: &latitudeCraiova, Longitude: &longitudeCraiova,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeCaracal, Latitude: &latitudeCaracal, Longitude: &longitudeCaracal,
					},
				},
			},
		},
		{
			name: "Input slice with some matching elements",
			respModel: []models.ActivityModelBumbal{
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodeCraiova, Latitude: &latitudeCraiova, Longitude: &longitudeCraiova,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeSurani, Latitude: &latitudeSurani, Longitude: &longitudeSurani,
					},
				},
				{
					Id:             nil,
					AddressApplied: nil,
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeSurani, Latitude: &latitudeSurani, Longitude: &longitudeSurani,
					},
				},
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodeCraiova, Latitude: &latitudeCraiova, Longitude: &longitudeCraiova,
					},
					DepotAddress: nil,
				},
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodePloiesti, Latitude: &latitudePloiesti, Longitude: &longitudePloiesti,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeBacau, Latitude: &latitudeBacau, Longitude: &longitudeBacau,
					},
				},
			},
			expectedResp: []models.ActivityModelBumbal{
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodeCraiova, Latitude: &latitudeCraiova, Longitude: &longitudeCraiova,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeSurani, Latitude: &latitudeSurani, Longitude: &longitudeSurani,
					},
				},
				{
					Id: nil,
					AddressApplied: &models.AddressAppliedModelBumbal{
						Zipcode: &zipcodePloiesti, Latitude: &latitudePloiesti, Longitude: &longitudePloiesti,
					},
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeBacau, Latitude: &latitudeBacau, Longitude: &longitudeBacau,
					},
				},
			},
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			filteredResp := filterResp(tc.respModel)

			// Check the output
			if len(filteredResp) != len(tc.expectedResp) {
				t.Errorf("Expected %d elements, but got %d", len(tc.expectedResp), len(filteredResp))
			}

			for i := 0; i < len(filteredResp); i++ {
				if !reflect.DeepEqual(filteredResp[i], tc.expectedResp[i]) {
					t.Errorf("Expected %+v, but got %+v", tc.expectedResp[i], filteredResp[i])
				}
			}
		})
	}
}
