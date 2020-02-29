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

func (t *task) deleteTask(db *sql.DB) error {
  statement := fmt.Sprintf("DELETE FROM tasks WHERE id=%d", t.ID)
  _, err := db.Exec(statement)
  return err
}

func (t *task) createTask(db *sql.DB) error {
  statement := fmt.Sprintf("INSERT INTO tasks(description, completed) VALUES('%s', %b)", t.Description, T.Completed)
  _, err := db.Exec(statement)

  if err != nil {
    return err
  }

  err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&t.ID)

  if err != nil {
    return err
  }

  return nil
}
