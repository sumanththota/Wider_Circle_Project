package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
    var err error
    dsn := "root:mysql@tcp(localhost:3306)/company"
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Warning: Failed to open DB connection:", err)
        return
    }
    // mitigating the risk of not having db
    if err = db.Ping(); err != nil {
        fmt.Println("Warning: Failed to ping the  DB connection:", err)
        db = nil
    }
}

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
    // enable CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    var employees []Employee
    fmt.Println("server got hit")

    // If DB is available, try fetching from it
    if db != nil {
         rows, err := db.Query("SELECT id, name, title, manager_id FROM employees")
            if err == nil {
                defer rows.Close()
                for rows.Next() {
                    var emp Employee
                    if err := rows.Scan(&emp.ID, &emp.Name, &emp.Title, &emp.ManagerID); err == nil {
                        employees = append(employees, emp)
                    }
                }
            }
    }

    // If found in DB, return those
    if len(employees) > 0 {
        json.NewEncoder(w).Encode(employees)
        return
    }

    // If not found in DB, fetch from external API
    resp, err := http.Get("https://gist.githubusercontent.com/chancock09/6d2a5a4436dcd488b8287f3e3e4fc73d/raw/fa47d64c6d5fc860fabd3033a1a4e3c59336324e/employees.json")
    if err != nil {
        http.Error(w, "Failed to fetch employee data", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&employees); err != nil {
        http.Error(w, "Failed to decode employee data", http.StatusInternalServerError)
        return
    }

    // Optional: Insert fetched employees into DB
    if db != nil{
        for _, emp := range employees {
                _, err := db.Exec("INSERT INTO employees (id, name, title, manager_id) VALUES (?, ?, ?, ?)",
                    emp.ID, emp.Name, emp.Title, emp.ManagerID)
                if err != nil {
                    fmt.Println("Failed to insert employee:", err)
                }
            }
    }
    json.NewEncoder(w).Encode(employees)
}
