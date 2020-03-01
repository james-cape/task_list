# Installations
go
Postgres

# Install
HTTP request handler:
go get github.com/gorilla/mux

postgres go driver
go get -u github.com/lib/pq

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
Postgres:
$ psql
jamescape=# CREATE TABLE tasks (
jamescape(# id serial PRIMARY KEY,
jamescape(# description VARCHAR(50) NOT NULL,
jamescape(# completed BOOLEAN NOT NULL);

Env variables:
"DB_USERNAME", "DB_PASSWORD"

Running localhost:
```
go build
```
```
./task_list
```


Creating a go.mod file:
go mod init github.com/james-cape/task_list
go mod tidy

## Endpoints Available

+ [Create a Task](#create_task)
+ [Update a Task](#update_task)
+ [Get All Tasks](#get_all_tasks)
+ [Get Task](#get_a_task)
+ [Delete Task](#delete_a_task)

## <a name="create_task"></a>Create a Task
`http://localhost:8080/task`

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
