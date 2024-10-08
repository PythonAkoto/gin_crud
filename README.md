# Building A CRUD App With Gin

I will be using Go's Gin framework to build a simple CRUD application.

First we need to initialise our Go application:
```
// in app directory
go init myapp

// in app directory - via guthub
go init github.com/pythonakoto/name_of_app
```

We need to install the gin package:
```
go get -u github.com/gin-gonic/gin
```

myapp/
├── cmd/
│   └── main.go                  # Entry point
├── controllers/                 # HTTP handlers/controllers
│   └── item_controller.go
├── services/                    # Business logic services
│   └── item_service.go
├── repositories/                # Data access layer (optional)
│   └── item_repository.go
├── models/                      # Entity/Model definitions
│   └── item.go
├── routes/                      # Route definitions
│   └── routes.go
├── templates/                   # HTML templates for rendering views
│   └── index.html               # The HTML file you want to render
├── static/                      # Static files (CSS, JS, images)
│   └── style.css
└── config/                      # Configuration, env variables
    └── config.go

For my application I am replacing `myapp` with my github repo: `github.com/pythonakoto/gin_crud`

The config.go file woouild bne used to create environment variables.

Install the following:
```
go get github.com/joho/godotenv
```
Now you fcan create your environment variables in a `.env` file:
```
// example 

APP_ENV=development
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
```

To run the app we can do the following in the command line:
```
 go run .\cmd\main.go
```

As your project grows, you'll want to add more dependencies. 
Each time you install a package using go get, your `go.mod` 
and `go.sum` files will be updated automatically. 
To update or tidy your go.mod file, you can run:
```
go mod tidy
```