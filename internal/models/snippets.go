package models

import (
	"database/sql"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. notice how
// the fields of the struct correspond to the field in pur MySQL snippets table\
type Snippet struct {
	ID      uint
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id uint) (*Snippet, error) {
	return nil, nil
}

// This wil return the most 10 recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
