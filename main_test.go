package main

import (
  "os"
  "fmt"
  "bytes"
  "log"
  "testing"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "strconv"
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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()
  a.Router.ServeHTTP(rr, req)

  return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
  if expected != actual {
    t.Errorf("Expected response code %d. Got %d\n", expected, actual)
  }
}

func TestEmptyTable(t *testing.T) {
  clearTable()

  req, _ := http.NewRequest("GET", "/tasks", nil)
  response := executeRequest(req)

  checkResponseCode(t, http.StatusOK, response.Code)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func TestGetNonExistentTask(t *testing.T) {
  clearTable()

  req, _ := http.NewRequest("GET", "/task/45", nil)
  response := executeRequest(req)

  checkResponseCode(t, http.StatusNotFound, response.Code)

  var m map[string]string
  json.Unmarshal(response.Body.Bytes(), &m)
  if m["error"] != "Task not found" {
    t.Errorf("Expected the 'error' key of the response to be set to 'Task not found'. Got '%s'", m["error"])
  }
}

func TestGetTask(t *testing.T) {
  clearTable()
  addTasks(1)

  req, _ := http.NewRequest("GET", "task/1", nil)
  response := executeRequest(req)

  checkResponseCode(t, http.StatusOK, response.Code)
}

func addTasks(count int) {
  if count < 1 {
    count = 1
  }

  for i := 0; i < count; i++ {
    statement := fmt.Sprintf("INSERT INTO tasks(description, completed) VALUES('%s', false)", ("Task " + strconv.Itoa(i+1)))
    a.DB.Exec(statement)
  }
}

func TestCreateTask(t *testing.T) {
  clearTable()

  payload := []byte(`{"description":"test task","completed":false}`)

  req, _ := http.NewRequest("POST", "/task", bytes.NewBuffer(payload))
  response := executeRequest(req)

  checkResponseCode(t, http.StatusCreated, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["description"] != "test task" {
    t.Errorf("Expected task description to be 'test task'. Got '%v'", m["description"])
  }

  if m["completed"] != false {
    t.Errorf("Expected task completed to be 'false'. Got '%v'", m["completed"])
  }

  if m["id"] != 1.0 {
    t.Errorf("Expected task ID to be '1'. Got '%v'", m["id"])
  }
}

func TestDeleteTask(t *testing.T) {
  clearTable()
  addTasks(1)

  req, _ := http.NewRequest("GET", "/task/1", nil)
  response := executeRequest(req)
  checkResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("DELETE", "/task/1", nil)
  response = executeRequest(req)
  checkResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("GET", "/task/1", nil)
  response = executeRequest(req)
  checkResponseCode(t, http.StatusNotFound, response.Code)
}
