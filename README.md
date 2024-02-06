# Backend-Go-API

## Overview

This project is a backend application for an e-commerce platform developed using the Go programming language with the Fiber web framework and MongoDB as the database.

## Features

- [RESTful API]() - The backend provides a set of RESTful APIs to handle various functionalities such as managing products, users and orders.

- [Product Management]() - The backend allows administrators to manage products, including adding new products, updating existing ones, and deleting products.

- [User Management]() - User Authentication and Authorization: The application supports user registration, login, and authentication using JSON Web Tokens (JWT). Additionally, it implements role-based access control to ensure secure access to different parts of the system.

- [Order Processing]() - Users can place orders, and the system handles order processing, including order creation, and order history.

- [Cart Functionality]() - The application supports shopping cart functionality, allowing users to add products to their carts and update quantities.

- [Middleware]() - Custom middleware functions are used to handle tasks such as access authentication.

## Technology Stack

- **Backend (Go):**

  - [Go]() - The programming language used for the backend development.
  - [Fiber](https://github.com/gofiber/fiber/v2) - A fast and lightweight web framework for Go, used for building the RESTful API.
  - [MongoDB]() - The NoSQL database used to store and retrieve data efficiently.
  - [Other Go Libraries]() - Various third-party Go libraries are utilized for tasks such as database interaction, validation, and error handling.

  - **Frontend (React):**
  - [React](https://reactjs.org/) - JavaScript library for building user interfaces
  - [axios](https://github.com/axios/axios) - HTTP client for making requests to the backend
