# Dion Arya Pamungkas_article_fullstack

This is a full-stack web application designed for managing and displaying articles. The backend is built with Golang and the frontend with React. The application utilizes MySQL as its database.

## Repository
[Project Repository](https://github.com/dionarya23/bahaso-test-v2)

## Prerequisites
- Go (Golang)
- Node.js and npm
- MySQL
- Git

## Backend Setup

### Environment Variables
Create a `.env` file in the root of the backend project and populate it with the following variables:

```bash
DB_NAME=your_database_name
DB_PORT=your_database_port
DB_HOST=your_database_host
DB_USERNAME=your_database_username
DB_PASSWORD=your_database_password

JWT_SECRET=your_jwt_secret
BCRYPT_SALT=your_bcrypt_salt_rounds

MAIL_USERNAME=your_email_username
MAIL_PASSWORD=your_email_password
MAIL_HOST=your_email_host
MAIL_PORT=your_email_port
MAIL_SENDER=your_email_sender_address

BASE_URL_FE=http://localhost:3000
```

### Running the Backend
Navigate to the backend directory and run the following commands:

```bash
# init go
go mod tidy

# Install dependencies
go get .

# Run the application
go run ./src/main.go
```

## Frontend Setup

Environment Variables
Create a ```.env``` file in the root of the frontend project and populate it with the following variables:

```bash
REACT_APP_CLOUDINARY_API_KEY=your_cloudinary_api_key
REACT_APP_CLOUDINARY_API_URL=your_cloudinary_api_url
REACT_APP_CLOUDINARY_UPLOAD_PRESET=your_cloudinary_upload_preset
REACT_APP_BASE_URL=http://localhost:8080
```

### Running the Frontend

Navigate to the frontend directory and run the following commands:
```bash
# Install dependencies
npm install

# Start the application
npm start
```
### Accessing the Application
The frontend will be available at http://localhost:3000 and the backend at http://localhost:8080.


### Authors
Dion Arya Pamungkas - Initial work
