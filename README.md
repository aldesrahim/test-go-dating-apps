# Dating App API

This is a simple and scalable **dating app API** built in Go using the [Gin](https://github.com/gin-gonic/gin) web framework. The application provides essential functionalities for user authentication, profile management, subscription handling, and swiping functionality. It demonstrates clean architecture, middleware usage, and role-based secured routes.

## Features

- **Authentication**:
  - User registration and login with JWT-based authentication.
- **Profile Management**:
  - View and update user profile details.
- **Swiping Mechanism**:
  - Swipe left (pass) or right (like) on available profile.
  - Daily swipe limits and filtered profile suggestions.
- **Subscription System**:
  - View subscription plans and active subscriptions.
  - Purchase subscriptions.
- **Secure Endpoints**:
  - Middleware ensures that only authenticated users can access secured routes.

## Available Routes

### **Public Routes**
- **Authentication**:
  - `POST /api/auth/register`: Register a new user.
  - `POST /api/auth`: Log in with email and password to receive a JWT token.

### **Secured Routes (Require JWT Authentication)**
- **Profile**:
  - `GET /api/profile`: View user profile.
  - `PUT /api/profile`: Update user profile details.
- **Subscriptions**:
  - `GET /api/subscription/active`: View the user's active subscriptions.
  - `GET /api/subscription/products`: List available subscription plans.
  - `POST /api/subscription/buy`: Purchase a subscription.
- **Swipes**:
  - `GET /api/swipes`: Get profile available for swiping.
  - `POST /api/swipes`: Swipe left or right on a profile.

## Technologies Used

- **Backend Framework**: [Gin](https://github.com/gin-gonic/gin) (HTTP web framework for Go)
- **Authentication**: JWT (JSON Web Tokens)
- **ORM**: Gorm (Object Relational Mapper)
- **Database**: MySQL (configurable)
- **Middleware**:
  - JWT-based authentication middleware.
  - Role-based request handling.

## How to Run the Application

### Prerequisites

1. Install [Go](https://golang.org/dl/).
2. Copy `.env.example` into `.env`
3. Run `openssl rand -base64 64` for `JWT_TOKEN`
4. Set up a MySQL database and update your `.env` file with the credentials.
5. Install the required dependencies:
   ```bash
   go mod tidy
   ```

### Run the Application
1. Start the server:
   ```bash
   go run main.go
   ```
2. Access the API at `http://localhost:8005/api`.

## Configuration

- Use a `.env` file to configure environment-specific variables, such as database credentials and JWT secrets.
- Example `.env`:
  ```
  DB_USER=root
  DB_PASSWORD=yourpassword
  DB_HOST=localhost
  DB_PORT=3306
  DB_NAME=dating_app
  JWT_SECRET=your_jwt_secret
  ```

## Project Structure

```plaintext
.
├── controllers      # Business logic and route handlers
├── middlewares      # JWT and authentication middleware
├── models           # Database models and relationships
├── routes           # API route definitions
├── utils            # Utility functions and helpers
├── main.go          # Application entry point
```

## Future Improvements

- Add unit and integration tests.
- Implement advanced swipe matching algorithms.
- Add admin features for managing users and subscriptions.
- Support multi-tenancy for deployment across different regions.
