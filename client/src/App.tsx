import React, { useEffect, useState } from "react";
import EmployeeNode, { type Employee } from "./components/EmployeeNode";

function App() {
  const [data, setData] = useState<Employee[] | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchEmployees = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/employees");
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const json: Employee[] = await response.json();

        // Create map
        const employeeMap: Record<number, Employee> = {};
        json.forEach((emp) => {
          emp.reports = [];
          employeeMap[emp.id] = emp;
        });

        // Build tree
        const rootEmployees: Employee[] = [];
        json.forEach((emp) => {
          if (emp.manager_id === null) {
            rootEmployees.push(emp);
          } else {
            const manager = employeeMap[emp.manager_id];
            if (manager) {
              manager.reports.push(emp);
            }
          }
        });

        console.log("Tree Structure", rootEmployees);
        setData(rootEmployees);
      } catch (err) {
        console.error("Error fetching:", err);
        setError("Failed to load employee data. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchEmployees();
  }, []);

  return (
    <div>
      <h1>Employee Org Chart</h1>
      {loading ? (
        <p>Loading...</p>
      ) : error ? (
        <p style={{ color: "red" }}>{error}</p>
      ) : (
        <ul>
          {data &&
            data.map((emp: Employee) => (
              <EmployeeNode key={emp.id} employee={emp} />
            ))}
        </ul>
      )}
    </div>
  );
}

export default App;
