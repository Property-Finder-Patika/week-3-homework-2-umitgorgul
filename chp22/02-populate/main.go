package main

import "fmt"

func main() {
	//  1. Phone numbers by last name
	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}
	//     Print the dulin's phone number.

	//  2. Product availability by Product ID
	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}
	//     Is Product ID 879401371 available?

	//  3. Multiple phone numbers by last name
	multiPhones := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}
	//     What is Greco's second phone number?

	//  4. Shopping basket by Customer ID
	basket := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}
	//     How many of 576872813 the customer 101 is going to buy?
	//                (Product ID)  (Customer ID)

	// Print dulin's phone number.
	name, phoneNumber := "dulin", "N/A"
	if v, ok := phones[name]; ok {
		phoneNumber = v // overrides phoneNumber (now its 03.37.77.63.06)
	}
	fmt.Printf("%s's phone number: %s\n", name, phoneNumber)

	// Is Product ID 879401371 available?
	id, status := 879401371, "available"
	if !products[id] {
		status = "not " + status // overrides status (now its not available)
	}
	fmt.Printf("Product ID #%d is %s\n", id, status)

	// What is Greco's second phone number?
	name, phoneNumber = "greco", "N/A"
	if phones := multiPhones[name]; len(phones) >= 2 {
		phoneNumber = phones[1] // overrides status (now its 03489900120)
	}
	fmt.Printf("%s's 2nd phone number: %s\n", name, phoneNumber)

	// How many of 576872813 the customer 101 is going to buy?
	cid, pid := 101, 576872813
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, basket[cid][pid], pid)
}
