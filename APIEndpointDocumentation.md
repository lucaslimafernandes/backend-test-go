# API Endpoint Documentation

## Authentication

### Sign Up

- **Endpoint**: `/auth/signup`
- **Method**: POST
- **Description**: Creates a new user account.
- **Request Body**:
```json
{
    "email": "user@example.com",
    "password": "password123",
    "fullname": "John Doe"
}
```
- **Response**:
    - Success: 201 Created
    - Error: 400 Bad Request

### Login

- **Endpoint**: `/auth/login`
- **Method**: POST
- **Description**: Logs in an existing user.
- **Request Body**:
```json 
{ 
    "email": "user@example.com", 
    "password": "password123" 
}
```
- **Response**:
    - Success: 200 OK
    - Error: 401 Unauthorized

### Get User Profile

- **Endpoint**: `/user/profile`
- **Method**: GET
- **Description**: Retrieves the authenticated user's profile.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Response**:
    - Success: 200 OK
    - Error: 401 Unauthorized

### Check Authentication

- **Endpoint**: `/user/isauth`
- **Method**: GET
- **Description**: Checks if the user is authenticated.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Response**:
    - Success: 200 OK
    - Error: 401 Unauthorized

## File Management

### List Files (v1)

- **Endpoint**: `/api/v1/file`
- **Method**: GET
- **Description**: Lists all files for the authenticated user.
- **Headers**:
    - Authorization: Bearer `<token>`
    - folder: folder
- **Response**:
    - Success: 200 OK
    - Error: 401 Unauthorized

### List Files (v2)

- **Endpoint**: `/api/v2/file`
- **Method**: GET
- **Description**: Lists all files for the authenticated user with additional metadata.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Response**:
    - Success: 200 OK
    - Error: 401 Unauthorized
- **Request Body**:
```json
{
	"folder": "folder"
}
```

### Upload File

- **Endpoint**: `/api/v1/file`
- **Method**: POST
- **Description**: Uploads a new file.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Request Body**: Multipart form data with file.
- **Headers**:
    - Content-Type: video/mp4
    - Authorization: Bearer `token`
    - file: file.mp4
    - path: folder
    - userid: 1
    - useremail: "user@example.com"
    - description: description
    - compress: true | false
- **Response**:
    - Success: 201 Created
    - Error: 400 Bad Request

### Mark File as Unsafe

- **Endpoint**: `/api/v1/unsafe`
- **Method**: POST
- **Description**: Marks a file as unsafe.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Request Body**:
```json
{
	"fileid": 1,
	"reviewerid": 4,  
	"unsafe": true
}
```
- **Response**:
    - Success: 200 OK
    - Error: 400 Bad Request

## Folder Management

### List Folders
- **Endpoint**: `/api/v1/folder`
- **Method**: GET
- **Description**: Lists all folders for the authenticated user.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Response**:
    - Success: 200 OK
    - Error: 401 Unauthorized

### Create Folder
- **Endpoint**: `/api/v1/folder`
- **Method**: POST
- **Description**: Creates a new folder.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Request Body**:
```json
{
	"userid": 1,
	"useremail": "user@example.com",
	"folder": "folder"
}
```
- **Response**:
    - Success: 201 Created
    - Error: 400 Bad Request

## Streaming

### Stream File
- **Endpoint**: `/stream/:filekey`
- **Method**: GET
- **Description**: Streams a file.
- **Headers**:
    - Authorization: Bearer `<token>`
- **Request Body**:
```json
{
    "userid": 1
}
```
- **Response**:
    - Success: 200 OK
    - Error: 401 Unauthorized


