package main

import (
  "log"
  "github.com/joho/godotenv"
  "os"
  "strconv"
)

func init() {
  if os.Getenv("DATABASE_URL") != "" {
    log.Print("This environment is production")
  } else if err := godotenv.Load(); err != nil {
    log.Print("No .env file found")
  }
}

func main() {
  var hostname = os.Getenv("DB_HOSTNAME")
  var host_port, _ = strconv.Atoi(os.Getenv("DB_SERVER"))
  var databasename = os.Getenv("DB_NAME")
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
