package main

import (
  "os"
  "log"
  "testing"
)

// Sets application we want to test as 'a'
var a App

func TestMain(m *testing.M) {
  var db_username string = os.Getenv("DB_USERNAME")
  var db_password string = os.Getenv("DB_PASSWORD")

  a = App{}
  a.Initialize(db_username, db_password, "go_task_list")

  ensureTableExists()

  code := m.Run()

  clearTable()

  os.Exit(code)
}

func ensureTableExists() {
  if _, err := a.DB.Exec(tableCreationQuery); err != nil {
    log.Fatal(err)
  }
}

func clearTable() {
  a.DB.Exec("DELETE FROM tasks")
  a.DB.Exec("ALTER TABLE tasks AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS tasks
(
  id INT AUTO_INCREMENT PRIMARY KEY,
  completed BOOLEAN NOT NULL,
  description VARCHAR(255) NOT NULL
)`
