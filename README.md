# Library CRUD App with MongoDB and Go

## Setup

1. Ensure MongoDB is running on your computer.

2. Download the dependencies:
   - For MongoDB operations: `go get go.mongodb.org/mongo-driver`
   - For working with `.env` files: `go get github.com/joho/godotenv`

## API Endpoints

- **GET** http://localhost:8000/get
  - Retrieves all books in the library.

- **GET** http://localhost:8000/get/{name}
  - Retrieves a specific book by its name.

- **POST** http://localhost:8000/create
  - Creates a new book entry in the library.

- **PUT** http://localhost:8000/update/{name}
  - Updates the details of a book specified by its name.

- **DELETE** http://localhost:8000/delete/{name}
  - Deletes a book from the library based on its name.
