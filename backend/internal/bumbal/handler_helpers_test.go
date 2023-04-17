package bumbal

import (
	"bzone/backend/internal/models"
	"bzone/backend/internal/test"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestRequestBumbalActivity(t *testing.T) {

	t.Run("Request Bumbal Activity", func(t *testing.T) {
		var err error
		var req *http.Request

		req, err = http.NewRequest("PUT", "/test", nil)
		assert.NoError(t, err)
		token, _ := test.CreateAccessTokenString("1")
		req.Header.Add("Authorization", token)
		rr := httptest.NewRecorder()
		body := map[string]string{
			"plot_name": "MyPlot",
			"plot_id":   "30",
		}

		var reqBody []byte
		reqBody, err = json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		_, errR := requestBumbalActivity(rr, req, reqBody, baseUrl)

		assert.Nil(t, errR)
	})
}

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
		// Test cases
		{
			name:         "Empty input slice",
			respModel:    []models.ActivityModelBumbal{},
			expectedResp: []models.ActivityModelBumbal{},
		},
		{
			name: "Input slice with no matching elements",
			respModel: []models.ActivityModelBumbal{
				{
					Id: nil, AddressApplied: nil,
					DepotAddress: &models.AddressModelBumbal{
						Zipcode: &zipcodeBacau, Latitude: &latitudeBacau, Longitude: &longitudeBacau,
					},
				},
				{
					Id: nil, AddressApplied: nil,
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
					Id: nil, AddressApplied: nil, DepotAddress: nil,
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

func TestGetClustersInfo(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		request  *http.Request
		expected ClustersInfo
		err      error
	}{
		{
			name: "Valid request with data",
			request: &http.Request{
				Body: io.NopCloser(strings.NewReader(`{"number_of_clusters":5,"number_of_candidate_clusters":10}`)),
			},
			expected: ClustersInfo{NrClusters: 5, NrCandidateClusters: 10},
			err:      nil,
		},
		{
			name: "Valid request with missing fields",
			request: &http.Request{
				Body: io.NopCloser(strings.NewReader(`{}`)),
			},
			expected: ClustersInfo{},
			err:      nil,
		},
		{
			name: "Invalid request with malformed JSON",
			request: &http.Request{
				Body: io.NopCloser(strings.NewReader(`{"number_of_clusters":5,"number_of_candidate_clusters":}`)),
			},
			expected: ClustersInfo{},
			err:      &json.SyntaxError{Offset: 46},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clustersInfo, err := getClustersInfo(tc.request)

			if !reflect.DeepEqual(clustersInfo, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, clustersInfo)
			}

			if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
				t.Errorf("Expected error of type %T, but got %T", tc.err, err)
			}
		})
	}
}

func TestGetZonesInfo(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		request  *http.Request
		expected ZonesInfo
		err      error
	}{
		{
			name: "Valid request with data",
			request: &http.Request{
				Body: io.NopCloser(strings.NewReader(`{"number_of_zones":5,"number_of_generations":10,
													"maximum_runtime":10}`)),
			},
			expected: ZonesInfo{NZones: 5, NGenerations: 10, MaxDuration: 10},
			err:      nil,
		},
		{
			name: "Valid request with missing fields",
			request: &http.Request{
				Body: io.NopCloser(strings.NewReader(`{}`)),
			},
			expected: ZonesInfo{},
			err:      nil,
		},
		{
			name: "Invalid request with malformed JSON",
			request: &http.Request{
				Body: io.NopCloser(strings.NewReader(`{"number_of_zones":5,"number_of_generations":}`)),
			},
			expected: ZonesInfo{},
			err:      &json.SyntaxError{Offset: 43},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clustersInfo, err := getZonesInfo(tc.request)

			if !reflect.DeepEqual(clustersInfo, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, clustersInfo)
			}

			if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
				t.Errorf("Expected error of type %T, but got %T", tc.err, err)
			}
		})
	}
}

func TestGetResponseData(t *testing.T) {
	// Define testing variables
	zipcodePloiesti, latitudePloiesti, longitudePloiesti := "100051", "44.933334", "26.033333"
	zipcodeSurani, latitudeSurani, longitudeSurani := "107545", "45.1955", "26.1853"

	// Test cases
	testCases := []struct {
		name     string
		response *http.Response
		expected models.ActivityListResponseBumbal
		err      error
	}{
		{
			name: "Valid response with items",
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(`{
                    "items": [
                        {
                            "id": null,
                            "address_applied": {
                                "zipcode": "100051",
                                "latitude": "44.933334",
                                "longitude": "26.033333"
                            },
                            "depot_address": {
                                "zipcode": "107545",
                                "latitude": "45.1955",
                                "longitude": "26.1853"
                            }
                        }
                    ],
                    "count_filtered": 1,
                    "count_unfiltered": 1,
                    "count_limited": 1
                }`)),
			},
			expected: models.ActivityListResponseBumbal{
				Items: &[]models.ActivityModelBumbal{
					{
						Id: nil,
						AddressApplied: &models.AddressAppliedModelBumbal{
							Zipcode:   &zipcodePloiesti,
							Latitude:  &latitudePloiesti,
							Longitude: &longitudePloiesti,
						},
						DepotAddress: &models.AddressModelBumbal{
							Zipcode:   &zipcodeSurani,
							Latitude:  &latitudeSurani,
							Longitude: &longitudeSurani,
						},
					},
				},
				CountFiltered:   1,
				CountUnfiltered: 1,
				CountLimited:    1,
			},
			err: nil,
		},
		{
			name: "Valid response without items",
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(`{
                    "items": null,
                    "count_filtered": 0,
                    "count_unfiltered": 0,
                    "count_limited": 0
                }`)),
			},
			expected: models.ActivityListResponseBumbal{
				Items:           nil,
				CountFiltered:   0,
				CountUnfiltered: 0,
				CountLimited:    0,
			},
			err: nil,
		},
		{
			name: "Invalid response with non-JSON data",
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`not json data`)),
			},
			expected: models.ActivityListResponseBumbal{},
			err:      &json.SyntaxError{},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			respModel, err := getResponseData(tc.response)

			if !reflect.DeepEqual(respModel, tc.expected) {
				t.Errorf("Expected response model to be %+v, but got %+v", tc.expected, respModel)
			}

			if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
				t.Errorf("Expected error type to be %T, but got %T", tc.err, err)
			}
		})
	}
}
