package functions

import (
	"encoding/json"
	"fmt"
	"noteapp/schema"
	"os"
)

// function to generate a new note ID based on existing notes
func generateNoteId(notes []schema.NoteData) int {
	maxId := 0
	for _, note := range notes {
		if note.Id > maxId {
			maxId = note.Id
		}
	}
	return maxId + 1
}

// function to save notes to json file
func saveNotesToFile(notes []schema.NoteData) error {
	err := os.MkdirAll("./db", 0755)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("./db/notes.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// function to validate note input
func validateNoteInput(title string, description string, myId int) (bool, bool, bool) {
	isValidTitle := len(title) >= 6
	isValidDescription := len(description) >= 6
	isValidUserId := myId > 0

	return isValidTitle, isValidDescription, isValidUserId
}

// function to create new note
func CreateNote(title string, description string, myId int, notes []schema.NoteData) {
	fmt.Println("#################################################################")
	fmt.Println("Input your title:")
	fmt.Scan(&title)
	
	fmt.Println("Input your description: ")
	fmt.Scan(&description)

	isValidTitle, isValidDescription, isValidUserId := validateNoteInput(title, description, myId)

	if isValidTitle && isValidDescription && isValidUserId {
		noteData := schema.NoteData{
			Id: generateNoteId(notes),
			Title: title,
			Description: description,
			UserId: myId,
		}

		notes = append(notes, noteData)

		err := saveNotesToFile(notes)
		if err != nil {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Unable to save to file:", err)
		} else {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Message: Note created successfully!")
		}

		fmt.Println(noteData)
		fmt.Println("----------------------------------------------------------------")
	} else {
		if !isValidTitle {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Title is too short.")
			fmt.Println("----------------------------------------------------------------")
		}
		if !isValidDescription {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Description is too short.")
			fmt.Println("----------------------------------------------------------------")
		}
		if !isValidUserId {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: User does not exist.")
			fmt.Println("----------------------------------------------------------------")
		}
	}
}

// function to view loggedin user note
func ViewMyNotes(notes []schema.NoteData, myId int) {
	var myNotes = make([]schema.NoteData, 0)

	for _, note := range notes {
		if note.UserId == myId {
			myNotes = append(myNotes, note)
		}
	}

	fmt.Println("----------------------------------------------------------------")
	fmt.Println(myNotes)
	fmt.Println("----------------------------------------------------------------")
}

// function for a user to delete a note by id
func DeleteMyNote(noteId int, myId int, notes []schema.NoteData) {
	fmt.Println("#################################################################")
	fmt.Println("Input the note ID to delete:")
	fmt.Scan(&noteId)

	var noteFound bool = false
	var noteIndex int = -1

	// First, find the note that belongs to the logged-in user and has the specified ID
	for i, note := range notes {
		if note.UserId == myId && note.Id == noteId {
			noteFound = true
			noteIndex = i
			break
		}
	}

	if noteFound {
		// Remove the note from the notes slice
		notes = append(notes[:noteIndex], notes[noteIndex+1:]...)
		
		// Save updated notes to file
		err := saveNotesToFile(notes)
		if err != nil {
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("Error: Unable to save changes to file:", err)
			fmt.Println("----------------------------------------------------------------")
		} else {
			fmt.Println("----------------------------------------------------------------")
			fmt.Printf("Message: Note with ID %d has been deleted successfully!\n", noteId)
			fmt.Println("----------------------------------------------------------------")
		}
	} else {
		fmt.Println("----------------------------------------------------------------")
		fmt.Printf("Error: Note with ID %d does not exist or does not belong to you.\n", noteId)
		fmt.Println("----------------------------------------------------------------")
	}
}
