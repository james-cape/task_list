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
  const (
    hostname = "localhost"
    host_port = 5432
    databasename = "go_task_list"
  )
  var username = os.Getenv("DB_USERNAME")
  var password = os.Getenv("DB_PASSWORD")

  port := os.Getenv("PORT")
  if port == "" {
      port = "8080" // Default port if not specified
  }

    a := App{}
    a.Initialize(host_port, hostname, username, password, databasename)

  a.Run(":" + port)

}
