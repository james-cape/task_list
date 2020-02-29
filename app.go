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

// Routes
func (a *App) initializeRoutes() {
  a.Router.HandleFunc("/tasks", a.getTasks).Methods("GET")
  a.Router.HandleFunc("/task", a.createTask).Methods("POST")
  a.Router.HandleFunc("/task/{id:[0-9]+}", a.getTask).Mthods("GET")
  a.Router.HandleFunc("/task/{id:[0-9]+}", a.deleteTask).Methods("DELETE")
}

// Response Helpers
func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

// Controllers
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

func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {
  tasks, err := getTasks(a.DB)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, tasks)
}

func (a *App) createTask(w http.ResponseWriter, r *http.Request) {
  var t task
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&t); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  defer r.Body.Close()

  if err := t.createTask(a.DB); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusCreated, t)
}

func (a *App) deleteTask(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid task ID")
    return
  }

  t := task{ID: id}
  if err := t.deleteTask(a.DB); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, map[string]string{"result":"success"})
}
