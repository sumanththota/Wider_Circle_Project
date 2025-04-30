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
('Michael Chen', 'CEO', NULL),
('Barrett Glasauer', 'CTO', 1),
('Chris Hancock', 'Engineering Manager', 2),
('Julian Early', 'Engineer', 3),
('Michael Lorton', 'Engineer', 3),
('Emily Pun', 'Designer', NULL),
('Andres Green', 'COO', 1);