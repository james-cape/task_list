package main

import (
  "database/sql"
  "errors"
)

type task struct {
  ID          int     `json:"id"`
  Completed   bool    `json:"completed"`
  Description string  `json:"description"`
}

func (u *task) getTask(db *sql.DB) error {
  return errors.New("Not implemented")
}
