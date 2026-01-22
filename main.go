package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Electronics", Description: "Devices and gadgets"},
	{ID: 2, Name: "Books", Description: "Printed and digital books"},
	{ID: 3, Name: "Clothing", Description: "Apparel and accessories"},
}

func main() {
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		} else if r.Method == "POST" {
			var newCategory Category

			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}

			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCategory)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			for _, category := range categories {
				if category.ID == id {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(category)
					return
				}
			}
			http.Error(w, "Category not found", http.StatusNotFound)

		} else if (r.Method == "PUT") {
			idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
		
			var updatedCategory Category
		
			err = json.NewDecoder(r.Body).Decode(&updatedCategory)
			if err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
		
			for i := range categories {
				if categories[i].ID == id {
					updatedCategory.ID = id
					categories[i] = updatedCategory
		
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(updatedCategory)
					return
				}
			}
			http.Error(w, "Category not found", http.StatusNotFound)
		} else if (r.Method == "DELETE") {
			idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}

			for i := range categories {
				if categories[i].ID == id {
					categories = append(categories[:i], categories[i+1:]...)

					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]string{
						"message": "Category deleted successfully",
					})
					return
				}
			}
		}
	})
	
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
			"message": "API category is running",	
		})
	})

	fmt.Println("Starting server on localhost:8080")
	err:= http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}