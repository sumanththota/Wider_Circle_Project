package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)



type Employee struct {
    Name      string  `json:"name"`
    ID        int     `json:"id"`
    Title     string  `json:"title"`
    ManagerID *int    `json:"manager_id"` // Use *int to allow null values
}
func main() {

	http.HandleFunc("/api/employees", employeesHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}


func employeesHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic for employees
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://gist.githubusercontent.com/chancock09/6d2a5a4436dcd488b8287f3e3e4fc73d/raw/fa47d64c6d5fc860fabd3033a1a4e3c59336324e/employees.json")

	if err != nil {
		http.Error(w, "Failed to fetch employee data", http.StatusInternalServerError)
        return
	}
	defer resp.Body.Close()

	var employees []Employee
	if err := json.NewDecoder(resp.Body).Decode(&employees); err != nil {
        http.Error(w, "Failed to decode employee data", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(employees); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}
