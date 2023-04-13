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

func Test_Solution2ZoneModels(t *testing.T) {
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

func Test_checkGenAlgHyperParamsBounds(t *testing.T) {
	tests := []struct {
		name    string
		params  GenAlgHyperParams
		wantErr bool
	}{
		{
			name:    "all params in bounds",
			params:  GenAlgHyperParams{42, 42, 42, 42, 42, 0.5, 0.5},
			wantErr: false,
		},
		{
			name:    "NOffspring = 0",
			params:  GenAlgHyperParams{0, 42, 42, 42, 42, 0.5, 0.5},
			wantErr: true,
		},
		{
			name:    "NParents = 0",
			params:  GenAlgHyperParams{42, 0, 42, 42, 42, 0.5, 0.5},
			wantErr: true,
		},
		{
			name:    "NGenerations = 0",
			params:  GenAlgHyperParams{42, 42, 0, 42, 42, 0.5, 0.5},
			wantErr: true,
		},
		{
			name:    "TournamentSize = 0",
			params:  GenAlgHyperParams{42, 42, 42, 0, 42, 0.5, 0.5},
			wantErr: true,
		},
		{
			name:    "MaxMutations = 0",
			params:  GenAlgHyperParams{42, 42, 42, 42, 0, 0.5, 0.5},
			wantErr: true,
		},
		{
			name:    "MutationRate < 0",
			params:  GenAlgHyperParams{42, 42, 42, 42, 42, -0.5, 0.5},
			wantErr: true,
		},
		{
			name:    "MutationRate > 1",
			params:  GenAlgHyperParams{42, 42, 42, 42, 42, 1.5, 0.5},
			wantErr: true,
		},
		{
			name:    "CrossoverRate < 0",
			params:  GenAlgHyperParams{42, 42, 42, 42, 42, 0.5, -0.5},
			wantErr: true,
		},
		{
			name:    "CrossoverRate > 1",
			params:  GenAlgHyperParams{42, 42, 42, 42, 42, 0.5, 1.5},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkGenAlgHyperParamsBounds(tt.params)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_parseActivity(t *testing.T) {
	tests := []struct {
		name     string
		activity models.ActivityModelBumbal
		want     Pos
		wantErr  bool
	}{
		{
			name: "valid activity",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "5621 HH",
				"52.37660663424486", "4.8972405493159235", "1012 LH"),
			want:    Pos{51.452446917398824, 5.462562099535168, 5621},
			wantErr: false,
		},
		{
			name: "invalid lat",
			activity: *makeActivity(t, "fhdsalkfjh", "5.462562099535168", "5621 HH",
				"52.37660663424486", "4.8972405493159235", "1012 LH"),
			want:    Pos{},
			wantErr: true,
		},
		{
			name: "invalid lon",
			activity: *makeActivity(t, "51.452446917398824", "dfljsfaljasl", "5621 HH",
				"52.37660663424486", "4.8972405493159235", "1012 LH"),
			want:    Pos{},
			wantErr: true,
		},
		{
			name: "zip too short",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "5",
				"52.37660663424486", "4.8972405493159235", "1012 LH"),
			want:    Pos{},
			wantErr: true,
		},
		{
			name: "invalid zip",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "kfsadkhl",
				"52.37660663424486", "4.8972405493159235", "1012 LH"),
			want:    Pos{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseActivity(tt.activity)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_parseDepot(t *testing.T) {
	tests := []struct {
		name     string
		activity models.ActivityModelBumbal
		want     Pos
		wantErr  bool
	}{
		{
			name: "valid activity",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "5621 HH",
				"52.37660663424486", "4.8972405493159235", "1012 LH"),
			want:    Pos{52.37660663424486, 4.8972405493159235, 1012},
			wantErr: false,
		},
		{
			name: "invalid lat",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "5621 HH",
				"dsafhkjl", "4.8972405493159235", "1012 LH"),
			want:    Pos{},
			wantErr: true,
		},
		{
			name: "invalid lon",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "5621 HH",
				"52.37660663424486", "dfshlafjhdk", "1012 LH"),
			want:    Pos{},
			wantErr: true,
		},
		{
			name: "zip too short",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "5621 HH",
				"52.37660663424486", "4.8972405493159235", "1"),
			want:    Pos{},
			wantErr: true,
		},
		{
			name: "invalid zip",
			activity: *makeActivity(t, "51.452446917398824", "5.462562099535168", "5621 HH",
				"52.37660663424486", "4.8972405493159235", "jkajsdhfh"),
			want:    Pos{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDepot(tt.activity)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
