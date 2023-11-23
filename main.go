package main

import (
	"bufio"   // Importing bufio for buffered IO operations
	"fmt"     // Importing fmt for formatting strings and printing
	"os"      // Importing os for accessing the standard input
	"strconv" // Importing strconv for converting strings to other types
	"strings" // Importing strings for string manipulation
)

// Function to get input from the user. It prints a prompt and reads a line of input.
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)                // Print the prompt to the user
	input, err := r.ReadString('\n') // Read a line of input until a newline character

	return strings.TrimSpace(input), err // Return the input with surrounding whitespace removed
}

// Function to create a new bill. It prompts the user for a bill name and creates a bill struct.
func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // Create a new reader for reading from standard input

	name, _ := getInput("Create a new bill name: ", reader) // Get bill name from user input

	b := newBill(name)                        // Create a new bill with the given name
	fmt.Println("Created the bill -", b.name) // Print confirmation message

	return b // Return the created bill
}

// Function to prompt the user for different bill actions.
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin) // Create a new reader for reading from standard input

	// Prompt the user to choose an option and read their input
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt { // Switch on the option selected by the user
	case "a": // If the user chooses to add an item
		name, _ := getInput("Item name: ", reader)   // Get item name from user input
		price, _ := getInput("Item price: ", reader) // Get item price from user input

		// Convert the price from string to float64
		p, err := strconv.ParseFloat(price, 64)
		if err != nil { // If there is an error in conversion
			fmt.Println("The price must be a number...") // Print an error message
			promptOptions(b)                             // Re-prompt the user for an option
		}
		b.addItem(name, p) // Add the item to the bill

		fmt.Println("item added -", name, price) // Print confirmation message
		promptOptions(b)                         // Re-prompt the user for an option
	case "t": // If the user chooses to add a tip
		tip, _ := getInput("Enter tip amount ($): ", reader) // Get tip amount from user input

		// Convert the tip from string to float64
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil { // If there is an error in conversion
			fmt.Println("The tip must be a number...") // Print an error message
			promptOptions(b)                           // Re-prompt the user for an option
		}
		b.updateTip(t) // Update the tip amount in the bill

		fmt.Println("tip has been updated to", tip) // Print confirmation message
		promptOptions(b)                            // Re-prompt the user for an option
	case "s": // If the user chooses to save the bill
		b.save()                                      // Save the bill to a file
		fmt.Println("bill has been saved as", b.name) // Print confirmation message
	default: // If the user enters an invalid option
		fmt.Println("That was not a valid option...") // Print an error message
		promptOptions(b)                              // Re-prompt the user for an option
	}
}

// The main function - entry point of the program
func main() {
	mybill := createBill() // Create a new bill
	promptOptions(mybill)  // Start the options prompt loop for the bill
}
