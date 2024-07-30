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
- Docker (optional)

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

To start the server, run:

```bash
./event-booking
```

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
