Todo App - GO
---

A simple todo app to improve golang skills.

#### Usage

##### Install dependencies
```
$ dep ensure
```

##### Run main (dev, tests)
```
$ go run main.go
```

##### Compiled binary
```
$ go build
$ ./todo-app
```

##### Endpoints
```
GET /todos (All todos)
GET /todos/{id} (Specific todo)
POST /todos/{id} (Create todo)
-> POST VALUES, json: {description: string, finished: bool, date: string}
DELETE /todos/{id} (Delete specific todo)
```

#### Roadmap
[x] CRUD api
[ ] Home to view and interact visually with todos
[ ] Unit tests