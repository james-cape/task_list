package main

import (
  "database/sql"
  "fmt"
  "log"
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

type App struct {
  Router *mux.Router
  DB     *sql.DB
}

// Creates the database connection and establishes routes
func (a *App) Initialize(host_port int, hostname, username, password, databasename string) {
  pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host_port, hostname, username, password, databasename)

  var err error
  a.DB, err = sql.Open("postgres", pg_con_string)
  if err != nil {
    log.Fatal(err)
  }

  a.Router = mux.NewRouter()
  a.initializeRoutes()
}

// Enable communication to front end
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Starts the application
func (a * App) Run(addr string) {
  log.Fatal(http.ListenAndServe(addr, a.Router))
}

// Routes
func (a *App) initializeRoutes() {
  a.Router.HandleFunc("/tasks", a.getTasks).Methods("GET")
  a.Router.HandleFunc("/task", a.createTask).Methods("POST")
  a.Router.HandleFunc("/task/{id:[0-9]+}", a.getTask).Methods("GET")
  a.Router.HandleFunc("/task/{id:[0-9]+}", a.deleteTask).Methods("DELETE")
  a.Router.HandleFunc("/task/{id:[0-9]+}", a.updateTask).Methods("PUT")
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
  enableCors(&w)
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

func (a *App) updateTask(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid task ID")
    return
  }

  var t task
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&t); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  defer r.Body.Close()
  t.ID = id

  if err := t.updateTask(a.DB); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, t)
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
