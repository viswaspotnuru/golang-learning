package models // Declares the package name as 'models', used for organizing related code

import (
	"rest-api/db" // Imports the db package for database connection
	"time"        // Imports the time package for handling date and time
)

// Defines the Event struct with fields representing event data
type Event struct {
	ID          int64     // Unique identifier for the event
	Name        string    `binding:"required"` // Event name, required for validation
	Description string    `binding:"required"` // Event description, required for validation
	Location    string    `binding:"required"` // Event location, required for validation
	DateTime    time.Time `binding:"required"` // Event date and time, required for validation
	UserID      int       // ID of the user who created the event
}

// Declares a package-level slice to hold events in memory (not used with DB in this code)
var events = []Event{}

// Method to save an Event to the database
func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)` // SQL query to insert a new event

	stmt, err := db.DB.Prepare(query) // Prepares the SQL statement, Prepare() + stmt.Exec() (when we inserted data into the database)
	if err != nil {
		return err // Returns error if statement preparation fails
	}
	defer stmt.Close() // Ensures the statement is closed after execution

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // Executes the statement with event data
	if err != nil {
		return err // Returns error if execution fails
	}
	id, err := result.LastInsertId() // Gets the ID of the newly inserted event
	e.ID = id                        // Sets the event's ID field
	return err                       // Returns any error from LastInsertId
}

// Function to retrieve all events from the database
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events" // SQL query to select all events
	rows, err := db.DB.Query(query) // Executes the query
	if err != nil {
		return nil, err // Returns error if query fails
	}
	defer rows.Close() // Ensures rows are closed after processing

	var events []Event // Slice to hold the retrieved events

	for rows.Next() { // Iterates over each row in the result set
		var event Event // Creates a new Event struct
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		// Scans the row data into the event struct fields

		if err != nil {
			return nil, err // Returns error if scanning fails
		}

		events = append(events, event) // Adds the event to the slice
	}

	return events, nil // Returns the slice of events and nil error
}
