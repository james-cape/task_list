# Welcome to Task List Backend!
This project provides simple backend functionality for CRUD actions on a task database. It is intended to be used in managing a to-do list.

This project is written in Go, uses a Postgres database, and is hosted on Heroku at https://task-list-backend-80224.herokuapp.com/.

## Contributor
[James Cape](https://github.com/james-cape)

## Publicly Deployed Links
Back-End: https://task-list-backend-80224.herokuapp.com/

Front-End: TBD

## Local Setup
### System Dependencies
* [Golang 1.14](https://golang.org/dl/)
* [Postgresql 11.2](https://www.postgresql.org/download/)
  * Set GOPATH:
  Add the following to ~/.bash_profile:
  ```
  export GOPATH=$HOME/go
  ```
  Source the ~/.bash_profile.
  ```
  source ~/.bash_profile
  ```

### Environment Dependencies
* **HTTP request handler:** gorilla/mux v1.7.4
```
$ go get github.com/gorilla/mux v1.7.4
```
* **Environment variables handler:** joho/godotenv v1.3.0
```
$ go get github.com/joho/godotenv v1.3.0
```
* **postgres go driver:** lib/pq v1.3.0
```
$ go get github.com/lib/pq v1.3.0
```

### Postgres Database Table
Create the tasks table
```
$ psql
 # CREATE TABLE tasks (
 # id serial PRIMARY KEY,
 # description VARCHAR(50) NOT NULL,
 # completed BOOLEAN NOT NULL);
```

### .env
Create a .env file in the project's root directory and add these environment variables.

```
DB_USERNAME=<Postgres user name>
DB_PASSWORD=<Postgres password>
APP_ENV=development
DB_NAME=go_task_list
DB_SERVER=5432
DB_HOSTNAME=localhost
```

### Testing
```
$ go test -v
```
Expected test results:
```
=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.01s)
=== RUN   TestGetNonExistentTask
--- PASS: TestGetNonExistentTask (0.00s)
=== RUN   TestGetTask
--- PASS: TestGetTask (0.00s)
=== RUN   TestCreateTask
--- PASS: TestCreateTask (0.00s)
=== RUN   TestUpdateTask
--- PASS: TestUpdateTask (0.00s)
=== RUN   TestDeleteTask
--- PASS: TestDeleteTask (0.00s)
PASS
ok  	github.com/james-cape/task_list	0.043s
```

### Compile and Run

Running localhost:
```
$ go build
$ ./task_list
```
You should now have a local server running and ready for http requests.
Postman collection:
https://www.getpostman.com/collections/c012d8fb9142d93e5c84

## Deployment
These instructions will explain how to host this app on Heroku.

### Create a Heroku Account
Register on [Heroku](https://id.heroku.com/login)

### Install Heroku CLI and Login
https://devcenter.heroku.com/articles/heroku-cli

Or, if you have Homebrew:
```
$ brew tap heroku/brew && brew install heroku
$ heroku login
```

### Initialize a Heroku Application and Git Remote
```
$ heroku apps:create <example app name>
```

### Add a Go Module
This go.mod file will autogenerate in the project's root folder.

This will autodetect dependencies used in the project, and guide Heroku to install them in the production environment.
```
$ go mod init github.com/james-cape/task_list
$ go mod tidy
```

### Add a Procfile
Add the following text line:
```
web: bin/task_list
```
`web` is a process type which indicates external http traffic can be received from Heroku's routers.

`bin/task_list` is the command for any Heroku dynos to run on startup.

### Create and Set Up a Remote Postgres Database
```
$ heroku addons:create heroku-postgresql:hobby-dev
$ heroku pg:psql
 # CREATE TABLE tasks (
 # id serial PRIMARY KEY,
 # description VARCHAR(50) NOT NULL,
 # completed BOOLEAN NOT NULL
 # );
```

### Configure the Production Environment Variables
```
$ heroku config:get DATABASE_URL -a task-list-backend-80224
```
The above command will provide your production environment variable information:

`postgres://<username>:<password>@<hostname/server>/<databasename>``

Navigate to https://dashboard.heroku.com/apps/task-list-backend-80224 > Settings > Reveal Config Vars and fill in the following Config Vars:
* DATABASE_URL
* DB_HOSTNAME
* DB_NAME
* DB_PASSWORD
* DB_SERVER
* DB_USERNAME

### Push Your Repo to Heroku
You now have your:
* Heroku app
* Heroku Postgres database
  * Tasks table
* Go Module
* Procfile
* Production Variables (Heroku Config Vars)

Next, commit any changes and push to Heroku:
```
$ git push heroku <branch name>:master
```
Try hitting available endpoints (below) using the production urls.

## Endpoints Available

+ [Create a Task](#create_task)
+ [Update a Task](#update_task)
+ [Get All Tasks](#get_all_tasks)
+ [Get Task](#get_a_task)
+ [Delete Task](#delete_a_task)

## <a name="create_task"></a>Create a Task
`http://localhost:8080/task`

`https://task-list-backend-80224.herokuapp.com/task`

A POST request to `/task/` takes a body with an object of keys:
* `"description":`
* `"completed":`

Example Request:
```
POST http://localhost:8080/task

Body; raw, JSON(application/json):
{
	"description": "task description",
	"completed": false
}
```

Example Response:
```
Status: 201 Created
{
    "id": 1,
    "completed": false,
    "description": "task description"
}
```

## <a name="update_task"></a>Update a Task
`http://localhost:8080/task/:id`

`https://task-list-backend-80224.herokuapp.com/task/:id`

A PUT request to `/task/:id` takes a body with an object of key:
* `"completed":`

Example Request:
```
PUT http://localhost:8080/task/1

Body; raw, JSON(application/json):
{
	"completed": true
}
```

Example Response:
```
Status: 200 OK
{
    "id": 1,
    "completed": true,
    "description": ""
}
```

## <a name="get_all_tasks"></a>Get All Tasks
`http://localhost:8080/tasks`

`https://task-list-backend-80224.herokuapp.com/tasks`

A GET request to `/tasks/` which takes no body

Example Request:
```
GET http://localhost:8080/tasks
```

Example Response:
```
Status: 200 OK
[
    {
        "id": 1,
        "completed": false,
        "description": "description_1"
    },
    {
        "id": 2,
        "completed": false,
        "description": "description_2"
    }
]
```

## <a name="get_a_task"></a>Get Task
`http://localhost:8080/task/:id`

`https://task-list-backend-80224.herokuapp.com/task/:id`

A GET request to `/task/:id` which takes no body

Example Request:
```
GET http://localhost:8080/task/1
```

Example Response:
```
Status: 200 OK
{
    "id": 1,
    "completed": false,
    "description": "description_1"
}
```

## <a name="delete_a_task"></a>Delete Task
`http://localhost:8080/task/:id`

`https://task-list-backend-80224.herokuapp.com/task/:id`

A DELETE request to `/task/:id` which takes no body

Example Request:
```
DELETE http://localhost:8080/task/1
```

Example Response:
```
Status: 200 OK
{
    "result": "success"
}
```
