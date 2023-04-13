package test

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func MockPlotData() bson.D {
	return bson.D{
		{"plot_id", "0"},
		{"plot_name", "Myplot"},

		{"plot_zones", bson.A{
			bson.D{
				{"zone_id", "2"},
				{"zone_ranges", bson.A{
					bson.D{
						{"zipcode_from", 1},
						{"zipcode_to", 4},
						{"iso_country", "NLD"},
						{"zone_range_id", 123},
					},

					bson.D{
						{"zipcode_from", 5},
						{"zipcode_to", 10},
						{"iso_country", "NLD"},
						{"zone_range_id", 2},
					},
				},
				},
				{"zone_fuel_cost", 0},
				{"zone_driving_time", 0},
			},
		}},
		{"plot_created_at", "0001-01-01T00:00:00Z"},
		{"plot_saved_at", "0001-01-01T00:00:00Z"},
		{"origin", "Algo"},
	}
}

func MockUsersData() bson.D {
	return bson.D{
		{"user_id", 1},

		{"user_plots", bson.A{
			bson.D{
				{"user_plot_id", "0"},
				{"user_plot_name", "myPlot"},
			},
			bson.D{
				{"user_plot_id", "1"},
				{"user_plot_name", "AccuratePlot"},
			},
		},
		},

		{"uuid", "user2432"},
	}
}
func CreateAccessTokenString(userId string) (string, error) {
	// Define the token claims, including the user ID
	claims := jwt.MapClaims{
		"user_id": userId,
	}

	// Define the token signing method and secret key
	signingMethod := jwt.SigningMethodHS256
	secret := []byte("my_secret_key")

	// Create the token
	token := jwt.NewWithClaims(signingMethod, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	// Return the Bearer string access token
	return fmt.Sprintf("Bearer %s", tokenString), nil
}
