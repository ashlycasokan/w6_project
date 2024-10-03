neaker CRUD API
This is a simple RESTful CRUD API built with Go for managing a collection of sneakers in your personal wardrobe. The API allows you to create, read, update, and delete sneakers, each with properties like model, size, color, and condition.

Features
Create a new sneaker
Read all sneakers or a sneaker by its ID
Update a sneaker's details
Delete a sneaker by its ID
Prerequisites
Go: Make sure you have Go installed on your machine. You can download it from here.
Getting Started
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/sneaker-crud-api.git
cd sneaker-crud-api
Run the application:

bash
Copy code
go run main.go
The server will start at http://localhost:8080.

API Endpoints
1. Create a Sneaker
Endpoint: POST /sneakers
Description: Creates a new sneaker.
Request Body (JSON):
json
Copy code
{
  "model": "Nike Air Max",
  "size": 9,
  "color": "Red",
  "condition": "new"
}
Response (JSON):
json
Copy code
{
  "id": 1,
  "model": "Nike Air Max",
  "size": 9,
  "color": "Red",
  "condition": "new"
}
2. Get All Sneakers
Endpoint: GET /sneakers
Description: Retrieves all sneakers.
Response (JSON):
json
Copy code
[
  {
    "id": 1,
    "model": "Nike Air Max",
    "size": 9,
    "color": "Red",
    "condition": "new"
  }
]
3. Get Sneaker by ID
Endpoint: GET /sneakers/{id}
Description: Retrieves a sneaker by its ID.
Response (JSON):
json
Copy code
{
  "id": 1,
  "model": "Nike Air Max",
  "size": 9,
  "color": "Red",
  "condition": "new"
}
4. Update Sneaker by ID
Endpoint: PUT /sneakers/{id}
Description: Updates the details of a sneaker by its ID.
Request Body (JSON):
json
Copy code
{
  "model": "Nike Air Max",
  "size": 10,
  "color": "Blue",
  "condition": "used"
}
Response (JSON):
json
Copy code
{
  "id": 1,
  "model": "Nike Air Max",
  "size": 10,
  "color": "Blue",
  "condition": "used"
}
5. Delete Sneaker by ID
Endpoint: DELETE /sneakers/{id}
Description: Deletes a sneaker by its ID.
Response: Returns 204 No Content on success.
Example curl Commands
1. Create a Sneaker
bash
Copy code
curl -X POST http://localhost:8080/sneakers \
    -H "Content-Type: application/json" \
    -d '{"model": "Nike Air Force 1", "size": 10, "color": "White", "condition": "new"}'
2. Get All Sneakers
bash
Copy code
curl -X GET http://localhost:8080/sneakers
3. Get a Sneaker by ID
bash
Copy code
curl -X GET http://localhost:8080/sneakers/1
4. Update a Sneaker
bash
Copy code
curl -X PUT http://localhost:8080/sneakers/1 \
    -H "Content-Type: application/json" \
    -d '{"model": "Nike Air Max", "size": 10, "color": "Blue", "condition": "used"}'
5. Delete a Sneaker
bash
Copy code
curl -X DELETE http://localhost:8080/sneakers/1
Project Structure
bash
Copy code
.
├── main.go            # Main Go file containing the CRUD logic
└── README.md          # Project documentation
Technologies Used
Go: Programming language used to build the API.
net/http: Go's built-in package to handle HTTP requests.
encoding/json: Go's package for encoding/decoding JSON.
