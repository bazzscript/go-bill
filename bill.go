package main

import (
	"fmt" // Importing the fmt package for formatting strings and printing output
	"os"  // Importing the os package for file operations like writing to a file
)

// Define a struct named 'bill' to represent a bill
type bill struct {
	name  string             // Name of the bill
	items map[string]float64 // Map to hold items and their prices
	tip   float64            // Tip amount
}

// Function to create and return a new bill given a name
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{}, // Initialize an empty map for items
		tip:   0,                    // Initialize tip as 0
	}
	return b // Return the new bill
}

// Method to add an item with its price to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price // Add or update the item in the items map
}

// Method to format the bill into a readable string
func (b *bill) format() string {
	fs := "Bill breakdown:\n" // Starting the formatted string
	var total float64 = 0     // Initialize total amount

	// Loop through each item in the bill
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v) // Format and add each item to the string
		total += v                                    // Add item's price to the total
	}

	// Add tip to the formatted string
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// Calculate and add total amount (including tip) to the formatted string
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs // Return the formatted bill string
}

// Method to update the tip amount for the bill
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip // Update the tip field of the bill
	b.tip = tip // Alternatively, can use this direct assignment
}

// Method to save the bill to a file
func (b *bill) save() {
	data := []byte(b.format())                              // Convert the formatted bill to a byte slice
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644) // Write the byte slice to a file
	if err != nil {                                         // Check for errors during file writing
		panic(err) // If an error occurs, stop execution and print the error message
	}
	fmt.Println("Bill saved to file") // Print a success message after saving the file
}
