# Installations
go
Mysql


# Install
HTTP request handler:
go get github.com/gorilla/mux

mysql go driver
go get github.com/go-sql-driver/mysql

Env variables
go get github.com/joho/godotenv


# Setting GOPATH and PATH
https:github.com/golang/go/wiki/SettingGOPATH
Bash
Add the following to ~/.bash_profile:
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/mysql/bin
```
Source the ~/.bash_profile.
```
source ~/.bash_profile
```

Create the database
$ mysql -u root -p
Enter mysql pw
CREATE DATABASE go_task_list;
USE go_task_list;
CREATE TABLE tasks (
   id INT AUTO_INCREMENT PRIMARY KEY,
   completed BOOLEAN NOT NULL,
   description VARCHAR(255) NOT NULL
);

Env variables:
"DB_USERNAME", "DB_PASSWORD"
