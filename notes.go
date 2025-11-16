package main

import "fmt"

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetNotes() ([]Note, error) {
	db := GetDB()

	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := `SELECT id, title, content FROM notes`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []Note

	for rows.Next() {
		var note Note

		err := rows.Scan(&note.ID, &note.Title, &note.Content)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return notes, nil
}

func SaveNotes(note *Note) error {

	db := GetDB()
	if db == nil {
		return fmt.Errorf("db is nil")
	}

	query := `INSERT INTO notes (title, content) VALUES (?, ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(note.Title, note.Content)
	return err

}
