CREATE DATABASE IF NOT EXISTS company;

USE company;

CREATE TABLE IF NOT EXISTS employees (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    title VARCHAR(100),
    manager_id INT,
    FOREIGN KEY (manager_id) REFERENCES employees(id)
);

INSERT INTO employees (name, title, manager_id) VALUES 
('Ram', 'CEO', NULL),
('Krishna', 'CTO', 1),
('Dad', 'Engineering Manager', 2),
('MOM', 'Engineer', 3),
('alpanso', 'Engineer', 3),
('Pun', 'Designer', NULL),
('Green', 'COO', 1);