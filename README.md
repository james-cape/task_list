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

Running localhost:
```
go build
```
```
./task_list
```
Navigate to localhost:8080/tasks



## Endpoints Available

+ [Create a Task](#create_task)
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
