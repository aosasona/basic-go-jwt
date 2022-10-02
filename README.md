# Basic Golang JWT API

This is a basic API written in Golang. It uses [Gorilla Mux](https://github.com/gorilla/mux) for routing and Sqlite3 as the data store/database ([Gorm](https://gorm.io/) as ORM).

# Running in development

To run the API in development, you will need to have [Golang](https://go.dev/) installed. Once you have that installed, you can run the following commands:

```bash
mv .env.example .env
```

This will rename your `.env.example` file to `.env`. You can then edit the `.env` file to change the port the API runs on or the JWT secret.

```bash
go build && ENV=development ./basic-jwt-api
```

This will build the executable file and start the server on port 8085 by default. The API will be running in development mode, which means that it will use your `.env` file directly by loading it with the [godotenv](github.com/joho/godotenv) package.

OR

```bash
ENV=development go run main.go
```

# Routes

**[POST]**

>`/auth/signup` - Sign up a new user
> 
>`/auth/login` - Login a user
> 
>`/notes` - Create a new note *(protected; requires bearer token)*


**[GET]**
>`/me` - Get current users *(protected; requires bearer token)*
>
>`/notes/{id}` - Get a note by ID *(protected; requires bearer token)*
> 
>`/notes` - Get all notes *(protected; requires bearer token)*