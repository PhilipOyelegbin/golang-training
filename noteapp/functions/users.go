package functions

import (
	"encoding/json"
	"fmt"
	"noteapp/schema"
	"os"
	"strings"
)

// function to generate a new user ID based on existing users
func generateUserId(users []schema.UserData) int {
	// Don't reload from file here - use the users already in memory
	maxId := 0
	for _, user := range users {
		if user.Id > maxId {
			maxId = user.Id
		}
	}
	return maxId + 1
}

// function to save users to JSON file
func saveUsersToFile(users []schema.UserData) error {
	// Create db directory if it doesn't exist
	err := os.MkdirAll("./db", 0755)
	if err != nil {
		return err
	}

	// Convert users slice to JSON (users already contains all data)
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON data to file
	err = os.WriteFile("./db/users.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// function to validate registration input
func validateRegistrationInput(users []schema.UserData, firstName, lastName, email, password string) (bool, bool, bool) {
	var validateNoExistingEmail bool = true
	for _, user := range users {
		if user.Email == email {
			validateNoExistingEmail = false
		}
	}

	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".") && validateNoExistingEmail
	isValidPassword := len(password) >= 6

	return isValidName, isValidEmail, isValidPassword
}

// function to register new user
func RegisterUser(firstName, lastName, email, password string, users []schema.UserData) {
	fmt.Println("#################################################################")

	fmt.Println("Input your firstName:")
	fmt.Scan(&firstName)

	fmt.Println("Input your lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Input your email address:")
	fmt.Scan(&email)

	fmt.Println("Input your password:")
	fmt.Scan(&password)

	isValidName, isValidEmail, isValidPassword := validateRegistrationInput(users, firstName, lastName, email, password)

	if isValidName && isValidEmail && isValidPassword {
		userData := schema.UserData{
			Id: generateUserId(users),
			FirstName: firstName,
			LastName: lastName,
			Email: email,
			Password: password,
		}

		users = append(users, userData)

		// Save users to JSON file
		err := saveUsersToFile(users)
		if err != nil {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Unable to save to file:", err)
		} else {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Message: User registered successfully!")
		}

		fmt.Println(userData)
		fmt.Println("----------------------------------------------------------------")
	} else {
		if !isValidName {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Firstname or Lastname length too short")
			fmt.Println("----------------------------------------------------------------")
		}
		if !isValidEmail {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Not a valid email address or email already exists")
			fmt.Println("----------------------------------------------------------------")
		}
		if !isValidPassword {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Password too short")
			fmt.Println("----------------------------------------------------------------")
		}
	}
}

// function to validate login input
func validateLoginInput(email, password string, users []schema.UserData) bool {
	if len(password) < 6 || !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false
	}

	for _, user := range users {
		if user.Email == email {
			return user.Password == password
		}
	}

	return false
}

// function to login as a user
func LoginUser(email, password string, users []schema.UserData) int {
	fmt.Println("#################################################################")
	var myId int

	fmt.Println("Input your email address:")
	fmt.Scan(&email)
	
	fmt.Println("Input your password: ")
	fmt.Scan(&password)

	isValidLogin := validateLoginInput(email, password, users)

	if isValidLogin {
		for _, user := range users {
			 if user.Email == email {
				myId = user.Id
			 }
		}
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("Message: Loggedin Successful...")
		fmt.Println("----------------------------------------------------------------")
	} else {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("Error: Invalid login details")
		fmt.Println("----------------------------------------------------------------")
	}

	return myId
}