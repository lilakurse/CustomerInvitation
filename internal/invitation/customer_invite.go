package invitation

import (
	"CustomerInvitation/internal/model"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math"
	"sort"
)

/*
This function will get eligible customers within 100km of the office.
Arguments:
	ipFile io.Reader
Return:
	[]model.CustomerList - list of eligible customers
	error - error if any error occurred
*/
func GetEligibleCustomer(ipFile io.Reader) ([]model.Customer, error) {
	// Initialize the eligibleCustomersList
	var eligibleCustomersList []model.Customer
	customer := model.Customer{}

	// Create new scanner to read from input file
	scanner := bufio.NewScanner(ipFile)
	// Read the line
	for scanner.Scan() {
		line := scanner.Bytes()
		// Decode line into Customer struct
		errDecoder := json.NewDecoder(bytes.NewReader(line)).Decode(&customer)
		if errDecoder != nil {
			log.Println(errDecoder)
			return nil, errDecoder
		}
		// If latitude or longitude is missing return nil
		if customer.Latitude == nil || customer.Longitude == nil {
			return nil, errors.New("latitude or longitude is missing")
		}

		// If the customer is not eligible for the invitation, continue
		if !isCustomerEligible(customer) {
			continue
		}
		// If eligible, add the customer to a list
		eligibleCustomer := model.Customer{UserID: customer.UserID, Name: customer.Name}
		eligibleCustomersList = append(eligibleCustomersList, eligibleCustomer)
	}

	return eligibleCustomersList, nil
}

/*
This function checks if customer is within the distance limit.
Arguments:
	customer model.CustomerList - customer
Return:
	bool - if customer is within the distance limit from the office
*/
func isCustomerEligible(customer model.Customer) bool {
	// Call the function that will calculate the difference
	distance := calculateDistance(*customer.Latitude, *customer.Longitude)
	// If the distance is within the distance limit return true
	if distance <= model.DistanceLimit {
		return true
	}
	return false
}

/*
This function calculate the distance between customer and the office.
Arguments:
	customerLatitude float64 - customer latitude
	customerLongitude float64 -  customer longitude
Return:
	float64 - distance between customer and the office
*/
func calculateDistance(customerLatitude float64, customerLongitude float64) float64 {
	// Convert customer and office latitudes to radians
	phi1 := customerLatitude * math.Pi / 180
	phi2 := model.OfficeLatitude * math.Pi / 180
	// Take absolute difference of customer and office longitudes in radians
	deltaLambda := math.Abs(model.OfficeLongitude-customerLongitude) * math.Pi / 180

	// Calculate the spherical distance
	distance := math.Acos(math.Sin(phi1)*math.Sin(phi2)+math.Cos(phi1)*math.Cos(phi2)*math.Cos(deltaLambda)) *
		model.R

	return distance
}

/*
This function generate the output and write it to the file.
Arguments:
	customers []model.CustomerList - list of eligible customers
*/
func GenerateOutputFile(customers []model.Customer, opFilePath string) {
	// Sort eligible customers by their User ID (ascending)
	sort.Slice(customers, func(i, j int) bool { return customers[i].UserID < customers[j].UserID })
	// Generate a JSON with all eligible customers
	output, err := json.MarshalIndent(customers, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	// Write results into output file
	// If the file does not exist ioutil.WriteFile will create it and write the data
	err = ioutil.WriteFile(opFilePath, output, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
