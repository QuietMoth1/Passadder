package main

import (
	"bufio"
	"fmt"
	"flag"
	"os"
    "github.com/fatih/color"
)

// Function to check if a file exists in a dictionary
func password_check(wordlist string, password_to_add string) { 
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer file.Close()

	// Splitting the file into lines
	scanner := bufio.NewScanner(file)
	
	result := 0 // Results counter
	
	for scanner.Scan() {
		// If the password is contained in a dictionary, then we add our counter to 1
		if scanner.Text() == password_to_add { 
			result++
		}
	} 


	if err := scanner.Err(); err != nil {
		fmt.Errorf("%v", &err)
	}	

	// If the result does not exceed 1 (it means that the password is not written to the dictionary), then we pass it to the add function
	if result < 1 {
		add_passwords(result, password_to_add, wordlist)
	} 
}


// Function to add a password to a file
func add_passwords(result int, password_to_add string, wordlist string) {

	f, err := os.OpenFile(wordlist, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    _, err = f.WriteString(password_to_add + "\n")
	if err != nil {
		panic(err)
	}

	password_to_add = color.RedString(password_to_add)
	fmt.Println("\n\tPassword -", password_to_add, "was added!") 
}

func main() {
	// TO-DO Add line-by-line reading of passwords from another file and passing them to the goroutine channel
	password_to_add := flag.String("password", "", "your file with passwords")
	wordlist := flag.String("wordlist", "", "your wordlist")
	flag.Parse()

	if len(*password_to_add) == 0 && len(*wordlist) == 0 {
		fmt.Println("Usage: passadder -password password123 -wordlist wordlist.txt")
		flag.PrintDefaults()
		os.Exit(0)
	} else {
		password_check(*wordlist, *password_to_add)
	}
}