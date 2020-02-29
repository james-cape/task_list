package main

import (
  "database/sql"
  "fmt"
  "log"
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
  _ "github.com/go-sql-driver/mysql"
)

type App struct {
  Router *mux.Router
  DB     *sql.DB
}

// Creates the database connection and establishes routes
func (a *App) Initialize(user, password, dbname string) {
  connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

  var err error
  a.DB, err = sql.Open("mysql", connectionString)
  if err != nil {
    log.Fatal(err)
  }

  a.Router = mux.NewRouter()
}

// Starts the application
func (a * App) Run(addr string) { }

func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

func (a *App) getTask(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid task ID")
    return
  }

  t := task{ID: id}
  if err := t.getTask(a.DB); err != nil {
    switch err {
      case sql.ErrNoRows:
        respondWithError(w, http.StatusNotFound, "Task not found")
      default:
        respondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
  }

  respondWithJSON(w, http.StatusOK, t)
}
