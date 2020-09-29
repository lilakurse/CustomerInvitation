package main

import (
	"CustomerInvitation/internal/invitation"
	"log"
	"os"
)

/*
This is the main function for the project. It interacts with the backend to process the input file and generate
the results after calculation. Given the path of the input file, it makes two function calls from the backend:
(i) GetEligibleCustomer(ipFile) -- this will process the input file and return all eligible customers
(ii) GenerateOutputFile(customers, opFilePath) -- this will output into a file all customers sorted by User ID (ascending)
	that are staying within the distance limit
*/

func main() {
	ipFilePath := "customer.txt"
	opFilePath := "output.txt"
	ipFile, err := os.Open(ipFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer ipFile.Close()

	eligibleCustomers, err := invitation.GetEligibleCustomer(ipFile)
	if err != nil {
		log.Fatal(err)
	}

	//invitation.GenerateInvitation(customers)
	invitation.GenerateOutputFile(eligibleCustomers, opFilePath)
}
