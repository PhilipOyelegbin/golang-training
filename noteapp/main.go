package main

import (
	"encoding/json"
	"fmt"
	"noteapp/functions"
	"noteapp/schema"
	"os"
)

var firstName, lastName, email, password  string
var title, description string
var myId, noteId int
var users = make([]schema.UserData, 0)
var notes = make([]schema.NoteData, 0)

func main() {
	// Load existing users from file at startup
	userErr := loadUsersFromFile()
	if userErr != nil {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("Error loading existing users:", userErr)
		fmt.Println("----------------------------------------------------------------")
	}

	// Load existing notes from file at startup
	noteErr := loadNotesFromFile()
	if noteErr != nil {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("Error loading existing notes:", noteErr)
		fmt.Println("----------------------------------------------------------------")
	}

	actionSelector()
}

// function to load existing users from JSON file
func loadUsersFromFile() error {
	data, err := os.ReadFile("./db/users.json")
	if err != nil {
		// File doesn't exist yet, that's okay
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return err
	}

	return nil
}

// function to load existing notes from JSON file
func loadNotesFromFile() error {
	data, err := os.ReadFile("./db/notes.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(data, &notes)
	if err != nil {
		return err
	}

	return nil
}

// function to select an action to perform
func actionSelector() {
	fmt.Println("#################################################################")
	fmt.Println("Welcome to the Note Application!")
	fmt.Println("Here you can perform various actions based on your needs.")
	fmt.Println("Please select an action:")
	fmt.Println("1) Register User")
	fmt.Println("2) Create Notes")
	fmt.Println("3) View Note")
	fmt.Println("4) Delete Note")
	fmt.Println("5) Exit")
	fmt.Println("#################################################################")
	
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		functions.RegisterUser(firstName, lastName, email, password, users)
	case 2:
		myId = functions.LoginUser(email, password, users)
		functions.CreateNote(title, description, myId, notes)
	case 3:
		myId = functions.LoginUser(email, password, users)
		functions.ViewMyNotes(notes, myId)
	case 4:
		myId = functions.LoginUser(email, password, users)
		functions.DeleteMyNote(noteId, myId, notes)
	case 5:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}