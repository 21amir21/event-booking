# Event Booking System

Welcome to the Event Booking System! This project is designed to help users book events seamlessly.

## Table of Contents

- [Event Booking System](#event-booking-system)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [Setup](#setup)
  - [Usage](#usage)
  - [API Functionalities](#api-functionalities)
    - [Authentication](#authentication)
    - [Event Management](#event-management)
    - [User Registration](#user-registration)
    - [Event Booking](#event-booking)

## Installation

To get a local copy up and running, follow these simple steps.

### Prerequisites

- Go
- [Docker (optional)](#using-docker)

### Setup

1. Clone the repository

   ```bash
   git clone https://github.com/21amir21/event-booking.git
   ```

2. Navigate to the project directory

   ```bash
   cd event-booking
   ```

3. Build the project

   ```bash
   go build
   ```

## Usage

### Setting Up Environment Variables

The application requires certain environment variables to be set, especially for JWT (JSON Web Tokens) configuration. These variables should be defined in a `.env` file in the root of your project.

1. **Create a `.env` file:**

   ```bash
   touch .env
   ```

2. **Add the following environment variables to the `.env` file:**

   ```env
   JWT_SECRET=your_jwt_secret_key
   JWT_EXPIRATION_TIME=7200  # Time in seconds,, no need to add it, the code now handles it
   ```

   - `JWT_SECRET`: A secret key used to sign the JWT tokens. It should be a strong, random string.
   - `JWT_EXPIRATION_TIME`: The expiration time for the JWT tokens in seconds.

To start the server, run:

```bash
./event-booking
```

## Using Docker

To build and run the application using Docker, follow these steps:

1. **Build the Docker Image:**

   ```bash
   docker build --no-cache -t event-booking .
   ```

2. **Run the Docker Container:**

   ```bash
   docker run -p 8080:8080 --rm -v $(pwd):/app -v /app/tmp --name event-booking-ctr event-booking
   ```

3. **View Logs:**

   If you need to view the logs of the running container, use:

   ```bash
   docker logs event-booking-ctr
   ```

By using Docker, you can ensure that the application runs in a consistent environment, regardless of the host operating system.

## API Functionalities

### Authentication

- The API uses JSON Web Tokens (JWT) for authenticating users.
- Users must login to receive a token, which must be included in the `Authorization` header for protected routes.

### Event Management

- Create, update, delete, and view events.
- Only authenticated users can create events.
- Events can only be updated or deleted by their creators.

### User Registration

- Users can sign up to create an account.
- Authentication is required to access most features.

### Event Booking

- Users can register for events.
- Users can cancel their registrations for events.
