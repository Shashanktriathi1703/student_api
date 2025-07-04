# Student API Project

![go_react_app](https://github.com/user-attachments/assets/fcf500ad-cc35-480e-9670-663a2057b93b)


## Project Description
This project is a full-stack student management application featuring a Go backend API and a React frontend. The backend provides RESTful endpoints to manage user data stored in a PostgreSQL database, while the frontend offers a user-friendly interface to interact with the API.

## Technologies Used
- **Backend:** Go, Gorilla Mux, PostgreSQL, rs/cors
- **Frontend:** React, Axios, React Router DOM, Create React App
- **Database:** PostgreSQL

## Project Structure
```
.
├── cmd/                  # Backend application entry point
│   └── main.go           # Main server setup and routing
├── internal/             # Backend internal packages
│   ├── config/           # Configuration loading
│   ├── database/         # Database connection setup
│   ├── handler/          # HTTP handlers
│   ├── model/            # Data models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic layer
├── frontend/             # React frontend application
│   ├── public/           # Public assets
│   ├── src/              # React source code
│   ├── package.json      # Frontend dependencies and scripts
│   └── README.md         # Frontend-specific documentation
├── go.mod                # Go module dependencies
├── go.sum                # Go module checksums
└── README.md             # This file
```

## Backend Setup and Running

### Prerequisites
- Go 1.18 or higher
- PostgreSQL database

### Configuration
The backend loads configuration from environment variables or configuration files (refer to `internal/config`).

### Running the Backend Server
1. Ensure your PostgreSQL database is running and accessible.
2. Configure the database connection parameters.
3. From the project root, run:
   ```bash
   go run cmd/main.go
   ```
4. The server will start on port `8080`.

## Frontend Setup and Running

### Prerequisites
- Node.js (v14 or higher)
- npm (v6 or higher)

### Running the Frontend
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm start
   ```
4. Open your browser and visit [http://localhost:3000](http://localhost:3000).

For more detailed frontend instructions, refer to the [frontend README](frontend/README.md).

## API Endpoints

| Method | Endpoint       | Description           |
|--------|----------------|-----------------------|
| POST   | /users         | Create a new user     |
| GET    | /users         | Get all users         |
| GET    | /users/{id}    | Get user by ID        |
| DELETE | /users/{id}    | Delete user by ID     |

## Contribution Guidelines
Contributions are welcome! Please fork the repository and create a pull request with your changes. Ensure your code follows the existing style and includes appropriate tests.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
