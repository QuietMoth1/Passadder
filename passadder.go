package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
    "github.com/fatih/color"
)

// Function to check if a file exists in a dictionary
func password_check(filepath string, user_password string) { 
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer file.Close()

	// Splitting the file into lines
	scanner := bufio.NewScanner(file)

	line := 1 // We read the lines where the password is located
	result := 0 // Results counter
	
	for scanner.Scan() {
		// If the password is contained in a dictionary, then we add our counter to 1
		if strings.Contains(scanner.Text(), user_password) { 
			result++
		} 

		line++
	} 


	if err := scanner.Err(); err != nil {
		fmt.Errorf("%v", &err)
	}	

	// If the result does not exceed 1 (it means that the password is not written to the dictionary), then we pass it to the add function
	if result < 1 {
		add_passwords(result, user_password, filepath)
	} else {
        user_password = color.RedString(user_password)
    	fmt.Println("Password -", user_password, "already exists -", result, "time!")    // Проверка на наличие пароля 
	}
	
}


// Function to add a password to a file
func add_passwords(result int, user_password string, filepath string) {
    user_password = color.RedString(user_password)
	fmt.Println("\n\tPassword - ", user_password, "was added!") 

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    _, err = f.WriteString("\n" + user_password)
	if err != nil {
		panic(err)
	}
}



func main() {
	// TO-DO Add line-by-line reading of passwords from another file and passing them to the goroutine channel
	fmt.Print("Enter password: ")
	var user_password string
	fmt.Scanf("%s", &user_password)

	filepath := "/root/passwords.txt" // use your dictionary file to add passwords to it  

	password_check(filepath, user_password)
}
