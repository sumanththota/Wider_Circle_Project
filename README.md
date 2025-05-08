

# Wider Circle Project

A full-stack application that visualizes organizational hierarchies using employee-manager relationships. It is built using React (frontend), Go (backend), and MySQL (database), with Docker for simplified setup and deployment.

## Project Overview

This project showcases the ability to integrate backend and frontend services with clean architecture, failover handling, and recursive data rendering.

- Backend built in Go serves employee data from a MySQL database and Api Endpoint
- React frontend renders an organizational chart hierarchically
- If the MySQL service is unreachable, the backend falls back to fetch from a static API
- Docker Compose is used to simplify MySQL container setup with preloaded data

## Tech Stack

- Frontend: React, TypeScript, Tailwind CSS
- Backend: Go (net/http)
- Database: MySQL (Dockerized)
- DevOps: Docker, Docker Compose

## Prerequisites

- Node.js
- Go (1.20+)
- Docker and Docker Compose

## Setup Instructions

### 1. Clone the repository

```
git clone https://github.com/sumanththota/Wider_Circle_Project.git
cd Wider_Circle_Project
```

### 2. Start MySQL with Docker

```
docker compose up
```

This starts a MySQL container with a preloaded `init.sql` file for the employee dataset.

### 3. Start the Go server

In a new terminal:

```
cd server
go run main.go
```

This starts a local Go server that serves employee data at `http://localhost:8080/employees`.

### 4. Start the React client

In another terminal:

```
cd client
npm install
npm run dev
```

This starts the frontend on Vite's default port, typically `http://localhost:5173`.

## Usage

1. Open `http://localhost:5173` in your browser.
2. The app will request employee data from the Go server.
3. If the server can connect to MySQL, it retrieves the data from the database.
4. If MySQL is down, it gracefully falls back to the external API source.
5. The frontend displays a tree view of the employee hierarchy, sorted by last name.

## File Structure Overview

```
Wider_Circle_Project/
├── client/                 # React frontend
│   ├── src/
│   └── ...
├── server/                 # Go backend
│   ├── main.go
│   └── ...
├── Docker/
│   ├── docker-compose.yml
│   └── mysql/
│       └── init.sql
└── Readme.md
```

## Notes

- Ensure Docker is running before executing `docker compose up`
- The MySQL container mounts `init.sql` to preload data automatically
- Make sure port `8080` (server) and `5173` (client) are available
