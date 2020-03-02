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
  statement := fmt.Sprintf("SELECT description, completed FROM tasks WHERE id=%d", t.ID)
  return db.QueryRow(statement).Scan(&t.Description, &t.Completed)
}

func (t *task) deleteTask(db *sql.DB) error {
  statement := fmt.Sprintf("DELETE FROM tasks WHERE id=%d", t.ID)
  _, err := db.Exec(statement)
  return err
}

func (t *task) createTask(db *sql.DB) error {
  statement := fmt.Sprintf("INSERT INTO tasks(description, completed) VALUES('%v', false)", t.Description)
  _, err := db.Exec(statement)

  if err != nil {
    return err
  }

  err = db.QueryRow("SELECT id FROM tasks ORDER BY id DESC LIMIT 1").Scan(&t.ID)

  if err != nil {
    return err
  }

  return nil
}

func (t *task) updateTask(db *sql.DB) error {
  statement := fmt.Sprintf("UPDATE tasks SET completed=%t WHERE id=%d", t.Completed, t.ID)
  _, err := db.Exec(statement)
  return err
}

func getTasks(db *sql.DB) ([]task, error) {
  statement := fmt.Sprintf("SELECT id, description, completed FROM tasks")
  rows, err := db.Query(statement)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  tasks := []task{}

  for rows.Next() {
    var t task
    if err := rows.Scan(&t.ID, &t.Description, &t.Completed); err != nil {
      return nil, err
    }
    tasks = append(tasks, t)
  }

  return tasks, nil
}
