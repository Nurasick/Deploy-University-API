# University Management API

This is the backend for a University Management System, built with Go and the Echo framework. It provides endpoints for managing students, groups, and user accounts in a university environment.

## Overview

The API implements a full-stack backend supporting a web frontend. It focuses on:

- Role-based access control (student, admin(for now))
- JWT-based authentication
- Efficient and concurrent-safe PostgreSQL operations
- Clean, modular project structure

## Key Features

- **User Management**: Create, update, and delete user accounts with different roles.
- **Student & Group Management**: CRUD operations for students and groups, including relational data handling.
- **Authentication & Authorization**: Secure JWT-based login system with role restrictions.
- **Database**: PostgreSQL integration with migrations and connection pooling for concurrency safety.
- **Documentation**: Swagger documentation for all endpoints.
- **Deployment Ready**: Backend is structured for deployment on services like Render.

## Project Structure

- `cmd/` — Entry point of the application  
- `config/` — Environment configuration and settings  
- `database/` — Migrations and database utilities  
- `helpers/jwt/` — JWT authentication helpers  
- `model/` — Data models for users, students, and groups  
- `pkg/` — Core packages with business logic and service handlers  

## About the Project

This backend powers a full-stack university management system used for:

- Managing student data and group assignments
- Handling user roles and access rights
- Providing API endpoints for a React frontend application

## Links
- Render: [University API on render.com]https://deploy-university-api.onrender.com/swagger//index.html#/
- Frontend: [University Management Frontend](https://github.com/Nurasick/University-Frontend)  
- Backend: This repository
