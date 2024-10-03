package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Sneaker struct {
	ID        int    `json:"id"`
	Model     string `json:"model"`
	Size      int    `json:"size"`
	Color     string `json:"color"`
	Condition string `json:"condition"` // "new", "used", "worn out"
}

var sneakers []Sneaker
var nextID = 1

// Create a new sneaker
func createSneaker(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newSneaker Sneaker
	// Decode the JSON request body into a Sneaker struct
	if err := json.NewDecoder(r.Body).Decode(&newSneaker); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newSneaker.ID = nextID
	nextID++
	sneakers = append(sneakers, newSneaker)

	w.Header().Set("Content-Type", "application/json")
	// Respond with the created sneaker in JSON format
	json.NewEncoder(w).Encode(newSneaker)
}

// Get all sneakers
func getSneakers(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Respond with the list of all sneakers in JSON format
	json.NewEncoder(w).Encode(sneakers)
}

// Get a single sneaker by ID
func getSneakerByID(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the sneaker ID from the URL
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid sneaker ID", http.StatusBadRequest)
		return
	}

	// Find and return the sneaker with the given ID
	for _, sneaker := range sneakers {
		if sneaker.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(sneaker)
			return
		}
	}
	http.Error(w, "Sneaker not found", http.StatusNotFound)
}

// Update an existing sneaker by ID
func updateSneaker(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the sneaker ID from the URL
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid sneaker ID", http.StatusBadRequest)
		return
	}

	// Find and update the sneaker with the given ID
	for i, sneaker := range sneakers {
		if sneaker.ID == id {
			var updatedSneaker Sneaker
			if err := json.NewDecoder(r.Body).Decode(&updatedSneaker); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			updatedSneaker.ID = id // Ensure the ID remains the same
			sneakers[i] = updatedSneaker
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedSneaker)
			return
		}
	}
	http.Error(w, "Sneaker not found", http.StatusNotFound)
}

// Delete a sneaker by ID
func deleteSneaker(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the sneaker ID from the URL
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid sneaker ID", http.StatusBadRequest)
		return
	}

	// Find and delete the sneaker with the given ID
	for i, sneaker := range sneakers {
		if sneaker.ID == id {
			sneakers = append(sneakers[:i], sneakers[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // No content response
			return
		}
	}
	http.Error(w, "Sneaker not found", http.StatusNotFound)
}

// Extract ID from URL path (e.g., /sneakers/1)
func extractID(path string) (int, error) {
	// Split the URL path into parts and extract the ID part
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid path")
	}
	return strconv.Atoi(parts[2]) // Convert ID to integer
}

func main() {
	// Handle requests to the /sneakers route
	http.HandleFunc("/sneakers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getSneakers(w, r) // Get all sneakers
		case http.MethodPost:
			createSneaker(w, r) // Create a new sneaker
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle requests to the /sneakers/{id} route
	http.HandleFunc("/sneakers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getSneakerByID(w, r) // Get sneaker by ID
		case http.MethodPut:
			updateSneaker(w, r) // Update sneaker by ID
		case http.MethodDelete:
			deleteSneaker(w, r) // Delete sneaker by ID
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
