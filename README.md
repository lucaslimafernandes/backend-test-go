# Backend Engineer Test

An API that serves as a cloud backup system


## Run with docker compose

To run the application using Docker Compose, follow these steps:

You need Docker, Gmail email account and and bucket S3 (I reccomend Supabase for this).

### Install Docker

To install docker on Linux:

```bash
curl -fsSL https://get.docker.com | bash
```

### Create an S3 Bucket in Supabase

#### Step 1: Sign Up and Log In

1. Go to the [Supabase website](https://supabase.io/).
2. Click on the "Start your project" button to sign up or log in if you already have an account.

#### Step 2: Create a New Project

1. After logging in, click on the "New project" button.
2. Fill in the required details:
   - **Project name**: Choose a unique name for your project.
   - **Organization**: Select or create an organization.
   - **Database password**: Set a password for your database.
   - **Region**: Choose the region closest to you for optimal performance.
3. Click on the "Create new project" button to proceed.

#### Step 3: Configure Storage

1. Once your project is created, navigate to the "Storage" section in the sidebar.
2. Click on "Create a new bucket".
3. Enter the following details:
   - **Bucket name**: Choose a name for your S3 bucket.
   - **Public**: Decide if you want the bucket to be public or private.
4. Click on "Create bucket".

#### Step 4: Retrieve Bucket Details

1. After creating the bucket, click on the bucket name to open its details.
2. Note down the bucket name and the API URL provided by Supabase. You will need these details for your environment variables.


### Configure enviroment variables

Configure the following environment variables (.env file) to run the application:

- **PORT**: Set the port number that your application will listen on.
  - Example: `PORT=3000`

- **DB_URL**: Connection string for your PostgreSQL database.
  - Example: `DB_URL="host=localhost user=postgres password=password dbname=mydatabase port=5432 sslmode=disable"`

- **SECRET**: Secret key used for JWT authentication.
  - Example: `SECRET=auth-api-jwt-secret`

- **S3_BUCKET**: Name of your S3 bucket where files will be stored.
  - Example: `S3_BUCKET=mybucket`

- **S3_REGION**: AWS region where your S3 bucket is located.
  - Example: `S3_REGION=sa-east-1`

- **S3_ENDPOINT**: Endpoint URL for your S3-compatible storage provider (if using one).
  - Example: `S3_ENDPOINT=http://s3-provider-url:9000`

- **S3_FILEPOINT**: Path prefix or directory structure inside the S3 bucket where files will be stored.
  - Example: `S3_FILEPOINT=s3/storage/v1/object/public/`

- **S3_ACCESS_KEY_ID**: Access key ID for authenticating with your S3 provider.
  - Example: `S3_ACCESS_KEY_ID=my-access-key-id`

- **S3_ACCESS_KEY**: Secret access key corresponding to the access key ID.
  - Example: `S3_ACCESS_KEY=my-secret-access-key`

- **RABBITMQ_HOST**: Connection URL for RabbitMQ server.
  - Example: `RABBITMQ_HOST=amqp://user:password@localhost:5672/`

- **SMTPHOST**: SMTP server host for sending emails.
  - Example: `SMTPHOST=smtp.gmail.com`

- **SMTPPORT**: SMTP server port number.
  - Example: `SMTPPORT=587`

- **SENDER_EMAIL**: Email address used as the sender for outgoing emails.
  - Example: `SENDER_EMAIL=your-email@gmail.com`

- **PASSWD_EMAIL**: Password or app-specific password for the sender email account.
  - Example: `PASSWD_EMAIL=your-email-password`


### Run

Start the application using Docker Compose:

```bash
sudo docker compose up -d
```


## API Endpoints

For detailed information on how to use the API endpoints, please refer to the [API Endpoint Documentation](./API%20Endpoint%20Documentation.md) and the included Postman collection [postman.json](./postman.json).

### Overview

The API provides endpoints for user authentication, file management, folder management, and streaming. Each endpoint is documented with its functionality, request methods, required headers, and request/response formats.

### How to Use

1. **API Endpoint Documentation**:
  - The [API Endpoint Documentation](./API%20Endpoint%20Documentation.md) provides comprehensive details about each endpoint, including descriptions, request methods, request bodies, headers, and possible responses.
  - Use this documentation to understand the purpose of each endpoint and how to interact with it.

2. **Postman Collection**:
   - The [postman.json](./postman.json) file is a Postman collection that includes all the endpoints with predefined requests.
   - Import this file into Postman to test and explore the API easily.
   - To import the collection:
     1. Open Postman.
     2. Click on the "Import" button.
     3. Select the `postman.json` file and import it.
     4. You can now see all the API endpoints organized in the collection and can start testing them.

### Available Endpoints

#### Authentication
- **Sign Up**: `/auth/signup` (POST)
- **Login**: `/auth/login` (POST)

#### User Management
- **Get User Profile**: `/user/profile` (GET)
- **Check Authentication**: `/user/isauth` (GET)

#### File Management
- **List Files (v1)**: `/api/v1/file` (GET)
- **List Files (v2)**: `/api/v2/file` (GET)
- **Upload File**: `/api/v1/file` (POST)
- **Mark File as Unsafe**: `/api/v1/unsafe` (POST)

#### Folder Management
- **List Folders**: `/api/v1/folder` (GET)
- **Create Folder**: `/api/v1/folder` (POST)

#### Streaming
- **Stream File**: `/stream/:filekey` (GET)

For each endpoint, you will find:
- **Endpoint URL**: The URL to access the endpoint.
- **Method**: The HTTP method (GET, POST, etc.) to use.
- **Description**: A brief description of the endpoint’s functionality.
- **Request Body**: The required format for sending data to the endpoint (if applicable).
- **Headers**: Any required headers, such as authorization tokens.
- **Response**: The expected response, including success and error codes.

Use this information to integrate the API into your applications and ensure correct usage of the endpoints.


## Contributing

Your contributions are welcome! If you encounter any bugs or have feature requests, please open an issue. To contribute code, follow these steps:

1. Fork the repository.
2. Clone your forked repository to your local machine.
3. Create a new branch (git checkout -b feature-or-bugfix-name).
4. Make your changes and commit them (git commit -m "Description of your changes").
5. Push your branch to your forked repository (git push origin feature-or-bugfix-name).
6. Open a pull request with a clear description of your changes.


## Goals

### Simple Mode

- [x] Users can create an account with:
   - [x] email address
   - [x] password
   - [x] full name
- [x] Users can upload files up to 200mb - up to 50mb (supabase limit)
- [x] Users can download uploaded files
- [x] Users can create folders to hold files

### Hard Mode

- [x] An admin user type for managing the content uploaded
- [x] Admins can mark pictures and videos as unsafe
- [x] Unsafe files automatically get deleted
- [ ] Users can stream videos and audio

### Ultra Mode

- [x] Compression
- [x] File History

## Extra features

- [x] RabbitMQ:
    - [x] New files notifications
    - [x] Queue to delete files


## Project Structure

```
.
├── controllers
│   ├── authController.go           // Authentication controllers
|   ├── authController_test.go      // Test for authentication controllers
│   ├── fileController.go           // Controllers for file upload/download
|   ├── fileController_test.go      // Test for file controllers
│   ├── folderController.go         // Controllers for folder management
|   └── folderControler_test.go     // Test for folder controllers
├── example_env                     // Example of environment variables file
├── go.mod
├── go.sum
├── main.go                         // Entry point of the application
├── middlewares
│   └── checkAuth.go                // Authentication middleware
├── migrate
│   └── migrate.go                  // Database migration scripts
├── models
│   ├── authInput.go                // Models for authentication input data
│   ├── database.go                 // Initialization and connection with the database
│   ├── file.go                     // File models
│   ├── folder.go                   // Folder models
│   ├── loadEnvs.go                 // Loading environment variables
│   ├── queue.go                    // Initialization and connection with the RabbitMQ
│   └── user.go                     // User models
├── services
│   └── rb                      
│   │  └── services                 // Notify and Delete queues
│   ├── fileService.go              // File business logic
│   └── folderService.go            // Folder business logic
├── README.md
├── Dockerfile                      // Docker file to containerize the application
└── docker-compose.yml              // docker-compose file to facilitate container management
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

