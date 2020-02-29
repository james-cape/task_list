package main

import (
  "log"
  "github.com/joho/godotenv"
  "os"
)

func init() {
  if os.Getenv("APP_ENV") == "production" {
    log.Print("This environment is production")
  } else if err := godotenv.Load(); err != nil {
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
