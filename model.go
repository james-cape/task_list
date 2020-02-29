package main

import (
  "database/sql"
  "fmt"
)

type task struct {
  ID          int     `json:"id"`
  Completed   bool    `json:"completed"`
  Description string  `json:"description"`
}

func (t *task) getTask(db *sql.DB) error {
  statement := fmt.Sprintf("SELECT description, completed FROM tasks WHERE id=%d", u.ID)
  return db.QueryRow(statement).Scan(&t.Description, &t.Completed))
}
