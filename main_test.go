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
  const (
    hostname = "localhost"
    host_port = 5432
    databasename = "go_task_list"
  )
  var username = os.Getenv("DB_USERNAME")
  var password = os.Getenv("DB_PASSWORD")

  a = App{}
  a.Initialize(host_port, hostname, username, password, databasename)

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
  a.DB.Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS tasks
(
  id serial PRIMARY KEY,
  completed BOOLEAN NOT NULL,
  description VARCHAR(50) NOT NULL
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

  req, _ := http.NewRequest("GET", "/task/1", nil)
  response := executeRequest(req)

  checkResponseCode(t, http.StatusOK, response.Code)
}

func addTasks(count int) {
  if count < 1 {
    count = 1
  }

  for i := 0; i < count; i++ {
    statement := fmt.Sprintf("INSERT INTO tasks(description, completed) VALUES('%s', %t)", ("Task " + strconv.Itoa(i+1)), false)
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
    t.Errorf("Expected task completed to be 'false'. Got '%t'", m["completed"])
  }

  if m["id"] != 1.0 {
    t.Errorf("Expected task ID to be '1'. Got '%v'", m["id"])
  }
}

func TestUpdateTask(t *testing.T) {
  clearTable()
  addTasks(1)

  req, _ := http.NewRequest("GET", "/task/1", nil)
  response := executeRequest(req)
  var originalTask map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &originalTask)

  payload := []byte(`{"completed":true}`)

  req, _ = http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
  response = executeRequest(req)

  checkResponseCode(t, http.StatusOK, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["id"] != originalTask["id"] {
    t.Errorf("Expected the id to remain the same(%v). Got %v", originalTask["id"], m["id"])
  }

  if m["completed"] == originalTask["completed"] {
    t.Errorf("Expected completed to change from '%v' to '%v'. Got '%v'", originalTask["completed"], m["completed"], m["completed"])
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
