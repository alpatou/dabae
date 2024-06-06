package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	tablename   = "snippets"
	ErrNoRecord = errors.New("models: no such record found")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	statement := fmt.Sprintf(`INSERT into %s  (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`, tablename)
	result, err := model.DB.Exec(statement, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (model *SnippetModel) Get(id int) (*Snippet, error) {

	statement := fmt.Sprintf(`SELECT id, title, content, created, expires FROM %s 
	WHERE expires > UTC_TIMESTAMP() AND id = ?`, tablename)

	row := model.DB.QueryRow(statement, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord

		} else {
			return nil, err
		}
	}

	return s, nil

}

func (model *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
