package genetic

import (
	"bzone/backend/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateMTSPInstance(t *testing.T) {
	tests := []struct {
		name       string
		activities []models.ActivityModelBumbal
		nRoutes    int
		want       MDMTSPInstance
	}{
		{
			name:       "no activities",
			activities: []models.ActivityModelBumbal{},
			nRoutes:    42,
			want: MDMTSPInstance{
				Activities: []Pos{},
				Depots:     []Pos{},
				NRoutes:    42,
			},
		},
		{
			name: "one activity",
			activities: []models.ActivityModelBumbal{
				*makeActivity(t, "1", "2", "3456AR", "4", "5", "6789HS"),
			},
			nRoutes: 42,
			want: MDMTSPInstance{
				Activities: []Pos{{1, 2, 3456}},
				Depots:     []Pos{{4, 5, 6789}},
				NRoutes:    42,
			},
		},
		{
			name: "multiple activities, one depot",
			activities: []models.ActivityModelBumbal{
				*makeActivity(t, "1", "2", "3456AB", "4", "5", "6789XZ"),
				*makeActivity(t, "10", "20", "3042PO", "4", "5", "6789XZ"),
			},
			nRoutes: 42,
			want: MDMTSPInstance{
				Activities: []Pos{{1, 2, 3456}, {10, 20, 3042}},
				Depots:     []Pos{{4, 5, 6789}},
				NRoutes:    42,
			},
		},
		{
			name: "multiple activities, multiple depots",
			activities: []models.ActivityModelBumbal{
				*makeActivity(t, "1", "2", "3456AB", "4", "5", "6789ZX"),
				*makeActivity(t, "10", "20", "3012QW", "40", "50", "6089AS"),
			},
			nRoutes: 42,
			want: MDMTSPInstance{
				Activities: []Pos{{1, 2, 3456}, {10, 20, 3012}},
				Depots:     []Pos{{4, 5, 6789}, {40, 50, 6089}},
				NRoutes:    42,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateMDMTSPInstance(tt.activities, tt.nRoutes)
			assert.Equal(t, tt.want, got)
		})
	}
}

func makeActivity(t testing.TB, actLat string, actLon string, actZip string, depotLat string, depotLon string, depotZip string) *models.ActivityModelBumbal {
	t.Helper()

	activity := models.NewActivityModel()
	address := models.NewAddressAppliedModel()
	address.SetLatitude(actLat)
	address.SetLongitude(actLon)
	address.SetZipcode(actZip)
	activity.SetAddressApplied(*address)

	depotAddress := models.NewAddressModel()
	depotAddress.SetLatitude(depotLat)
	depotAddress.SetLongitude(depotLon)
	depotAddress.SetZipcode(depotZip)
	activity.SetDepotAddress(*depotAddress)

	return activity
}

func TestSolution2ZoneModels(t *testing.T) {
	tests := []struct {
		name string
		sol  Solution
		want []models.ZoneModel
	}{
		{"no zones", Solution{Routes: []Route{}},
			[]models.ZoneModel{}},
		{"1 zone, 0 zips", Solution{Routes: []Route{{Activities: []Pos{}}}},
			[]models.ZoneModel{{ZoneRanges: []models.ZoneRangeModel{}}}},
		{"1 zone, 1 zip", Solution{Routes: []Route{{Activities: []Pos{{Zipcode: 1}}}}},
			[]models.ZoneModel{{ZoneRanges: []models.ZoneRangeModel{{ZipcodeFrom: 1, ZipcodeTo: 1}}}}},
		{"1 zone, 2 zips", Solution{Routes: []Route{{Activities: []Pos{{Zipcode: 1}, {Zipcode: 2}}}}},
			[]models.ZoneModel{{ZoneRanges: []models.ZoneRangeModel{{ZipcodeFrom: 1, ZipcodeTo: 1}, {ZipcodeFrom: 2, ZipcodeTo: 2}}}}},
		{"2 zones, 0 zips", Solution{Routes: []Route{{Activities: []Pos{}}, {Activities: []Pos{}}}},
			[]models.ZoneModel{{ZoneRanges: []models.ZoneRangeModel{}}, {ZoneRanges: []models.ZoneRangeModel{}}}},
		{"2 zones, 2 same zips", Solution{Routes: []Route{{Activities: []Pos{{Zipcode: 1}}}, {Activities: []Pos{{Zipcode: 1}}}}},
			[]models.ZoneModel{{ZoneRanges: []models.ZoneRangeModel{{ZipcodeFrom: 1, ZipcodeTo: 1}}}, {ZoneRanges: []models.ZoneRangeModel{{ZipcodeFrom: 1, ZipcodeTo: 1}}}}},
		{"2 zones, 2 diff zips", Solution{Routes: []Route{{Activities: []Pos{{Zipcode: 1}}}, {Activities: []Pos{{Zipcode: 2}}}}},
			[]models.ZoneModel{{ZoneRanges: []models.ZoneRangeModel{{ZipcodeFrom: 1, ZipcodeTo: 1}}}, {ZoneRanges: []models.ZoneRangeModel{{ZipcodeFrom: 2, ZipcodeTo: 2}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution2ZoneModels(tt.sol)
			assert.Equal(t, tt.want, got)
		})
	}
}
