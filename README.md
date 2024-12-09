# Test Go Dating App

A lightweight backend service for a dating application built using **Go** and the **Gin web framework**. This application provides features for user authentication, profile management, subscription handling, and swiping functionality, with JWT-based secured routes.

## Features

- **Authentication**:
  - Register a new user.
  - Log in with email and password to get a JWT token.
- **Profile Management**:
  - View and update user profiles.
- **Swiping**:
  - Browse available profiles for swiping.
  - Like or pass on profiles with daily swipe limits.
- **Subscription System**:
  - View active subscriptions and available plans.
  - Purchase a subscription.

## API Endpoints

### Public Routes
- `POST /api/auth/register`: Register a new user.
- `POST /api/auth`: Log in and obtain a JWT.

### Secured Routes (Require JWT Authentication)
- **Profile**:
  - `GET /api/profile`: Get the user's profile.
  - `PUT /api/profile`: Update profile details.
- **Swipes**:
  - `GET /api/swipes`: Fetch profiles available for swiping.
  - `POST /api/swipes`: Perform a swipe action (like or pass).
- **Subscriptions**:
  - `GET /api/subscription/active`: View active subscriptions.
  - `GET /api/subscription/products`: List subscription plans.
  - `POST /api/subscription/buy`: Purchase a subscription.

## Tech Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) (HTTP web framework for Go)
- **Authentication**: JWT (JSON Web Tokens)
- **Database**: MySQL (or configurable)

## Getting Started

### Prerequisites
1. Install [Go](https://golang.org/dl/).
2. Set up a MySQL database and configure `.env`.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/aldesrahim/test-go-dating-apps.git
   cd test-go-dating-apps
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Use a long and complex key to enhance security. You can generate a secure key using a tool like openssl:
   ```
   openssl rand -base64 64
   ```
4. Configure the `.env` file with your database credentials:
   ```
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=dating_app
   JWT_SECRET=your_secret_key
   ```
5. Run the application:
   ```bash
   go run main.go
   ```

## Contributing
Feel free to fork the project and submit pull requests for new features or improvements.

## License
This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
