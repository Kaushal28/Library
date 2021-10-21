# Library
Simple library API to store and retrieve books using Go and MongoDB

## Setup

To start the application, run `docker-compose up -d`

## Endpoints available

- List all books (GET /books)
- Filter books by ID (GET /books/{id})
- Add new book(s) (POST /books)

## cURL Commands

1. List all books:

    '''
    curl --location --request GET 'localhost:8080/books'
    '''

2. Filter books by ID:

    '''
    curl --location --request GET 'localhost:8080/books/1'
    '''

3. Add new book(s)

    '''
    curl --location --request POST 'localhost:8080/books' \
    --header 'Content-Type: application/json' \
    --data-raw '[
        {
            "id": "1",
            "title": "Harry Potter and the Philosopher's Stone",
            "author": "J. K. Rowling",
            "cost": 599,
            "year": 1997,
            "publisher": "Bloomsbury"
        }
    ]'
    '''