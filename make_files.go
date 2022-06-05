// This is a Golang program to Create a text file

package main // Package main

import ( // Import packages
	"fmt"
	"os"
)

func createFile(fileName, message string) {
	file, err := os.Create(fileName) // Create file
	if err != nil {                  // Check for errors
		fmt.Printf("Unable to open file: %s", err) // Print error if unable to open file
	}

	len, err := file.WriteString(message) // Write message to file

	if err != nil {
		fmt.Printf("Unable to write data: %s", err) // Print error if unable to write data
	}
	file.Close() // Close file

	fmt.Printf("%d character written successfully into file", len) // Print message if file is created successfully
}

func main() { // Main function
	go createFile("helloWorld.txt", "Hello World")          // Call createFile function using a goroutine (a new thread in the background)
	createFile("myName.txt", "My name is INSERT_NAME_HERE") // Call createFile function
}
