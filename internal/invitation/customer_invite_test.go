package invitation

import (
	"CustomerInvitation/internal/model"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetEligibleCustomer(t *testing.T) {

	var tests = []struct {
		input    string
		expected []model.Customer
	}{
		// This customer should be eligible since he/she is within the distance limit (~24km far from the office)
		{
			input: "{\"latitude\": \"53.1229599\", \"user_id\": 6, \"name\": \"John Smith\", \"longitude\": \"-6.2705202\"}",
			expected: []model.Customer{{UserID: 6, Name: "John Smith"}},
		},
		// This customer should not be eligible since he/she is not within the distance limit (~300km far from the office)
		{
			input: "{\"latitude\": \"51.92893\", \"user_id\": 1, \"name\": \"Bob Marley\", \"longitude\": \"-10.27699\"}",
			expected: nil,
		},
		// This should return nil since latitude is missing in the input
		{
			input: "{\"user_id\": 6, \"name\": \"John Smith\", \"longitude\": \"-6.2705202\"}",
			expected: nil,
		},
		// This should return nil since input is empty
		{
			input: " ",
			expected: nil,
		},
	}

	// Test the GetEligibleCustomer
	for _, test := range tests {
		actual, _ := GetEligibleCustomer(strings.NewReader(test.input))
		assert.Equal(t, test.expected, actual, "Expected and actual customer list is not equal")
	}
}

func Test_isCustomerEligible(t *testing.T) {
	// Initialize test data
	latitude1 := 53.1229599
	longitude1 := -6.2705202
	latitude2 := 12.1982765
	longitude2 := -1.6575436
	var tests = []struct {
		input    model.Customer
		expected bool
	}{	// This customer should be eligible since he/she is within the distance limit (~24km far from the office)
		{
			input:    model.Customer{Latitude: &latitude1, UserID: 1, Name: "Vanessa Queen", Longitude: &longitude1},
			expected: true,
		},
		// This customer should not be eligible since he/she is within the distance limit (~4600km far from the office)
		{
			input:    model.Customer{Latitude: &latitude2, UserID: 2, Name: "Alice Smith", Longitude: &longitude2},
			expected: false,
		},
	}

	for _, test := range tests {
		actual := isCustomerEligible(test.input)
		assert.Equal(t, test.expected, actual, "Expected results are not equal to actual")
	}
}

func Test_calculateDistance(t *testing.T) {
	var tests = []struct {
		longitude float64
		latitude  float64
		expected  float64
	}{
		{
			latitude: 52.833502,
			longitude:  -8.522366,
			expected:  161.36,
		},
		{
			latitude: 50,
			longitude:  -7,
			expected:  374.83,
		},
		{
			latitude: 0,
			longitude:  0,
			expected:  5959.28,
		},
	}

	// Check if the calculated distance is same as expected
	for _, test := range tests {
		actual := calculateDistance(test.latitude, test.longitude)
		assert.InEpsilon(t, test.expected,  actual, 1e-2, "Expected distance does not match")
	}
}