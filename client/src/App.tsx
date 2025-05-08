import React, { useEffect, useState } from "react";
import EmployeeNode from "./components/EmployeeNode";




function App() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);


useEffect(() => {

  fetch("http://localhost:8080/api/employees")
    .then((res) => {
      return res.json();
    })
    .then((json) => {
      //create map
      const employeeMap = {};
      json.forEach(emp => {
          emp.reports = [];
          employeeMap[emp.id] = emp
          });
      // build tree
      const rootEmployees = [];
      json.forEach(emp => {
          if (emp.manager_id === null){
              rootEmployees.push(emp);
          }else {
              const manager = employeeMap[emp.manager_id];
              if(manager) {
                  manager.reports.push(emp);
                  }
              }

          });
      console.log("tree Structure", rootEmployees);
      setData(rootEmployees);
      setLoading(false);


    })
    .catch((err) => {
      console.error("Error fetching:", err);
      setLoading(false);

    });
}, []);



return (
  <div>
    <h1>Employee Org Chart</h1>
     {loading ? <p>Loading...</p> : (
        <ul>
        {data && data.map((emp: Employee) => (
            <EmployeeNode key={emp.id} employee={emp} />
        ))}
        </ul>
     )}
  </div>
);
}

export default App;