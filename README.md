# Backend Engineer Test

Create an API that serves as an cloud backup system

## todo 

folder controller

## Run

### Initialize DB postgres

    ```
    go run migrate/migrate.go

    go run main.go
    ```

## Goals

### Simple Mode

- [x] Users can create an account with:
   - [x] email address
   - [x] password
   - [x] full name
- [x] Users can upload files up to 200mb - up to 50mb (supabase limit)
- [x] Users can download uploaded files
- [x] Users can create folders to hold files

## Project Structure

```
.
├── controllers
│   ├── authController.go       // Authentication controllers
│   ├── fileController.go       // Controllers for file upload/download
│   └── folderController.go     // Controllers for folder management
├── dbinit
│   ├── database.go             // Initialization and connection with the database
│   └── loadEnvs.go             // Loading environment variables
├── example_env                 // Example of environment variables file
├── go.mod
├── go.sum
├── main.go                     // Entry point of the application
├── middlewares
│   ├── checkAuth.go            // Authentication middleware
│   └── adminAuth.go            // Administrator authorization middleware (for Hard Mode)
├── migrate
│   └── migrate.go              // Database migration scripts
├── models
│   ├── authInput.go            // Models for authentication input data
│   ├── user.go                 // User models
│   ├── file.go                 // File models
│   └── folder.go               // Folder models
├── repositories
│   ├── userRepository.go       // User repositories
│   ├── fileRepository.go       // File repositories
│   └── folderRepository.go     // Folder repositories
├── services
│   ├── authService.go          // Authentication business logic
│   ├── fileService.go          // File business logic
│   └── folderService.go        // Folder business logic
├── README.md
├── Dockerfile                  // Docker file to containerize the application
└── docker-compose.yml          // docker-compose file to facilitate container management
```



## Test

### Simple Mode
- Users can create an account with:
    - email address
    - password
    - full name
- Users can upload files up to 200mb
- Users can download uploaded files
- Users can create folders to hold files

### Hard Mode
- An admin user type for managing the content uploaded
- Admins can mark pictures and videos as unsafe
- Unsafe files automatically get deleted
- Users can stream videos and audio

### Ultra Mode
- Compression
- File History

### Bonus
- Revokable session management
- Multiple admin reviews before file is deleted

### How to pick what to work on
At minimum you must implement everything in simple mode. You're free to pick and choose what else you
implement along side it. The harder the task, the better your chances. Though make sure to finish the **Simple Mode**
first.

### Tools/Stack

- NodeJs (TypeScript & Express) or Golang
- Postgres for pure data
- Redis
- Docker
- Postman
- S3 or any other shared cloud storage provider

### Tests

Unit tests are a must, submissions without tests will be ignored.


### Time Duration

7 days

### Submission

1. Your API endpoints should be well documented in POSTMAN.
2. Code should be hosted on a git repository, Github preferably.
3. The API should be hosted on a live server (e.g. https://heroku.com)
4. Your app should be `containerized` using `docker`.
5. Share with us through email the:
    - Repository
    - Hosted API URL
    - Postman Collection
    - A list of tasks you did beyond **Simple Mode**