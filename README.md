<h1 align="center" id="title">JobKart</h1>

<p id="description">Jobkart is a full-stack job application portal built using Go (Fiber) GORM and a HTML/CSS/JavaScript frontend designed to streamline job applications and candidate management. It supports both employer and employee roles features Google OAuth2 and JWT-based authentication and allows seamless interaction between job seekers and recruiters. The application is fully Dockerized for easy deployment.</p>

  
  
<h2>🧐 Features</h2>

Here're some of the project's best features:

*   ✅ Google OAuth2 & JWT Authentication for both employers and employees
*   🧑‍💼 Separate profile pages for employers and employees
*   📝 Job application functionality for employees to apply to postings
*   📬 Application review system allowing employers to view and manage candidate applications
*   🔍 Job search by profile/keyword helping employees discover relevant listings
*   Vacancy listings displayed on the homepage for easy browsing
*   🐳 Dockerized setup for consistent development and easy deployment

## Setup Instructions
### **1. Install Golang**
1. Download and install Go from [https://golang.org/dl/](https://golang.org/dl/)
2. Verify installation:
   ```sh
   go version
   ```

### **2. Set Up PostgreSQL Database**
1. Install PostgreSQL:
   - **Linux** (Ubuntu/Debian):
     ```sh
     sudo apt update && sudo apt install postgresql postgresql-contrib
     ```
   - **MacOS**:
     ```sh
     brew install postgresql
     ```
   - **Windows**:
     Download and install from [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
2. Start PostgreSQL service:
   ```sh
   sudo service postgresql start
   ```
3. Create a new database:
   ```sh
   sudo -u postgres psql
   CREATE DATABASE food_delivery;
   ```
4. Set up a PostgreSQL user:
   ```sh
   CREATE USER food_admin WITH ENCRYPTED PASSWORD 'securepassword';
   ALTER ROLE food_admin SET client_encoding TO 'utf8';
   ALTER ROLE food_admin SET default_transaction_isolation TO 'read committed';
   ALTER ROLE food_admin SET timezone TO 'UTC';
   GRANT ALL PRIVILEGES ON DATABASE food_delivery TO food_admin;
   ```

### **3. Install GoFiber & GORM**
```sh
go mod init food_delivery

go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

### **4. Install Authentication Libraries**
```sh
go get github.com/dgrijalva/jwt-go
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
```

### **5. Clone the Repository & Setup Backend**
```sh
git clone https://github.com/DevanshS9881/Job_Portal-GO.git
cd Job_Portal
```
- **Create a `.env` file** with:
  ```env
  DB_HOST=localhost
  DB_USER=admin
  DB_PASSWORD=securepassword
  DB_NAME=jobkart
  JWT_SECRET=your_secret_key
  GOOGLE_CLIENT_ID=your_google_client_id
  GOOGLE_CLIENT_SECRET=your_google_client_secret
  ```
- **Run the server**:
  ```sh
  go run main.go
  ```

### **6. Set Up Frontend**
1. Navigate to `frontend/` folder:
   ```sh
   cd frontend
   ```
2. Open `index.html` in a browser:
   ```sh
   open index.html  # MacOS
   start index.html  # Windows
   xdg-open index.html  # Linux
   ```

### **7. Running the Full Application**
- **Start the backend**:
  ```sh
  go run main.go
  ```
- **Open the frontend** in a browser:
  ```sh
  http://localhost:8080/.html](http://127.0.0.1:3000/frontend/index5.html
  ```
## API Endpoints
| Method | Endpoint                       | Description                                       |
| ------ | ------------------------------ | ------------------------------------------------- |
| POST   | `/register`                    | Register a new user                               |
| POST   | `/updateProfileEmployee/:id`   | Update employee profile (requires JWT)            |
| POST   | `/updateProfileEmployer/:id`   | Update employer profile (requires JWT)            |
| GET    | `/getProfile/:id`              | Get user profile (requires JWT)                   |
| DELETE | `/deleteUser/:id`              | Delete user account (requires JWT)                |
| POST   | `/role`                        | Get user role from JWT                            |
| POST   | `/addJob/:id`                  | Add a new job posting (requires JWT)              |
| PUT    | `/updateJob/:id/:Employer_id`  | Update an existing job (requires JWT)             |
| DELETE | `/deleteJob/:id`               | Delete a job posting (requires JWT)               |
| GET    | `/showJob/:id`                 | Show job details by job ID                        |
| GET    | `/getJob/:id`                  | Alias for show job by ID (requires JWT)           |
| GET    | `/getJobs/:id`                 | Get all jobs posted by a specific employer        |
| GET    | `/allJobs`                     | Get all available job listings                    |
| GET    | `/jobs/profiles/:profile`      | Search jobs by profile                            |
| POST   | `/apply/:Emid/:jobID`          | Apply for a job (requires JWT)                    |
| GET    | `/review/:Employer_id/:job_id` | Review applications for a job (requires JWT)      |
| POST   | `/accept/:id/`                 | Accept a job application                          |
| GET    | `/getApplications/:id`         | Get applications submitted by a specific employee |
