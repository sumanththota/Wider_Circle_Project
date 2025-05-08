import React from "react";
export interface Employee {
    id: number;
    name: string;
    title: string;
    manager_id: number | null;
    reports: Employee[];
}
interface Props {
  employee: Employee;
}

const EmployeeNode: React.FC<Props> = ({ employee }) => {
  //SORT REPORTS BY LAST NAME — EXACT PLACEMENT
  const sortedReports = [...employee.reports].sort((a, b) => {
    const lastA = a.name.trim().split(" ").slice(-1)[0];
    const lastB = b.name.trim().split(" ").slice(-1)[0];
    return lastA.localeCompare(lastB);
  });

  return (
    <li>
      <div>{employee.title}: {employee.name}</div>

      {sortedReports.length > 0 && (
        <ul>
          {sortedReports.map((report) => (
            <EmployeeNode key={report.id} employee={report} />
          ))}
        </ul>
      )}
    </li>
  );
};

export default EmployeeNode;