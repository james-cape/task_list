package main

import (
  "log"
  "github.com/joho/godotenv"
  "os"
)

func init() {
  if err := godotenv.Load(); err != nil {
    log.Print("No .env file found")
  }
}

func main() {
  var db_username string = os.Getenv("DB_USERNAME")
  var db_password string = os.Getenv("DB_PASSWORD")

  a := App{}
  a.Initialize(db_username, db_password, "go_task_list")

  a.Run(":8080")
}